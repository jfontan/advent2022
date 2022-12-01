package advent01

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

const testData = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func Test01Example(t *testing.T) {
	r := bytes.NewBufferString(testData)

	values, err := read(r)
	require.NoError(t, err)

	pos, max := max(values)
	require.Equal(t, 3, pos)
	require.Equal(t, 24000, max)
}

func Test01Solution(t *testing.T) {
	f, err := os.Open("data.txt")
	require.NoError(t, err)
	defer f.Close()

	values, err := read(f)
	require.NoError(t, err)

	pos, max := max(values)
	fmt.Printf("pos: %d, max: %d\n", pos+1, max)
}

func Test02Example(t *testing.T) {
	r := bytes.NewBufferString(testData)

	values, err := read(r)
	require.NoError(t, err)

	spew.Dump(values)

	sum := max3sum(values)
	require.Equal(t, 45000, sum)
}

func Test02Solution(t *testing.T) {
	f, err := os.Open("data.txt")
	require.NoError(t, err)
	defer f.Close()

	values, err := read(f)
	require.NoError(t, err)

	sum := max3sum(values)
	fmt.Printf("sum: %d\n", sum)
}
