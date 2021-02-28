package second_attempt

import (
	"testing"

	testify "github.com/stretchr/testify/require"
)

func TestGetSingleton(t *testing.T) {
	assert := testify.New(t)

	initial := GetSingleton()
	assert.Equal(42, initial.GetValue(), "State wasn't initialized as expected")

	gotAgain := GetSingleton()
	initial.SetValue(47)
	assert.Equal(47, gotAgain.GetValue(), "State was NOT changed!")

	// State value is able to be manipulated directly, may bypass validation checks in methods.
	gotAgain.value = 42
	assert.Equal(42, gotAgain.GetValue())

	// State variable is able to be changed to a new instance from within the same package, doesn't meet requirements.
	state = Singleton{
		value: 43,
	}
	assert.Equal(43, GetSingleton().GetValue())
}
