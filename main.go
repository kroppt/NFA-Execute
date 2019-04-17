package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	. "github.com/kroppt/IntSet"
)

var flagin string
var pstate bool

func init() {
	flag.StringVar(&flagin, "in", "stdin", "file for input string")
	flag.BoolVar(&pstate, "print", false, "prints intermediate states")
}

func parseEdge(trans []map[rune]Set, line string) {
	n := len(trans)
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
	if n1 >= n {
		fmt.Fprintf(os.Stderr, "error parsing edge: index %d out of bounds \n", n1)
		os.Exit(1)
	}
	n2, err := strconv.Atoi(strs[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing edge index \"%s\"\n", strs[1])
		os.Exit(1)
	}
	if n2 >= n {
		fmt.Fprintf(os.Stderr, "error parsing edge: index %d out of bounds \n", n2)
		os.Exit(1)
	}
	if len([]rune(strs[2])) != 1 {
		fmt.Fprintf(os.Stderr, "error parsing edge character \"%s\"\n", strs[2])
		os.Exit(1)
	}
	r := []rune(strs[2])[0]
	if _, ok := trans[n1][r]; ok {
		trans[n1][r].Add(n2)
	} else {
		trans[n1][r] = NewSetInit(n2)
	}
}

func εClosure(trans []map[rune]Set, s Set) (ns Set) {
	// iterate until no change
	ns = s.Copy()
	var Δs Set
	iterf := func(i int) {
		if s, ok := trans[i]['ε']; ok {
			Δs.Union(s)
		}
	}
	Δb := true
	for Δb {
		Δs = NewSet()
		ns.Range(iterf)
		Δb = ns.Union(Δs)
	}
	return ns
}

func rmEmpty(strs []string) (out []string) {
	out = make([]string, 0)
	for _, s := range strs {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	flag.Parse()
	// load NFA
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "there must be at least 1 argument for the input file")
		os.Exit(1)
	}
	buf, err := ioutil.ReadFile(args[len(args)-1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fstr := strings.Replace(string(buf), "\r", "", -1)
	// remove carriage returns
	strs := strings.Split(fstr, "\n")
	strs = rmEmpty(strs)
	n, err := strconv.Atoi(strs[0])
	if n < 2 {
		fmt.Fprintln(os.Stderr, "there must be at least 2 nodes")
		os.Exit(1)
	}
	accept := make([]bool, n)
	states := strings.Split(strs[1], " ")
	if len(states) <= 0 {
		fmt.Fprintln(os.Stderr, "there must be at least 1 accepting state")
		os.Exit(1)
	}
	for _, str := range states {
		i, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading error states")
			os.Exit(1)
		}
		if i >= n {
			fmt.Fprintf(os.Stderr, "accepting state %d is outside bounds\n", i)
			os.Exit(1)
		}
		accept[i] = true
	}
	var trans = make([]map[rune]Set, n)
	for i := range trans {
		trans[i] = make(map[rune]Set)
		trans[i]['ε'] = NewSetInit(i)
	}
	for _, str := range strs[2:] {
		parseEdge(trans, str)
	}
	// read input
	infile := os.Stdin
	if flagin != "stdin" {
		if infile, err = os.Open(flagin); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		defer infile.Close()
	}
	input := bufio.NewReader(infile)
	r, _, err := input.ReadRune()
	// begin algorithm
	var oldState Set
	initState, _ := trans[0]['ε']
	initState.Add(0)
	currState := εClosure(trans, initState)
	if pstate {
		fmt.Println(currState.Print())
	}
	for err == nil {
		oldState = currState
		currState = NewSet()
		oldState.Range(func(i int) {
			if s, ok := trans[i][r]; ok {
				currState.Union(s)
			}
		})
		currState = εClosure(trans, currState)
		if currState.IsEmpty() {
			fmt.Fprintln(os.Stderr, "input not accepted")
			os.Exit(1)
		}
		if pstate {
			fmt.Println(currState.Print())
		}
		r, _, err = input.ReadRune()
	}
	currState.Range(func(i int) {
		if accept[i] {
			fmt.Println("input accepted")
			os.Exit(0)
		}
	})
	fmt.Fprintln(os.Stderr, "input not accepted")
	os.Exit(1)
}
