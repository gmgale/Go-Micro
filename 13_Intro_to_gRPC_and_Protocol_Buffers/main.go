package main

import (
	"net"
	"os"

	protos "github.com/gmgale/go_micro/13_Intro_to_gRPC_and_Protocol_Buffers/protos/currency"
	"github.com/gmgale/go_micro/13_Intro_to_gRPC_and_Protocol_Buffers/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {

	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}


// func RegisterCurrencyServer(s grpc.ServiceRegistrar, srv CurrencyServer) {
// 	s.RegisterService(&Currency_ServiceDesc, srv)
// }
