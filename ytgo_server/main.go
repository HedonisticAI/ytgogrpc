package main

import (
	"log"
    "net"
    pb "github.com/HedonisticAI/ytgogrps/proto"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
	//"os"
	"strings"
)
var in_mem map[string][]byte



func main() {

	/*var mod bool = false
	if(os.Args[1]) == "--async"{
		mod = true
	}*/
	in_mem = make(map[string][]byte)
    listener, err := net.Listen("tcp", ":5300")

    if err != nil {
        grpclog.Fatalf("failed to listen: %v", err)
    }

    opts := []grpc.ServerOption{}
    grpcServer := grpc.NewServer(opts...)

    pb.RegisterReverseServer(grpcServer, &server{})
    grpcServer.Serve(listener)
}

type server struct{}

func (s *server) Do(c context.Context, request *pb.Request)(response *pb.Response, err error) {
	log.Printf(ParseID(request.Message))
	var ans [][]byte
	var str []string = strings.Split(request.Message,",")
	for r  :=range str {
		_, ok :=in_mem[str[r]] 
		if  ok {
			log.Printf("in cash")
			ans = append(ans, in_mem[ParseID(str[r])])
		}
		if !ok {
			var data,err1 = downloadFile(ParseID(str[r]))
			if(err1 != nil){
			response = &pb.Response{
					Image: nil,
				}
				return response , nil
			}
			ans = append(ans,data)
			in_mem[ParseID(str[r])] = data
    		response = &pb.Response{
        		Image: ans,
    		}
    			return response, nil   
			}
			response = &pb.Response{
					Image: nil,
				}
				return response , nil
			}
			return nil, nil
	}
		