package pad

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightPad(t *testing.T) {
	testString := "meheh"
	padded := Right(testString, 33, "X")

	assert.Equal(t, len(padded), 33+len(testString))
}
