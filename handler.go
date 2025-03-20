package lab2

import (
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Input)
	if err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}
	expression := strings.TrimSpace(string(data))
	if expression == "" {
		return fmt.Errorf("empty input")
	}

	result, err := PrefixToLisp(expression)
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(result))
	return err
}

