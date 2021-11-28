# mtnmomo-go

[![Build](https://github.com/NdoleStudio/mtnmomo-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/mtnmomo-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/mtnmomo-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/mtnmomo-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/mtnmomo-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/mtnmomo-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/mtnmomo-go)](https://goreportcard.com/report/github.com/NdoleStudio/mtnmomo-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/mtnmomo-go)](https://github.com/NdoleStudio/mtnmomo-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/mtnmomo-go?color=brightgreen)](https://github.com/NdoleStudio/mtnmomo-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/mtnmomo-go)](https://pkg.go.dev/github.com/NdoleStudio/mtnmomo-go)


This package provides a generic `go` client template for the MTN Mobile Money API

## Installation

`mtnmomo-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/mtnmomo-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/mtnmomo-go"
```


## Implemented

- [API User](#api-user)
  - `POST {baseURL}/apiuser`: Create API User
  - `POST {baseURL}/apiuser/{APIUser}/apikey`: Create API Key
  - `GET {baseURL}/apiuser/{APIUser}`: Get API user information
- [Collection](#collection)
  - `POST {baseURL}/collection/token/`: Create access token

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
    "github.com/NdoleStudio/mtnmomo-go"
)

func main()  {
    client := mtnmomo.New(
        mtnmomo.WithBaseURL("https://sandbox.momodeveloper.mtn.com/v1_0"),
        mtnmomo.WithSubscriptionKey(""/* Subscription key */),
    )
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
apiUser, response, err := statusClient.APIUser.CreateAPIUser(context.Background())
if err != nil {
    //handle error
}
```

### API User

#### `POST {baseURL}/apiuser`: Create API User

Used to create an API user in the sandbox target environment.

```go
userID, response, err := client.APIUser.CreateAPIUser(
	context.Background(),
	uuid.NewString(),
	"providerCallbackHost",
)

if err != nil {
    log.Fatal(err)
}

log.Println(response.HTTPResponse.StatusCode) // 201
```

#### `POST {baseURL}/apiuser/{APIUser}/apikey`: Create API Key

Used to create an API key for an API user in the sandbox target environment.

```go
apiKey, response, err := client.APIUser.CreateAPIKey(context.Background(), "userID")

if err != nil {
    log.Fatal(err)
}

log.Println(apiKey) // e.g "f1db798c98df4bcf83b538175893bbf0"
```

#### `GET {baseURL}/apiuser/{APIUser}`: Get API user information

Used to get API user information.

```go
apiUser, response, err := client.APIUser.CreateAPIKey(context.Background(), "userID")

if err != nil {
    log.Fatal(err)
}

log.Println(apiUser.TargetEnvironment) // e.g "sandbox"
```

### Collection

#### `POST {baseURL}/collection/token/`: Create access token

This operation is used to create an access token which can then be used to authorize and authenticate towards the other end-points of the API.

```go
authToken, response, err := client.Collection.Token(context.Background())
if err != nil {
	log.Fatal(err)
}

log.Println(authToken.AccessToken) // e.g eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9....
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

### Security

If you discover any security related issues, please email arnoldewin@gmail.com instead of using the GitHub issues.

## Credits

- [Acho Arnold](https://github.com/achoarnold)
- [All Contributors](../../contributors)
- [Go HTTP Client Template](https://github.com/NdoleStudio/go-http-client)


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
