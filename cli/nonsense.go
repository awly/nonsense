package main

import (
	"fmt"
	"github.com/captaincronos/nonsense"
	"math/rand"
	"os"
	"time"
)

const (
	MAXGEN = 100 // max generated words
)

func main() {
	rand.Seed(time.Now().Unix())
	chains, err := nonsense.Build(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = chains.Gen(os.Stdout, MAXGEN); err != nil {
		fmt.Println(err)
	}
}
