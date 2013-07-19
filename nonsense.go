// package nonsense generates random text using Markov chains
// uses 2-word prefixes
package nonsense

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

const (
	npref = 2   // prefix length
	naw   = " " // not a word
)

type Chain map[prefix][]string

// Build: read space-separated strings from in and Build markov chains from them
func Build(in io.Reader) (Chain, error) {
	var w string       // word in input
	var p prefix       // current prefix
	for i := range p { // initialize "empty" prefix
		p[i] = naw
	}
	c := make(Chain)

	s := bufio.NewScanner(in)
	for s.Scan() {
		// split line into space-separated tokens and add them
		// append \n to preserve some structure
		for _, w = range strings.Split(s.Text()+"\n", " ") {
			c[p] = append(c[p], w)
			p.advance(w)
		}
	}
	c[p] = append(c[p], naw)
	if err := s.Err(); err != nil {
		return nil, err
	}

	return c, nil
}

// Gen: using given c, write random text to out, space-separated
// user is responsible for seeding PRNG from math/rand
func (c Chain) Gen(out io.Writer, max int) error {
	var next string    // next word to be written
	var p prefix       // current prefix
	for i := range p { // initialize "empty" prefix
		p[i] = naw
	}
	var err error
	for i := 0; i < max; i++ {
		next = c[p][rand.Intn(len(c[p]))] // get random suffix for given prefix
		if next == naw {
			break
		}
		if _, err = fmt.Fprint(out, next, " "); err != nil {
			return err
		}
		p.advance(next)
	}
	return nil
}

type prefix [npref]string

func (p *prefix) advance(next string) {
	for i := 0; i < len(p)-1; i++ { // move all elements backwards by 1
		p[i] = p[i+1]
	}
	p[len(p)-1] = next // promote suffix to last prefix element
}
