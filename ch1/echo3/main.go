// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func alg1(args []string){
	for index, arg := range args[0:] {
		fmt.Println(index, arg);
	}
}

func alg2(args []string){
    fmt.Println(strings.Join(args[0:], " "))
}

//!+
func main() {
	fmt.Println("Algorithm 1:")
	start := time.Now()
	alg1(os.Args)
	secs := time.Since(start).Seconds()
	fmt.Printf("%f seconds\n", secs)
	fmt.Println("Algorithm 2:")
	start = time.Now()
	alg2(os.Args)
	secs = time.Since(start).Seconds()
	fmt.Printf("%f seconds\n", secs)
}

//!-
