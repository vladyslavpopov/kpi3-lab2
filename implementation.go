package lab2

import (
	"fmt"
	"strings"
)

func PrefixToLisp(input string) (string, error) {
	if len(input) == 0 {
		return "", fmt.Errorf("empty input")
	}
	tokens := strings.Fields(input)
	index := 0

	result, err := parsePrefix(tokens, &index, true)
	if err != nil {
		return "", err
	}
	return result, nil
}

func parsePrefix(tokens []string, index *int, topLevel bool) (string, error) {
	if *index >= len(tokens) {
		return "", fmt.Errorf("invalid expression")
	}

	token := tokens[*index]
	*index++

	if isOperator(token) {
		operands := []string{}

		for i := 0; i < 2; i++ {
			operand, err := parsePrefix(tokens, index, false)
			if err != nil {
				return "", err
			}
			operands = append(operands, operand)
		}

		if topLevel {
			for *index < len(tokens) {
				operand, err := parsePrefix(tokens, index, false)
				if err != nil {
					return "", err
				}
				operands = append(operands, operand)
			}
		}

		return fmt.Sprintf("(%s %s)", tokenToLisp(token), strings.Join(operands, " ")), nil
	}

	return token, nil
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}

func tokenToLisp(token string) string {
	if token == "^" {
		return "pow"
	}
	return token
}
