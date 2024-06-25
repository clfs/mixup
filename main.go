package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
)

func main() {
	log.SetPrefix("mixup: ")
	log.SetFlags(0)

	deterministicFlag := flag.Bool("d", false, "enable deterministic mode")
	helpFlag := flag.Bool("h", false, "show help")

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		return
	}

	var lines []string

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fn := func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	}

	if *deterministicFlag {
		rand.New(rand.NewPCG(0, 0)).Shuffle(len(lines), fn)
	} else {
		rand.Shuffle(len(lines), fn)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
