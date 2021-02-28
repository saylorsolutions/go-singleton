package first_attempt

import (
	"testing"

	testify "github.com/stretchr/testify/require"
)

func TestGetSingleton(t *testing.T) {
	assert := testify.New(t)

	initial := GetSingleton()
	assert.Equal(42, initial, "State wasn't initialized as expected")

	initial += 5
	gotAgain := GetSingleton()
	assert.Equal(42, gotAgain, "State was changed!")

	// State is able to be changed to a new "instance", doesn't meet requirements.
	State = 43
	assert.Equal(43, GetSingleton())
}
