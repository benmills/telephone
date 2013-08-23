# Telephone

[![Build Status](https://travis-ci.org/benmills/telephone.png)](https://travis-ci.org/benmills/telephone)

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

## Installation

Get telephone with the `go get` command: `go get github.com/benmills/telephone`. You can then import `github.com/benmills/telephone` in your go projects.

## Todo

* Error handling
* Supporting all HTTP verbs
* Support sending JSON easily
