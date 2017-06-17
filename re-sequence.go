package main

import (
	"os"
	"bufio"
	"fmt"
)

func concurrentResequence(filePath string) {
	//read shorts file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()


	//concurrently read shorts from file
	buffer := make(chan string)

	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			buffer <- scanner.Text()
		}
		close(buffer)
	}()

	var shortsCount = 0
	for b :=  range buffer {
		go insertToSlice(b)
		shortsCount++
		if shortsCount%1000 == 0{
			fmt.Printf(" . %d", shortsCount)
		}
	}
	fmt.Println()
	fmt.Printf("read %d shorts\n", shortsCount)

	////generate output.txt file
	generateOutputFile()
}


func insertToSlice(b string){
	indexes := bruteSearch(ORIGIN, b)
	if indexes != -1 {
		start := indexes
		strByte := []byte(b)
		for i:= start; i < start + len(b); i++{
			if i < LENGTH {
				OUTPUT[i] = strByte[i-start]
			}
		}
	}
}

func resequence(filePath string){
	//read shorts file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//read each line of shorts
	//and store to buffer
	scanner := bufio.NewScanner(file)

	var buffer string

	for scanner.Scan(){
		buffer = scanner.Text()
		i := bruteSearch(ORIGIN, buffer)
		if i != -1 {
			insertToSlice(buffer)
		}
	}

	generateOutputFile()
}




func generateOutputFile() {
	//fmt.Println(string(OUTPUT[:LENGTH]))
	// []byte to string
	output := string(OUTPUT[:LENGTH])
	f, err := os.Create("output.txt")
	if err != nil{
		panic(err)
	}

	w := bufio.NewWriter(f)
	w.WriteString(output)
}