package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"bufio"
)

func generateGenome(length int) {
	//generate random string
	buffer := make([]byte, length)
	for i := range buffer {
		buffer[i] = ALPHABET[rand.Intn(len(ALPHABET))]
	}
	//save as input.txt file
	err := ioutil.WriteFile("input.txt", buffer, 0644)
	if err != nil{
		panic(err)
	}
}

func generateShortReads(){
	// shortMin	  : minimum size of short read
	// shortMax	  : maximum size of short read
	// gapMax 	  : maximum overlapping size between two short reads
	shortMin, shortMax, gapMax := 100, 300, 10

	//open file
	f , err := os.Open("input.txt")
	if err != nil{
		fmt.Println("reading txt file error")
		panic(err)
	}
	//read original genome sequencing
	b := make([]byte, LENGTH)
	n, err := f.Read(b)
	fmt.Printf("read %d bytes\n", n)

	//temporal storage for short reads
	buffer := []string{}
	lastEnd := 0 //index of last short read's end

	// pick parts and append to buffer
	for lastEnd < LENGTH {

		// make randomly overlap between gapMax range among short reads
		lastEnd -= rand.Intn(gapMax)

		// random number between 100 ~ 300
		randLength := rand.Intn(shortMax - shortMin) + shortMin
		end := lastEnd + randLength

		tmp := string(b[Max(lastEnd, 0):Min(end, LENGTH-1)])

		buffer = append(buffer, tmp)
		lastEnd = end
	}
	// Save as ORIGIN string
	ORIGIN = strings.Join(buffer, "")

	// Shuffle short reads slice
	shuffleSlice(buffer)

	//store as a file shorts.txt
	f2, err := os.Create("shorts.txt")
	if err != nil{
		panic(err)
	}

	w := bufio.NewWriter(f2)
	for i := range buffer {
		_, err = w.WriteString(buffer[i] + "\n")
		if err != nil{
			fmt.Print(err); panic(err)
		}
	}
}

func shuffleSlice(s []string) {
	for i := len(s) - 1; i > 0 ; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

