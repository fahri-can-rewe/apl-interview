package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	conf, err := parseArgs(args)
	if err != nil {
		return err
	}

	apiClient, err := buildAPIClient(conf.APIBaseURL)
	if err != nil {
		return err
	}

	wp, err := fetchWordPair(apiClient)
	if err != nil {
		return err
	}

	if err := validateWordPair(wp.FirstWord, wp.SecondWord); err != nil {
		return err
	}
	isAnagram := areAnagrams(wp.FirstWord, wp.SecondWord)
	fmt.Printf("Word 1: %s\n", wp.FirstWord)
	fmt.Printf("Word 2: %s\n", wp.SecondWord)
	fmt.Printf("Are Anagrams: %v\n", isAnagram)
	return nil
}
