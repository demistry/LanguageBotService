package main

import (
	"LanguageBotService/grpcUtil"
	"LanguageBotService/wordgen"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT       = "50051"
	TcpAddress = "0.0.0.0:"
)

func Init(){
	lis, err := net.Listen("tcp", TcpAddress + PORT)
	if err != nil{
		log.Fatal("Could not initialize listener due to ", err.Error())
	}
	fmt.Println("Running server on port: 50051")
	s := grpc.NewServer()
	grpcUtil.RegisterBotServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}

type server struct {

}

func (s server) GenerateWord(ctx context.Context, request *grpcUtil.WordRequest) (*grpcUtil.WordResponse, error) {
	fmt.Println("Calling generate word...")
	charCount := request.GetWordCount()
	word,_ := wordgen.GetRandomWord(int(charCount))
	return &grpcUtil.WordResponse{
		Word:  word       ,
		PartOfSpeech: "Part of speech is i dunno",
		Definition:   []string{},
		Example:      []string{},
	}, nil
}


