## Golang Password Generation Library

[![Build Status](https://travis-ci.com/szybia/go-password.svg?branch=master)](https://travis-ci.com/szybia/go-password)
[![GoDoc](https://godoc.org/github.com/szybia/go-password?status.svg)](https://godoc.org/github.com/szybia/go-password)

Performance-oriented library for the generation of random passwords.

Inspired by Seth Vargos [go-password](https://github.com/sethvargo/go-password/) implementation.

### Installation

```sh
$ go get -u github.com/szybia/go-password/password
```
* * *
cli tool for generating passwords:
```sh
$ go get -u github.com/szybia/go-password/randpw
```

### Usage

```golang
package main

import (
	"fmt"

	"github.com/szybia/go-password/password"
)

func main() {
	//	Generate a random password of length 50
	pass, err := password.GenerateLength(50)
	if err != nil {
		panic(err)
	}
	fmt.Println(pass)
}
```

Output:
```text
Rg,FCu/4{'(ZWB4HR5D~%R[vA{ITM^!i5\{pkq;VnIj?=O"CL3
```
* * *
You can also create your own generators with a specific character set as shown below:

```golang
//	Create generator with specified character set
g := password.NewGenerator(&password.CharSet{
	Lowercase: "abc",
	Uppercase: "ABC",
	Digits:    "012",
	Symbols:   "",
})

//	Generate a password consisting of
//	5 lowercase, uppercase letters and 5 digits
pass, err := g.Generate(5, 5, 5, 0) //  output: 1Cc1c02cAB2CBba

//  Generate a password of length 15
pass, err := g.GenerateLength(15)   //  output: a0BC1B1cbA1a10B
```
