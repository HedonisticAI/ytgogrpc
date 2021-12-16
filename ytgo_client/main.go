package main

import (
	"fmt"
	"strings"
	pb "github.com/HedonisticAI/ytgogrps/proto"
)

func main() {
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		conn, err := grpc.Dial("http://localhost:8080/", opts...)
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		var s string
		CreateFolder("pic")
		checkErr(err)
		fmt.Printf("Enter Youtube Url or urls separated by ',':")
		fmt.Scanln(&s)
		var urls = strings.Split(s, ",")
		for i:=0; i < len(urls); i++ {
			//fmt.Println(urls[i])
			client := pb.NewReverseClient(conn)
    request := &pb.Request{
        Message: urls[i],
    }
    response, err := client.Do(context.Background(), request)

    if err != nil {
        grpclog.Fatalf("fail to dial: %v", err)
    }

   fmt.Println(response.Message)
			
		}


}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
