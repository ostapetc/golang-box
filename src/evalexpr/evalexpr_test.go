package evalexpr

import (
	"fmt"
	"testing"
)

func TestEvalExpression(t *testing.T) {
	tests := []struct{
		expr   string
		expect int32
	}{
		//{"(2+1)*(5+5)", 30},
		//{"(2*5)+(2*5)", 20},
		//{"(2+5)*(2+5)", 49},
		{"(2+5)*(2+5)+(2*1)", 51},
	}

	for _, test := range tests {
		result := EvalExpression(test.expr)

		if result != test.expect {
			t.Errorf("For expr '%s', expected '%d' got '%d'", test.expr, test.expect, result)
		} else {
			fmt.Println("Success", test.expr)
		}
	}
}
