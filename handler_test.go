package lab2

import (
  "bytes"
  "strings"
  "testing"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
)

func TestComputeHandler_Success(t *testing.T) {
	input := "+ 2 3"
	reader := strings.NewReader(input)
	var buf bytes.Buffer

	handler := &ComputeHandler{
		Input:  reader,
		Output: &buf,
	}

	err := handler.Compute()
	require.NoError(t, err, "compute must successfully process a valid input expression")

	expected := "(+ 2 3)"
	assert.Equal(t, expected, buf.String(), "The result obtained should correspond to the expected one.")
}

func TestComputeHandler_InvalidInput(t *testing.T) {
	reader := strings.NewReader("")
	var buf bytes.Buffer

	handler := &ComputeHandler{
		Input:  reader,
		Output: &buf,
	}

	err := handler.Compute()
	require.Error(t, err, "Compute should return an error on an empty input expression")
}
