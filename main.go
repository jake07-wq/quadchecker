package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
				res.WriteByte('o')
			} else if i == 0 || i == y-1 {
				res.WriteByte('-')
			} else if j == 0 || j == x-1 {
				res.WriteByte('|')
			} else {
				res.WriteByte(' ')
			}
		}
		res.WriteByte('\n')
	}
	return res.String()
}

func quadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				res.WriteByte('/')
			} else if i == 0 && j == x-1 {
				res.WriteByte('\\')
			} else if i == y-1 && j == 0 {
				res.WriteByte('\\')
			} else if i == y-1 && j == x-1 {
				res.WriteByte('/')
			} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
				res.WriteByte('*')
			} else {
				res.WriteByte(' ')
			}
		}
		res.WriteByte('\n')
	}
	return res.String()
}

func quadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && (j == 0 || j == x-1) {
				res.WriteByte('A')
			} else if i > 0 && (j == 0 || j == x-1) && i == y-1 {
				res.WriteByte('C')
			} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
				res.WriteByte('B')
			} else {
				res.WriteByte(' ')
			}
		}
		res.WriteByte('\n')
	}
	return res.String()
}

func quadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if j == 0 && (i == 0 || i == y-1) {
				res.WriteByte('A')
			} else if j > 0 && (i == 0 || i == y-1) && j == x-1 {
				res.WriteByte('C')
			} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
				res.WriteByte('B')
			} else {
				res.WriteByte(' ')
			}
		}
		res.WriteByte('\n')
	}
	return res.String()
}

func quadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == y-1 && j == x-1 && i > 0 && j > 0) {
				res.WriteByte('A')
			} else if (i == 0 && j == x-1) || (i == y-1 && j == 0) {
				res.WriteByte('C')
			} else if i == 0 || i == y-1 || j == 0 || j == x-1 {
				res.WriteByte('B')
			} else {
				res.WriteByte(' ')
			}
		}
		res.WriteByte('\n')
	}
	return res.String()
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if len(input) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	y := len(input)
	x := len(input[0])
	fullStr := strings.Join(input, "\n") + "\n"

	type quadFunc struct {
		name string
		f    func(int, int) string
	}

	quads := []quadFunc{
		{"quadA", quadA},
		{"quadB", quadB},
		{"quadC", quadC},
		{"quadD", quadD},
		{"quadE", quadE},
	}

	var matches []string
	for _, q := range quads {
		if q.f(x, y) == fullStr {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", q.name, x, y))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}
