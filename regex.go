package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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

	myrand := random(0, 3)

	output := responses[myrand]
	return (output)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Talk to Eliza:")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	fmt.Println(ElizaResponse(text))

} //main
