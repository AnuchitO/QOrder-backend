# Menus resources

## GET /menus

Get all of menus items
* Content-Type: `application/json`
* Accept: `application/json`

===

### Example response

* Status: `200`
* Content-Type: `application/json; charset=utf-8`

```
[{"id":"57e3f8e85e812900039b1234","menu_name":"ต้มยำกุ้ง","price":99}]
```

## POST /menus

* Content-Type: `application/json`
* Accept: `application/json`

```json
{
  "menu_name": "ต้มยำกุ้ง",
  "price": 99.5
}
```
===

### Example response
* Status: `201`

## PUT /menus/57e3f8e85e812900039b1234
* Content-Type: `application/json`
* Accept: `application/json`

```json
{
  "menu_name": "ต้มยำกุ้ง",
  "price": 99
}
```

===

### Example response
* Status: `200`

## DELETE /menus/57e3f8e85e812900039b1234
* Content-Type: `application/json`
* Accept: `application/json`

===

### Example response
* Status: `200`

