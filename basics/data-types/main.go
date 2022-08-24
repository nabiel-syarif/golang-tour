package main

import (
	"fmt"
	"math"
)

func main() {
	// zero value for numeric is 0
	// zero value for boolean is false
	// zero value for string is empty string ("")

	numberTypes()
	stringType()
	booleanType()
}

func numberTypes() {
	discreteNumber := func() {
		// 64 bit large in 64bit computer (32 bit large in 32 bit computer) (signed integer)
		var maxInt int = math.MaxInt
		// max 8 bit with negative value (signed integer)
		var maxInt8 int8 = math.MaxInt8
		// max 16 bit with negative value (signed integer)
		var maxInt16 int16 = math.MaxInt16
		// max 32 bit with negative value (signed integer)
		var maxInt32 int32 = math.MaxInt32
		// max 64 bit with negative value (signed integer)
		var maxInt64 int64 = math.MaxInt64

		// max 64 bit in 64bit computer (32bit in 32bit computer) without negative value (unsigned integer)
		var maxUint uint = math.MaxUint
		// max 64 bit without negative value (unsigned integer)
		var maxUint8 uint = math.MaxUint8
		// max 64 bit without negative value (unsigned integer)
		var maxUint16 uint = math.MaxUint16
		// max 64 bit without negative value (unsigned integer)
		var maxUint32 uint = math.MaxUint32
		// max 64 bit without negative value (unsigned integer)
		var maxUint64 uint = math.MaxUint64

		fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n", maxInt, maxInt8, maxInt16, maxInt32, maxInt64, maxUint, maxUint8, maxUint16, maxUint32, maxUint64)
	}

	floatingNumber := func() {
		var maxFloat32 = math.MaxFloat32
		var maxFloat64 = math.MaxFloat64

		fmt.Printf("%v\n%v\n", maxFloat32, maxFloat64)
	}

	complexNumber := func() {
		// 32 bit each for real and imaginary parts
		var complex64Number complex64 = 2i
		// 64 bit each for real and imaginary parts
		var complex128Number complex128 = 1e10 + 200i

		fmt.Printf("%v\n%v\n", complex64Number, complex128Number)
	}

	aliases := func ()  {
		// alias for uint8
		var byteVar byte = 255

		// alias for int32, can also be used for char with unicode supports
		var runeVar rune = 'c'

		fmt.Printf("%v\n%v\n", byteVar, runeVar)
	}

	discreteNumber()
	floatingNumber()

	complexNumber()
	aliases()

}

func stringType() {
	var name string = "Nabiel"
	fmt.Println(name)
}

func booleanType() {
	var trueVar bool = true
	var falseVar bool = false
	fmt.Println(trueVar, falseVar)
}
