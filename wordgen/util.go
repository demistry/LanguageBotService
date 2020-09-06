package wordgen

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type WordObject struct{
	Word string
	Definition string
}
func shuffleWords(){
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(totalWords), func(i, j int) { totalWords[i], totalWords[j] = totalWords[j], totalWords[i] })
}

//Generates and returns a random word based on character count
func getRandomWord(characterCount int) (string,error){
	shuffleWords()
	lengthOfTotalWords := len(totalWords)
	var idx = -1
	for i := range totalWords{
		if len(totalWords[i]) == characterCount{
			idx = i
			break
		}
	}
	if idx >= 0{
		return strings.TrimSpace(totalWords[idx]), nil
	} else {
		return "",errors.New( fmt.Sprintf("Could not find word with character length %d in word bank", characterCount))
	}
}


func parseWords(data []byte) []string{
	var word string
	var totalWords []string
	if string(data[len(data)-1]) != NewLineChar{ //checks if the end of the txt file has no new line and appends one to it
		data = append(data,ByteNewline...)
	}
	for _,deet := range data{
		character := string(deet)
		if character != NewLineChar { //Check for new line
			word += character
		}else{
			totalWords = append(totalWords,word)
			word = ""
		}
	}
	return totalWords
}