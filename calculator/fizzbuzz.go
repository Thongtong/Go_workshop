package calculator

import "strconv"

func FizzBuzz(number int) string {
	// if number == 1 {
	// 	return "1"
	// } else if number == 2 {
	// 	return "2"
	// }
	// return ""

	// if number == 3 || number == 6 || number == 9{
	// 	return "Fizz"
	// }
	
if number%3 == 0 && number%5 == 0 {
	return "FizzBuzz"
 }
	
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}


	return strconv.Itoa(number)
}