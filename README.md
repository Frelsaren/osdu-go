# OSDU-Go

## Description
OSDU-Go is a Go client library for interacting with the OSDU (Open Subsurface Data Universe) data platform. It provides a set of tools and utilities to facilitate the integration and management of subsurface data within the OSDU ecosystem.

NB: Package is still under development, not all endpoints have been implemented and/or tested

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)

## Installation
To install OSDU-Go, you need to have Go installed on your machine. You can then use `go get` to fetch the library.

```bash
go get github.com/Frelsaren/osdu-go
```

## Usage
```go
package main

import (
        "context"
        "net/url"

        "github.com/Frelsaren/osdu-go/osdu"
)

func main() {
        ctx := context.Background()
        BaseURL, e := url.Parse("http://localhost:8080/")
        if e != nil {
                panic(e)
        }
        token := acquireToken()
        partition := "default"

        client := osdu.Client{
                BaseURL:   BaseURL,
                Token:     &token,
                Partition: &partition,
        }

        client.Initialize()

        params := osdu.GetRecordOfKindParams{
                Kind: "osdu:wks:master-data--Field:1.0.0",
        }

        client.Storage.GetRecordsOfKind(ctx, params)
}
```
