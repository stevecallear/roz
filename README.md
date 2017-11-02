# Roz
[![Build Status](https://travis-ci.org/stevecallear/roz.svg?branch=master)](https://travis-ci.org/stevecallear/roz)
[![codecov](https://codecov.io/gh/stevecallear/roz/branch/master/graph/badge.svg)](https://codecov.io/gh/stevecallear/roz)
[![Go Report Card](https://goreportcard.com/badge/github.com/stevecallear/roz)](https://goreportcard.com/report/github.com/stevecallear/roz)

Roz is a slug generator for Go.

## Getting Started
`go get github.com/stevecallear/roz`

``` go
package main

import (
    "fmt"

    "github.com/stevecallear/roz`"
)

func main() {
    s := roz.New("My", "URL Slug")
    fmt.Println(s)
    // Output: "my-url-slug"
}
```