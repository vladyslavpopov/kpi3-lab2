package lab2

import (
  "fmt"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestPrefixToLisp(t *testing.T) {
  res, err := PrefixToLisp("+ * 2 3 4")
  if assert.Nil(t, err) {
    assert.Equal(t, "(+ (* 2 3) 4)", res)
  }
}

func TestPrefixToLispMoretOperands(t *testing.T) {
  res, err := PrefixToLisp("* + 2 3 - 8 4 ^ 5 2 / 10 5")
  if assert.Nil(t, err) {
    assert.Equal(t, "(* (+ 2 3) (- 8 4) (pow 5 2) (/ 10 5))", res)
  }
}

func TestInvalidInput(t *testing.T) {
  _, err := PrefixToLisp("")
  assert.NotNil(t, err)

  _, err = PrefixToLisp("+ 5 *")
  assert.NotNil(t, err)
}

func ExamplePrefixToLisp() {
  res, _ := PrefixToLisp("+ * 2 3 - 8 4")
  fmt.Print(res)
  // Output: (+ (* 2 3) (- 8 4))
}
