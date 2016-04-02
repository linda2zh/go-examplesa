
package main

import "fmt"

const (
	size    = 1000
	verbose = false
)

type LL interface {
	head() int
	tail() LL
}

type baseLL struct {
	fst int
}

func (s baseLL) head() int { return s.fst }
func (s baseLL) tail() LL  { return baseLL{1 + s.fst} }

type filtLL struct {
	fst     int
	preFilt LL
	p       int
}

func (s filtLL) head() int { return s.fst }
func (s filtLL) tail() LL  { return filter(s.preFilt, s.p) }

func filter(in LL, p int) LL {
	newHead := in.head()
	newRest := in.tail()
	for newHead%p == 0 {
		newHead = newRest.head()
		newRest = newRest.tail()
	}
	return filtLL{newHead, newRest, p}
}

type recursLL struct {
	fst     int
	preProc LL
}

func (s recursLL) head() int { return s.fst }
func (s recursLL) tail() LL {
	newTail := filter(s.preProc, s.fst)
	return recursLL{newTail.head(), newTail.tail()}
}

func sieve() LL {
	l2 := baseLL{2}
	return recursLL{l2.head(),
		filter(l2.tail(), l2.head())}
}

func main() {
	s := sieve()
	var r [size - 1]bool
	for _ = range r {
		if verbose {
			fmt.Println(s.head())
		}
		s = s.tail()
	}
	fmt.Println(size, "th prime is ", s.head())
}
