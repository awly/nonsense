package main

import (
	"github.com/captaincronos/nonsense"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	DEFW = 100
	MAXW = 1000
)

var aliceChains nonsense.Chain

func main() {
	rand.Seed(time.Now().Unix())

	inf, err := os.Open("alice.txt")
	if err != nil {
		log.Println(err)
		return
	}
	aliceChains, err = nonsense.Build(inf)
	if err != nil {
		log.Println(err)
		return
	}

	http.HandleFunc("/", handler)
	if err = http.ListenAndServe(":8081", nil); err != nil {
		log.Println(err)
	}
}

func handler(rw http.ResponseWriter, req *http.Request) {
	c := aliceChains
	if req.Body != nil && req.ContentLength > 0 {
		var err error
		c, err = nonsense.Build(req.Body)
		if err != nil {
			log.Println(err)
			return
		}
	}

	max, err := strconv.Atoi(req.URL.Query().Get("len"))
	if err != nil {
		max = DEFW
	}
	if max > MAXW {
		max = MAXW
	}

	if err = c.Gen(rw, max); err != nil {
		log.Println(err)
		return
	}
}
