package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	NPREF  = 2   // prefix length
	MAXGEN = 100 // max generated words
	NAW    = " " // not a word
)

type prefix [NPREF]string

func (p *prefix) advance(next string) {
	for i := 0; i < len(p)-1; i++ { // move all elements backwards by 1
		p[i] = p[i+1]
	}
	p[len(p)-1] = next // promote suffix to last prefix element
}

func main() {
	rand.Seed(time.Now().Unix())
	chains := make(map[prefix][]string)
	build(chains, os.Stdin)
	generate(chains, os.Stdout)
}

// build: read space-separated strings from in and build markov chains from them
func build(c map[prefix][]string, in io.Reader) {
	var w string       // word in input
	var p prefix       // current prefix
	for i := range p { // initialize "empty" prefix
		p[i] = NAW
	}

	s := bufio.NewScanner(in)
	for s.Scan() {
		// split line into space-separated tokens and add them
		// append \n to preserve some structure
		for _, w = range strings.Split(s.Text()+"\n", " ") {
			c[p] = append(c[p], w)
			p.advance(w)
		}
	}
	if err := s.Err(); err != nil {
		fmt.Println(err)
	}
}

// generate: using given c, write random text to out, space-separated
func generate(c map[prefix][]string, out io.Writer) {
	var next string    // next word to be written
	var p prefix       // current prefix
	for i := range p { // initialize "empty" prefix
		p[i] = NAW
	}
	var err error
	for i := 0; i < MAXGEN; i++ {
		next = c[p][rand.Intn(len(c[p]))] // get random suffix for given prefix
		if next == NAW {
			break
		}
		if _, err = fmt.Fprint(out, next, " "); err != nil {
			fmt.Println(err)
			break
		}
		p.advance(next)
	}
}
