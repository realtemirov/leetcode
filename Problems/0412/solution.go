package problem0412

import "fmt"

func FizzBuzz(n int) []string {
	resp := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			resp[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			resp[i-1] = "Fizz"
		} else if i%5 == 0 {
			resp[i-1] = "Buzz"
		} else {
			resp[i-1] = fmt.Sprintf("%d", i)
		}
	}

	return resp
}
