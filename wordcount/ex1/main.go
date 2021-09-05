package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) == 1 {
		log.Error("No files to process")
		return
	}

	result := make(map[string]int)

	start := time.Now()
	for _, f := range os.Args[1:] {
		processFile(result, f)
	}

	defer fmt.Printf("Processing took: %v\n", time.Since(start))
	printResult(result)
}

func processFile(result map[string]int, f string) {
	var w string
	r, err := os.Open(f)
	if nil != err {
		log.Warn(err)
		return
	}
	defer r.Close()

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		w = strings.ToLower(sc.Text())
		result[w] = result[w] + 1
	}
}

func printResult(result map[string]int) {
	fmt.Printf("%-10s%s\n", "Count", "Word")
	fmt.Printf("%-10s%s\n", "-----", "----")

	for w, c := range result {
		fmt.Printf("%-10v%s\n", c, w)
	}
}
