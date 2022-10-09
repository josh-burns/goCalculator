package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/knetic/govaluate.v2"
)

func main() {
	operators := []string{"+", "-", "/", "*"}
	isSumComplete := false
	var sum []string

	fmt.Println("Enter Sum, or 'c' to clear sum")
	for isSumComplete == false {
		var input string
		fmt.Println(sum)
		fmt.Scan(&input)

		if input == "=" {
			isSumComplete = true
		}

		// if is number
		if _, err := strconv.Atoi(input); err == nil {
			sum = append(sum, input)
		}

		// if is operator, but not double occurrence of operator
		for _, a := range operators {
			if input == a && sum[len(sum)-1] != a {
				sum = append(sum, input)
			}
			if sum[len(sum)-1] == a {
				sum = sum[:len(sum)-1]
				sum = append(sum, input)
			}
		}

		// if is a letter
		if input == "c" {
			sum = nil
		} else {
			for _, r := range input {
				if unicode.IsLetter(r) {
					fmt.Println("Enter Number or Symbol (+, -, /, *, =)")

				}
			}
		}
	}

	if isSumComplete {
		fmt.Println("sum complete")

		sumString := strings.Join(sum, " ")
		sumToEvaluate, err := govaluate.NewEvaluableExpression(sumString)

		if err != nil {
			fmt.Println("Invalid Sum...", err)
		}

		result, err := sumToEvaluate.Evaluate(nil)

		if err != nil {
			fmt.Println("Invalid Sum...", sum, err)
		}

		fmt.Println(result)
	}
}
