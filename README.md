

# util
[![Documentation](https://godoc.org/https://github.com/rasheedhamdawi/util?status.svg)](https://godoc.org/github.com/rasheedhamdawi/util)
[![Build Status](https://travis-ci.org/rasheedhamdawi/util.svg?branch=master)](https://travis-ci.org/rasheedhamdawi/util)
[![Go Report Card](https://goreportcard.com/badge/github.com/rasheedhamdawi/util)](https://goreportcard.com/report/github.com/rasheedhamdawi/util)
[![codecov](https://codecov.io/gh/rasheedhamdawi/util/branch/master/graph/badge.svg)](https://codecov.io/gh/rasheedhamdawi/util)



## Example

```go
package main

import (
	"fmt"
	"github.com/rasheedhamdawi/util"
)

func main() {

	// init
	list := new(util.ArrayList)

	// add items
	list.Add("Go")
	list.Add("Python")
	list.Add("Java")
	list.Add("JavaScript")

	// get the length of the list
	list.Size() // 4

	// get item
	list.Get(1) // "Python"

	// get index of item
	list.IndexOf("Java") // 2

	// remove item
	list.Remove("JavaScript")

	// Iteration
	iterator := list.Iterator()

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}

	// Loop
	printItem := func(_ int, item util.Element) {
		fmt.Println(item)
	}

	list.ForEach(printItem)
}

```
check the [documentation](https://godoc.org/github.com/rasheedhamdawi/util) for more
