/*
 * Copyright (C) 2020 Baidu, Inc. All Rights Reserved.
 */
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpcTest/server/services"
	"log"
	"net/http"
)

func main()  {
	creds,err := credentials.NewServerTLSFromFile("keys/ssl.crt","keys/ssl.key")
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))
	//lis,_ := net.Listen("tcp",":8081")
	//rpcServer.Serve(lis)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("%+v",*request)
		rpcServer.ServeHTTP(writer,request)
	})
	httpServer := http.Server{
		Addr:":8081",
		Handler:mux,
	}
	httpServer.ListenAndServeTLS("keys/ssl.crt","keys/ssl.key")
}