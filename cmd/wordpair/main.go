package main

import (
	"apl-interview/anagram"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

const fiveSecondsTimeout = 5 * time.Second

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()
	if err := run(ctx, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, args []string) error {
	conf, err := parseArgs(args)
	if err != nil {
		return err
	}

	ac, err := buildAPIClient(conf.APIBaseURL)
	if err != nil {
		return err
	}

	wp, err := fetchWordPair(ctx, ac)
	if err != nil {
		return err
	}

	if err := anagram.ValidateWordPair(wp.FirstWord, wp.SecondWord); err != nil {
		return err
	}
	isAnagram := areAnagrams(wp.FirstWord, wp.SecondWord, anagram.FreqMapChecker{})
	fmt.Printf("Word 1: %s\n", wp.FirstWord)
	fmt.Printf("Word 2: %s\n", wp.SecondWord)
	fmt.Printf("Are Anagrams: %v\n", isAnagram)
	return nil
}
