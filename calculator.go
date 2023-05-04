package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func calculate(expression string) (float64, error) {
    // Use regular expressions to extract the numbers, operators, and parentheses from the expression
    reNum := regexp.MustCompile(`\d+\.?\d*`)
    reOp := regexp.MustCompile(`[-+*/]`)
    reParen := regexp.MustCompile(`\(([^\(\)]+)\)`)
    numbers := reNum.FindAllString(expression, -1)
    operators := reOp.FindAllString(expression, -1)
    parens := reParen.FindAllStringSubmatch(expression, -1)
    // Check for missing numbers, operators, or parentheses
    if len(numbers) == 0 {
        return 0, fmt.Errorf("missing numbers in expression")
    }
    if len(operators) == 0 {
        return 0, fmt.Errorf("missing operators in expression")
    }
    // Process parentheses recursively
    for _, p := range parens {
        innerExpr := p[1]
        innerResult, err := calculate(innerExpr)
        if err != nil {
            return 0, err
        }
        expression = strings.Replace(expression, p[0], strconv.FormatFloat(innerResult, 'f', -1, 64), 1)
    }
    numbers = reNum.FindAllString(expression, -1)
    operators = reOp.FindAllString(expression, -1)
    // Convert the numbers to floats
    var numFloats []float64
    for _, numStr := range numbers {
        num, err := strconv.ParseFloat(numStr, 64)
        if err != nil {
            return 0, fmt.Errorf("invalid number '%s' in expression", numStr)
        }
        numFloats = append(numFloats, num)
    }
    // Process multiplication and division next
    for i := 0; i < len(operators); i++ {
        if operators[i] == "*" || operators[i] == "/" {
            num1 := numFloats[i]
            num2 := numFloats[i+1]
            var result float64
            if operators[i] == "*" {
                result = num1 * num2
            } else {
                result = num1 / num2
            }
            operators = append(operators[:i], operators[i+1:]...)
            numFloats = append(numFloats[:i], append([]float64{result}, numFloats[i+2:]...)...)
            i--
        }
    }

    // Process addition and subtraction last
    result := numFloats[0]
    for i := 0; i < len(operators); i++ {
        if operators[i] == "+" {
            result += numFloats[i+1]
        } else {
            result -= numFloats[i+1]
        }
    }

    // Return the final result
    return result, nil
}




