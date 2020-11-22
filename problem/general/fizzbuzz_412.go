package general

import "strconv"

func fizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		temp := ""
		if i%3 == 0 {
			temp = "Fizz"
		}
		if i%5 == 0 {
			temp += "Buzz"
		}
		if i%3 != 0 && i%5 != 0 {
			temp = strconv.Itoa(i)
		}
		result = append(result, temp)
	}
	return result
}
