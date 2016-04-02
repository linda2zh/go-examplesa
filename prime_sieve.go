package main
// Implements an algorithm similar to the Haskell
// let f (p:as) = p : f [a|a<-as, mod a p > 0] in take 10 $ f [2..]

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
	baseLL
	preFilt LL
	p       int
}

func (s filtLL) tail() LL { return filter(s.preFilt, s.p) }

func filter(in LL, p int) LL {
	newHead := in.head()
	newRest := in.tail()
	for newHead%p == 0 {
		newHead = newRest.head()
		newRest = newRest.tail()
	}
	return filtLL{baseLL{newHead}, newRest, p}
}

type recursLL struct {
	baseLL
	preProc LL
}

func (s recursLL) tail() LL {
	newTail := filter(s.preProc, s.head())
	return recursLL{baseLL{newTail.head()}, newTail.tail()}
}

func sieve() LL {
	l2 := baseLL{2}
	return recursLL{baseLL{l2.head()},
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
