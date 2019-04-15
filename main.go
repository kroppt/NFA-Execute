package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseEdge(m []map[rune]*set, line string) {
	n := len(m)
	strs := strings.Split(line, " ")
	if len(strs) != 3 {
		fmt.Fprintf(os.Stderr, "error parsing edge \"%s\"\n", line)
		os.Exit(1)
	}
	n1, err := strconv.Atoi(strs[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing edge index \"%s\"\n", strs[0])
		os.Exit(1)
	}
	if n1 > n {
		fmt.Fprintf(os.Stderr, "error parsing edge: index %d out of bounds \n", n1)
		os.Exit(1)
	}
	n2, err := strconv.Atoi(strs[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing edge index \"%s\"\n", strs[1])
		os.Exit(1)
	}
	if n2 > n {
		fmt.Fprintf(os.Stderr, "error parsing edge: index %d out of bounds \n", n2)
		os.Exit(1)
	}
	if len([]rune(strs[2])) != 1 {
		fmt.Fprintf(os.Stderr, "error parsing edge character \"%s\"\n", strs[2])
		os.Exit(1)
	}
	r := []rune(strs[2])[0]
	m[n1][r] = newSet(n2)
}

func main() {
	// load NFA
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	in := scan.Text()
	n, err := strconv.Atoi(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	var trans = make([]map[rune]*set, n)
	for i := range trans {
		trans[i] = make(map[rune]*set)
	}
	for scan.Scan() {
		in = scan.Text()
		parseEdge(trans, in)
	}
}
