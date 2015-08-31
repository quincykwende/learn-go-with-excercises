package main

import (
	"fmt"
	"strings"
)

//usd to eur rate i.e 1usd == 0.7eur
const RATE float64 = 0.7


func main(){
	
	var currency string
	var amount float64
	
	fmt.Printf("Currency Calculator:\nConverts USD to Eur and vice versa \nEnter amount: ")
	
	fmt.Scanf("%f", &amount)
	
	fmt.Printf("Enter present currency (USD or Eur): ")
	
	fmt.Scanf("%s", &currency)
	
	//convert currency string to uppercase to minimise error
	currency = strings.ToUpper(currency)
	
	amt, cur, err := convert(amount, currency)
	
	if err == true {
		fmt.Printf("Wrong Currency used, Please choose either USD or EUR")
	} else {
		fmt.Printf("%f %s = %f %s", amount, currency, amt, cur)
	}
	
}

/* 
 * returns 3 values (new amount, new currency and status; true for error)
 * convert from USD to EUR and viceversa
 */
func convert(amt float64, c string) (float64, string, bool) {
	
	if c == "EUR" {
		return amt / RATE, "USD", false
	} else if c == "USD" {
		return amt * RATE, "EUR", false
	} 
	
	return 0, "", true
}
