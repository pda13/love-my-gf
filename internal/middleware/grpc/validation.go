package grpcmw

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

type validator interface {
	Validate() error
}

func ValidationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if validator, ok := req.(validator); ok {
			if err := validator.Validate(); err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
			}
		} else {
			if validateAll := reflect.ValueOf(req).MethodByName("ValidateAll"); validateAll.IsValid() {
				if err := validateAll.Call(nil)[0].Interface(); err != nil {
					return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
				}
			}
		}

		return handler(ctx, req)
	}
}
