package main

import (
	"fmt"
	"strings"
	pb "github.com/HedonisticAI/ytgogrps/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"context"
	"log"
	"os"
)

func main() {
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		conn, err := grpc.Dial("127.0.0.1:5300", opts...)
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
	if response.Image != nil{
		for r:= range response.Image{
			file,_ := CreateFile("pic/"+t_rand()+".jpg")
			os.WriteFile(file.Name(),response.Image[r],0644)
		}
	} else {
		log.Printf("nothing found")
	}


		}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
