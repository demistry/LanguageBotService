package main

import (
	"LanguageBotService/grpcUtil"
	"LanguageBotService/wordgen"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)


func Init(){
	PORT := os.Getenv("PORT")
	TcpAddress := os.Getenv("TCP_ADDRESS")

	lis, err := net.Listen("tcp", TcpAddress + PORT)
	if err != nil{
		log.Fatal("Could not initialize listener due to ", err.Error())
	}
	fmt.Println("Running server on port:",PORT)
	s := grpc.NewServer()
	grpcUtil.RegisterBotServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}

type server struct {

}

func (s server) GenerateWord(ctx context.Context, request *grpcUtil.WordRequest) (*grpcUtil.WordResponse, error) {
	charCount := request.GetWordCount()
	wordRes, err := wordgen.GetAndReturnWordForCount(int(charCount))
	if err != nil{
		return &grpcUtil.WordResponse{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		}, nil
	}
	var def []string
	var example [] string
	var partsOfSpeech [] string
	for _,word := range wordRes.WordData{
		for _,meaning := range word.Meaning {
			fmt.Println("Meaning POS here is ", meaning.PartOfSpeech)
			partsOfSpeech = append(partsOfSpeech, meaning.PartOfSpeech)
			for _,defin := range meaning.Definitions{
				if defin.Definition != ""{
					def = append(def,defin.Definition)
				}
				if defin.Example != ""{
					example = append(example,defin.Example)
				}
			}
		}
	}
	resp := &grpcUtil.WordResponse{
		Status:  true,
		Message: "Successfully received generated word",
		Data:    &grpcUtil.WordData{
			Word:         wordRes.WordText,
			PartOfSpeech: partsOfSpeech,
			Definition:   def,
			Example:      example,
		},
	}
	return resp, nil
}


