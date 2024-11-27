package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func printFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("open file fail: %s", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	file.Close()
}

func main() {
	printFile("./test.txt")
}
