package advent02

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const testData = `A Y
B X
C Z`

func Test01Example(t *testing.T) {
	r := bytes.NewBufferString(testData)

	s := first(r)

	require.Equal(t, 15, s)
}

func Test01(t *testing.T) {
	r, err := os.Open("data")
	require.NoError(t, err)
	defer r.Close()

	s := first(r)

	println(s)
}

func Test02Example(t *testing.T) {
	r := bytes.NewBufferString(testData)

	s := second(r)

	require.Equal(t, 12, s)
}

func Test02(t *testing.T) {
	r, err := os.Open("data")
	require.NoError(t, err)
	defer r.Close()

	s := second(r)

	println(s)
}
