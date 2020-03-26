package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	num := flag.Int("n", 4, "number of words")
	min := flag.Int("min", 3, "min length of words")
	max := flag.Int("max", 8, "max length of words")
	useCommonWords := flag.Bool("common", true, "use common words")
	flag.Parse()

	validWord := func(s string) bool { return len(s) > *min && len(s) < *max }

	var words []string
	if *useCommonWords {
		words = filterCommonWords(validWord)
	} else {
		words = loadSuitableWords(validWord)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	passwords := []string{}
	for i := 0; i < *num; i++ {
		word := words[r.Intn(len(words))]
		passwords = append(passwords, strings.ToLower(word))
	}
	password := strings.Join(passwords, "-")
	fmt.Println(password)
}

func loadSuitableWords(p func(string) bool) []string {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	words := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if p(word) {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

func filterCommonWords(p func(string) bool) []string {
	words := make([]string, 0)
	for _, word := range commonWords {
		if p(word) {
			words = append(words, word)
		}
	}

	return words
}
