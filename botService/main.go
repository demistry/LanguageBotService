package main

import (
	"LanguageBotService/grpcUtil"
	"LanguageBotService/translator"
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
type server struct {

}

func (s server) TranslateNow(ctx context.Context, translate *grpcUtil.Translate) (*grpcUtil.Translate, error) {
	return nil, nil
}

func main(){
	fmt.Println("Hello go 1.15")
	translator.Hi()

	wordgen.Init()
	if word, err := wordgen.GetRandomWord(7); err != nil{
		fmt.Println("Error is ", err)
	} else{
		fmt.Println("Word is", word)
	}

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



