package wordgen

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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
	err := godotenv.Load("../secrets.env")
	if err != nil {
		log.Fatal("Error in loading .env file in app")
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

func GetAndReturnWordForCount(characterCount int) (*WordResponse, error){
	localWord, err := getRandomWord(characterCount)
	if err != nil{return nil,err}
	//make external api call here
	wordObj, err := callDictionary(localWord)
	if err != nil{return nil,err}
	return wordObj,nil
}

func callDictionary(word string) (*WordResponse,error){
	var wordResp []WordObject
	log.Println("word checked is ", word)
	res, err := http.Get(dictionary_api_base + strings.TrimSpace(word))
	if err != nil{
		return nil, errors.New(fmt.Sprintf("call to dictionary api failed due to %v",err))
	}
	buffer,_ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(buffer, &wordResp)
	if err != nil {
		fmt.Printf("Error from sending word %s is %v\n", word, err)
		return nil,errors.New(fmt.Sprintf("could not find definition for word %s",word))
	}
	return &WordResponse{
		WordText: word,
		WordData: wordResp,
	}, nil
}