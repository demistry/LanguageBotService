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
	word, err := wordgen.GetAndReturnWordForCount(int(charCount))
	if err != nil{
		return nil, err
	}
	return &grpcUtil.WordResponse{
		Word:  word.Word       ,
		PartOfSpeech: "Part of speech is i dunno now sha",
		Definition:   []string{word.Definition},
		Example:      []string{},
	}, nil
}


