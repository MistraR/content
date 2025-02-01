package main

import (
	"content-manage/api/operate"
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
)

func main() {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := operate.NewAppClient(conn)
	//reply, err := client.CreateContent(context.Background(), &operate.CreateContentRequest{
	//	Content: &operate.Content{
	//		Title:       "Mistra",
	//		VideoUrl:    "https://127.0.0.1:8000",
	//		Author:      "Mistra Author",
	//		Description: "test",
	//	},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("[grpc] CreateContent #{reply} \n", reply)
	reply, err := client.UpdateContent(context.Background(), &operate.UpdateContentRequest{
		Content: &operate.Content{
			Id:          1,
			Title:       "Mistra",
			VideoUrl:    "https://127.0.0.1:8000",
			Author:      "Mistra Author",
			Description: "test33333",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] UpdateContent #{reply} \n", reply)
}
