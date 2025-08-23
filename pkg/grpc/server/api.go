package server

import "github.com/anonpragmatic/gowaves/pkg/grpc/generated/waves/node/grpc"

type GrpcHandlers interface {
	grpc.AccountsApiServer
	grpc.AssetsApiServer
	grpc.BlockchainApiServer
	grpc.BlocksApiServer
	grpc.TransactionsApiServer
}
