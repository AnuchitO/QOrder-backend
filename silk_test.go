package main

import (
	"net/http/httptest"
	"testing"

	"github.com/matryer/silk/runner"
)

func TestOrdersEndpoint(t *testing.T) {
	s := httptest.NewServer(newAPI().MakeHandler())
	defer s.Close()

	//runner.New(t, s.URL).RunGlob(filepath.Glob("*.silk.md"))
	runner.New(t, s.URL).RunFile("./orders.silk.md")
}
