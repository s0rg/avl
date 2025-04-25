[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/avl)](https://pkg.go.dev/github.com/s0rg/avl)
[![License](https://img.shields.io/github/license/s0rg/avl)](https://github.com/s0rg/avl/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/avl)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/avl?sort=semver)](https://github.com/s0rg/avl/tags)

[![CI](https://github.com/s0rg/set/workflows/ci/badge.svg)](https://github.com/s0rg/set/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/avl)](https://goreportcard.com/report/github.com/s0rg/avl)
[![Maintainability](https://qlty.sh/badges/717f9cfb-49e0-475d-bcb6-3efafa6823e4/maintainability.svg)](https://qlty.sh/gh/s0rg/projects/avl)
[![Code Coverage](https://qlty.sh/badges/717f9cfb-49e0-475d-bcb6-3efafa6823e4/test_coverage.svg)]
![Issues](https://img.shields.io/github/issues/s0rg/avl)

# avl

Generic [AVL Tree](https://en.wikipedia.org/wiki/AVL_tree) implementation for golang

# features

- simple API
- generic
- range-iter support
- zero-dependency
- 100% test-covered

# example

```go
package main

import (
	"fmt"

	"github.com/s0rg/avl"
)

func main() {
	tree := avl.New[int, string]()

    tree.Add(4, "four")
    tree.Add(2, "two")
    tree.Add(5, "five")
    tree.Add(1, "one")
    tree.Add(3, "three")

    tree.Del(1)

    if val, ok := tree.Get(2); ok {
        fmt.Println("value of 2:", val)
    }

    for k, v := range tree.Iter {
        fmt.Println("key", k, "value", v)
    }
}
```

# license

MIT
