package main
import (
	"fmt"
)

/*
 Request for a number from the user and return it's factorial
*/

func main() {
	
	var number int64
	
	fmt.Println("Factorial \nEnter a number to get it's factorial: ")
	fmt.Scan(&number)
	
	fmt.Printf("The Factorial of %d is %d", abs(number), factorial(number))
}

// function that returns factorial
func factorial(number int64) int64{
	
	if number < 0 {
		number = abs(number)
	}
	
	if number == 0 {
		return 1
	}
			
	return number * factorial(number-1)
	
}

// returns absolute value of a negative number
func abs(number int64) int64{
	
	if number < 0 {
		return number * -1
	}
	
	return number
}
