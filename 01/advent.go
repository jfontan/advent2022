package advent01

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

func read(r io.Reader) ([]int, error) {
	s := bufio.NewScanner(r)

	var values []int
	var value int
	var valid = false
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if l == "" {
			if valid {
				values = append(values, value)
			}
			valid = false
			value = 0
			continue
		}

		valid = true
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		value += i
	}

	if s.Err() != nil {
		return nil, s.Err()
	}

	if valid {
		values = append(values, value)
	}

	return values, nil
}

func max(values []int) (int, int) {
	var pos = 0
	var max = 0

	for i, v := range values {
		if v > max {
			pos = i
			max = v
		}
	}

	return pos, max
}

func max3sum(values []int) int {
	var max = []int{0, 0, 0}

	for _, v := range values {
		if v > max[2] {
			max[2] = v
			sort.Sort(sort.Reverse(sort.IntSlice(max)))
		}
	}

	return max[0] + max[1] + max[2]
}
