/*
 * Copyright (C) 2020 Baidu, Inc. All Rights Reserved.
 */
package services

import "context"

type ProdService struct {

}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock:25},nil

}