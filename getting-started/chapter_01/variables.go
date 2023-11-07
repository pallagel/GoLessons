package main

import "fmt"

func main() {
	//typical declaration
	//var <variable name> <type>
	var name string
	var number int //int has signed integers and unsigned Integers
	var pi float32
	var google complex128
	var isVisible bool
	var simple byte  //Alias for uint8
	var uniCode rune //represent Unicode characters. Alias for int32

	//Now its time to assigne values to all the variables
	name = "Peter Pan"
	number = 12
	pi = 3.1415926535897932384626433832795
	google = 10 ^ 100
	isVisible = true
	simple = 255
	uniCode = 'A'

	//Here is how to print all values in the variables.
	fmt.Printf("My name is %v\n", name)
	fmt.Printf("My Age is %v\n", number)
	fmt.Printf("Do you know the value of PI? %.10f\n", pi)
	fmt.Printf("The google is a big number. %v\n", google)
	fmt.Printf("The visibility is set to %v\n", isVisible)
	fmt.Printf("A simple value is %v\n", simple)
	fmt.Printf("The unicode value for 'A' is %x\n", uniCode)
	fmt.Println("#############################################\n")

	//Go also provide a convenience way to declare and assigne values to variables
	//Insted of above 2 steps, lets follow the new way
	username := "lpallage"
	age := 12
	radius := 5.33333333333
	quadrillion := 1000000000000000
	isTrue := true
	oneByte := 254
	code := 'H'

	//lets print all the variables
	fmt.Printf("My name is %v\n", username)
	fmt.Printf("My Age is %v\n", age)
	fmt.Printf("Do you know the value of radius? %.10f\n", radius)
	fmt.Printf("The quadrillion is a big number. %v\n", quadrillion)
	fmt.Printf("The visibility is set to %v\n", isTrue)
	fmt.Printf("A simple value is %v\n", oneByte)
	fmt.Printf("The unicode value for 'H' is %x\n", code)
}
