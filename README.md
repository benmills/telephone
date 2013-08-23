# Telephone

Interacting with HTTP services is a fundamental part of many Go applications and libraries. Telephone improves on `net/http` by make it easier to create requests using any method (`GET`, `POST`, `PUT`, ect) and including parameters in the request.

## Examples

Making a GET request

```go
response := telephone.Request{url: "http://api.com"}.Get()
```

Making a GET request with URL parameters

```go
request := telephone.Request{
  url: "http://api.com",
  parameters: telephone.Parameters{
    "user_id": "1234",
  }
}

response := request.Get()
//GET http://api.com?user_id=1234
```

## Todo

* Error handling
* Supporting all HTTP verbs
* Support sending JSON easily
