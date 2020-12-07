/*
 * Copyright (C) 2020 Baidu, Inc. All Rights Reserved.
 */
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpcTest/client/services"
	"log"
)

func main() {
	creds,err := credentials.NewClientTLSFromFile("keys/ssl.crt","zsy.com")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)
	prodRes,err :=prodClient.GetProdStock(context.Background(),
		&services.ProdRequest{ProdId:2})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
}
