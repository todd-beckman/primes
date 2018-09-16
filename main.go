package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/todd-beckman/prime/primes"
	"github.com/todd-beckman/prime/test_utilities"
)

const (
	K = 6
)

func main() {
	args := os.Args
	if len(args) != 2 && len(args) != 3 {
		panic("Expected args n and optionally k, both integers")
	}

	var err error
	var n int

	if n, err = strconv.Atoi(args[1]); err != nil {
		panic("Expected integer n as the first arg")
	}

	k := K
	if len(args) == 3 {
		if k, err = strconv.Atoi(args[2]); err != nil || k >= 10000 {
			panic("Expected optional integer k<10000 as the second arg")
		}
	}

	primeCheck := func() bool {
		return primes.SimpleCheckIfPrime(n, k)
	}

	isPrime, elapsed := timing.GetBoolWithTiming(primeCheck)

	if isPrime {
		fmt.Println(fmt.Sprintf("%d is prime", n))
	} else {
		fmt.Println(fmt.Sprintf("%d is not prime", n))
	}

	if elapsed.Nanoseconds() > 2 {
		fmt.Println(fmt.Sprintf("Took %s", elapsed.String()))
	}
}
