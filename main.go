package main

import (
	"time"
	"math/rand"
	"fmt"
)
//length of original dna sample
const LENGTH int = 10000000
var (
	ALPHABET = []byte("ACGT")
	ORIGIN string
	OUTPUT [LENGTH]byte
)

func init(){
	//setting random seed
	rand.Seed(time.Now().UnixNano())

	// Genome Sequence generator
	generateGenome(LENGTH)
	// Short reads generator
	generateShortReads()
}

func main(){
	// Execution time check set up
	start := time.Now()

	//resequence("shorts.txt")
	concurrentResequence("shorts.txt")

	// Execution time catch
	elapsed := time.Since(start)
	// Execution time print
	fmt.Printf("took %s \n", elapsed)
}
