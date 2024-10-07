package main

import (
	"fmt"
	"strings"
)

var (
	words   = []string{"never", "going", "to", "give", "you", "up"}
	passage = "we are no strangers to love you know the rules and so do i a full commitment is what i am thinking of you would not get this from any other guy i just want to tell you how i am feeling got to make you understand never going to give you up never going to let you down never going to run around and desert you never going to make you cry never going to say goodbye never going to tell a lie and hurt you we have known each other for so long your heart has been aching but you are too shy to say it inside we both know what has been going on we know the game and we are going to play it and if you ask me how i am feeling do not tell me you are too blind to see never going to give you up never going to let you down never going to run around and desert you never going to make you cry never going to say goodbye never going to tell a lie and hurt you never going to give you up never going to let you down never going to run around and desert you never going to make you cry never going to say goodbye never going to tell a lie and hurt you we have known each other for so long your heart has been aching but you are too shy to say it inside we both know what has been going on we know the game and we are going to play it i just want to tell you how i am feeling got to make you understand never going to give you up never going to let you down never going to run around and desert you never going to make you cry never going to say goodbye never going to tell a lie and hurt you never going to give you up never going to let you down never going to run around and desert you never going to make you cry never going to say goodbye never going to tell a lie and hurt you never going to give you up never going to let you down never going to run around and desert you never going to make you cry never going to say goodbye never going to tell a lie and hurt you"
)

func main() {
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word] = 0
	}

	passageWords := strings.Fields(passage)
	for _, word := range passageWords {
		if _, ok := wordCounts[word]; ok {
			wordCounts[word]++
		}
	}

	for word, count := range wordCounts {
		fmt.Println("Found", count, "counts of", word)
	}
}
