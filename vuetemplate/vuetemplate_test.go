package vuetemplate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVue(t *testing.T) {
	v := Vue{
		VarName: "app",
	}
	assert.Equal(t, "ab", v.String())
}
func TestJSElement(t *testing.T) {
	testCases := []struct {
		Got    JSElement
		Expect string
	}{
		{
			Got: JSElement{
				CONSTANT,
				"var1",
				"val1",
			},
			Expect: `const var1 = "val1";`,
		},
		{
			Got: JSElement{
				VARIABLE,
				"var1",
				"val1",
			},
			Expect: `var var1 = "val1";`,
		},
		{
			Got: JSElement{
				LETSTMT,
				"var1",
				"val1",
			},
			Expect: `let var1 = "val1";`,
		},
	}
	for _, tc := range testCases {
		jse := tc.Got
		assert.Equal(t, tc.Expect, jse.String())
	}
}
func TestComponent(t *testing.T) {
}