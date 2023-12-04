package accesslog

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type interceptor struct {
	log *logrus.Logger
}

func NewInterceptor(log *logrus.Logger) *interceptor {
	return &interceptor{log: log}
}

func (i *interceptor) Interceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	start := time.Now()

	reply, err := handler(ctx, req)

	i.log.Infof("after incoming call=%v req=%#v reply=%#v time=%v err=%v",
		info.FullMethod,
		req,
		reply,
		time.Since(start),
		err)

	return reply, err
}
