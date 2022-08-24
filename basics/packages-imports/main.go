package main

import (
	pkg_1 "github.com/nabiel-syarif/golang-tour/basics/packages-imports/pkg-1"
	pkg_2 "github.com/nabiel-syarif/golang-tour/basics/packages-imports/pkg-2"
)

// exported function, types, etc should begins with capital letter
// all function, types, etc that doesn't begin with capital letter considered as private

func main() {
	pkg_1.HelloPkg1()
	pkg_2.HelloPkg2()
}
