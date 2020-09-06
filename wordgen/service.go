package wordgen

import (
	"github.com/joho/godotenv"
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
	dictionary_api_base string
)

func Init() {
	err := godotenv.Load("./secrets.env")
	if err != nil {
		log.Fatal("Error in loading .env file in app",err)
	}

	dictionary_api_base = os.Getenv("DICTIONARY_BASE_API")
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
		Word:      dictionary_api_base + localWord,
		Definition: localWord + " Definition",
	}, nil
}