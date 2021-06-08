// run -gcflags=-G=3

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func Zero[T any]() (_ T) { return }

type T[X any] int

func (T[X]) M() {
	var have interface{} = Zero[X]()
	var want interface{} = Zero[MyInt]()

	if have != want {
		println("FAIL")
	}
}

type I interface{ M() }

type MyInt int
type U = T[MyInt]

var x = U(0)
var i I = x

func main() {
	x.M()
	U.M(x)
	(*U).M(&x)

	i.M()
	I.M(x)
}
