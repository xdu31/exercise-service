package util

import (
	"context"
	"fmt"

	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SetOK(ctx context.Context, format string, args ...interface{}) error {
	return gateway.SetStatus(ctx, status.New(codes.OK, fmt.Sprintf(format, args...)))
}
