package wordgen

import (
	"log"
	"os"
)

const (
	NewLineChar = "\n"
	FileName = "wordlist.txt"
)

var (
	ByteNewline = []byte(NewLineChar)
	totalWords  []string
)

func Init() {
	file,err := os.Open(FileName)
	if err != nil{
		log.Fatalf("Error in opening file %v", err)
	}
	defer  file.Close()
	fileInfo,_ := os.Stat(FileName)
	data := make([]byte,int(fileInfo.Size()))
	_,err = file.Read(data)
	totalWords = parseWords(data)
}

func GetAndReturnWordForCount(characterCount int) (*WordObject, error){
	localWord, err := getRandomWord(characterCount)
	if err != nil{return nil,err}
	//make external api call here
	return &WordObject{
		Word:       localWord,
		Definition: localWord + " Definition",
	}, nil
}