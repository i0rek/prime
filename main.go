package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
)

func naive(limit int) []int {
	primes := []int{}
	for i := 0; i <= limit; i++ {
		n := float64(i)
		if n == 0 || n == 1 {
			continue
		}
		if n == 2 {
			primes = append(primes, 2)
			continue
		}
		found := false
		for i := 2.0; i < n; i++ {
			if math.Mod(n, i) == 0 {
				found = true
				break
			}
		}
		if !found {
			primes = append(primes, i)
		}
	}
	return primes
}

func sieve(limit int) []int {
	primes := []int{}
	sieb := map[int]bool{}
	for i := 2; i <= limit; i++ {
		sieb[i] = true
	}
	for i := 2; i <= int(math.Sqrt(float64(limit))); i++ {
		for k, v := range sieb {
			if k <= i || v == false {
				continue
			}
			if math.Mod(float64(k), float64(i)) == 0 {
				sieb[k] = false
			}
		}
	}
	for k, v := range sieb {
		if v == true {
			primes = append(primes, k)
		}
	}
	return primes
}

func main() {
	var limit int
	var verbose bool
	var algorithm string
	flag.IntVar(&limit, "limit", 0, "Last number to check.")
	flag.BoolVar(&verbose, "verbose", false, "Wether or not to print the primes.")
	flag.StringVar(&algorithm, "algorithm", "naive", "Which algorithm to use.")
	flag.Parse()
	primes := []int{}
	switch algorithm {
	case "naive":
		primes = naive(limit)
	case "sieve":
		primes = sieve(limit)
	}
	if verbose {
		sort.Ints(primes)
		for _, p := range primes {
			fmt.Println(p)
		}
	}
}