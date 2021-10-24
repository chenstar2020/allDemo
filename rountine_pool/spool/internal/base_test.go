package internal

import (
	"fmt"
	"testing"
)

func TestMathDecimal(t *testing.T) {
	f := 3.12345
	new := fmt.Sprintf("%0.2f", f)
	fmt.Println(new)
}
