go-search
=========

[![Build Status](https://travis-ci.org/magdkudama/go-search.png?branch=master)](https://travis-ci.org/magdkudama/go-search)

"go-search" is a simple library to find elements in a slice. It accepts multiple input types. It's been written in 30 minutes (including travis, tests and doc), so bugs are likely to appear, yes :)

[Click to view documentation (auto-generated)](https://godoc.org/github.com/magdkudama/go-search)

## Installation

```
go get github.com/magdkudama/go-search
```

## Quick Start

```go
package main

import (
	"github.com/magdkudama/search"
	"fmt"
)

func main() {
	found, pos, err := search.Search()

	fmt.Println("Was element found?: ", found)
	fmt.Println("In what position?: ", pos)
	fmt.Println("Any error?: ", err)
}
```

## Contributors

- Magd Kudama [magdkudama]
