package main

import "fmt"

type Calculator2d interface {
	CalculateArea() float64
}

var _ Calculator2d = &SquareCalculator{}

type SquareCalculator struct {
	Side float64
}

// implement stringer interface to square calculator
var _ fmt.Stringer = &SquareCalculator{}

func (s SquareCalculator) String() string {
	return fmt.Sprintf("Side : %f\n", s.Side)
}

func (s SquareCalculator) CalculateArea() float64 {
	return s.Side * s.Side
}

type NilValueInterface interface {
	Call()
}
type NilValue struct {
}

func (n *NilValue) Call() {
	if n == nil {
		fmt.Println("Called with a nil receiver")
	}
}

func main() {
	var calculator2d Calculator2d = SquareCalculator{
		Side: 1e2,
	}
	fmt.Println(calculator2d.CalculateArea())

	var nilValueInterface NilValueInterface
	var nilValue *NilValue
	nilValueInterface = nilValue
	nilValueInterface.Call()

	var emptyInterface interface{} = 100
	emptyInterface = SquareCalculator{}
	emptyInterface = "Empty interface can be any type"
	fmt.Println(emptyInterface)
}
