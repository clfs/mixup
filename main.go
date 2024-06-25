package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
)

func main() {
	log.SetPrefix("mixup: ")
	log.SetFlags(0)

	var lines []string

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})

	for _, line := range lines {
		fmt.Println(line)
	}
}
