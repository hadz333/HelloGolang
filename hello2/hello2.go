package main 

import ( 
	"fmt"
	"errors"
	"math"
)

type person struct {
	name string 
	age int
}

func main() {
	// declaration of array
	a := []int{5, 2, 3, 1, 4}
	// array doesn't have fixed size 
	a = append(a, 13)
	fmt.Println(a)

	// maps in Go
	vertices := make(map[string]int)

	vertices["Triangle"] = 2
	vertices["Square"] = 3
	vertices["Hexagon"] = 5
	fmt.Println(vertices)
	delete(vertices, "Square")
	fmt.Println("After delete: ", vertices)

	// for loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	// for loops are the only kind of loop in Go
	for i < 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("This is an array for loop example")
	arr := []string{"Hello", "Hi", "Ni Hao"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
	fmt.Println("This is a map for loop example")
	m := make(map[string]int)
	m["A"] = 0
	m["B"] = 1

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	fmt.Println(sum(2, 8))

	num := 9
	// func can return multiple values
	result, err := sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Square root of", num, "is", result)
	}

	// structs
	structExample()

	// pointer example
	pointerExample()
}

func structExample() {
	p := person{name: "Hank Hill", age: 42}
	fmt.Println(p)
	fmt.Println(p.age)
}

func pointerExample() {
	i := 7
	fmt.Println("Pointer example")
	fmt.Println(i)
	increment(&i)
	fmt.Println(&i) // memory address
	fmt.Println(i)
}

func increment(x *int) {
	*x++;
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

