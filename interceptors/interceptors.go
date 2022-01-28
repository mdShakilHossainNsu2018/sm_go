package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("StreamInterceptor Called: ", info.FullMethod)
	return handler(srv, ss)
}

func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("UnaryInterceptor called: ", info.FullMethod)

	return handler(ctx, req)

}
