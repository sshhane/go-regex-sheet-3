package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
)

// random genorates random number
// http://golangcookbook.blogspot.ie/2012/11/generate-random-number-in-given-range.html
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// ElizaResponse takes input and picks output
func ElizaResponse(input string) string {

	//saved responses array
	responses := [3]string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	fatherResp := [3]string{
		"Why don’t you tell me more about your father?",
		"Lets talk about your father",
		"Your father?  Please tell me more",
	}

	// List the reflections.
	reflections := [][]string{
		{"your", "my"},
		{"you", "I"},
		{"me", "you"},
	}

	// https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)

	// capture "I'm, Im, Iam, I am"
	capture := regexp.MustCompile(`\bI'?\s*a?m\b([^.?!]*)`)

	myrand := random(0, 3)
	var output string

	//check, using a regular expression, if the input contains the word “father”.
	if father, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); father {
		output = fatherResp[myrand]
	} else {
		if match := capture.MatchString(input); match {

			output = capture.ReplaceAllString(input, "How do you know you are$1?")

		} else {
			output = responses[myrand]
		}
	}

	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	// Put the tokens back together.
	// output = strings.Join(tokens, ``)

	return (output)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Talk to Eliza:")
	text, _ := reader.ReadString('\n')
	// fmt.Println(text)

	fmt.Println(ElizaResponse(text))

} //main
