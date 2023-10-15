package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	gpt2 "jarvis/out_bound/grpc/proto/stubs"
	"log"
)

var (
	client gpt2.ChatServiceClient
)

func CreateClient(host string, port string) {
	conn, err := grpc.Dial(host+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("Can not connect to gpt2 server: ", err)
	}

	client = gpt2.NewChatServiceClient(conn)
}

func GetGeneratedText(text string) string {
	req := gpt2.TextRequest{Text: text, MaxLength: 30, NumReturnSequences: 3}
	res, err := client.GenerateText(context.Background(), &req)
	if err != nil {
		log.Printf("Error in calling gpt2", err)
		return "Internal server error"
	}
	return res.GetText()[0]
}
