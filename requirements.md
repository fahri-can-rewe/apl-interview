# Golang - Word-Pair Anagram Checker

Our customer has a Word-Pair API that returns two words in JSON format.
Recently he has noticed that some of the word pairs have very similar letters and suspects that some of them might be anagrams of each other.

He has asked you to help him identify these anagrams.

## Task
Create a Golang application that connects to the Word-Pair API, processes the response, and checks if the two words are anagrams of each other.
The Word-Pair API is documented at: https://interview.sowula.at/docs

## Requirements

### 1. Run the App
```sh
go run main.go
```

The app should:
- Call the API (https://interview.sowula.at/word-pair)
- Parse the JSON response
- Check if the two words are anagrams
- Print the two words and whether they are anagrams to stdout

Example output:
```
Word 1: listen
Word 2: silent
Are Anagrams: true
```

### 2. Structure
- Add your own packages/files/interfaces (client, model, solver, …)
- Keep code testable

### 3.	Configuration
In a later stage, this application will be deployed in different environments and integrated more closely with other services.
Therefore, the base URL of the Word-Pair API should be configurable via a command line argument:

- go run main.go --apiBaseUrl https://host:port

## Notes
- You are free to import and library for this task.
- Write code as if it could later run in production.
- No full test coverage required now, but keep testability in mind.
- Think about how you would later containerize/deploy this (Docker, K8s).

## Allowed tools and AI usage
You are free to use any IDE, build tool, or libraries you prefer.

You may also use ChatGPT or other AI tools to help you, **except** for generating a solution regarding the anagram logic itself.

The anagram detection logic must be your own work — please don't ask:
- Write me a function that checks if two words are anagrams.
- How do I check if two words are anagrams?


