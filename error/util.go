package error

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func ConvertToGRPCError(err error) error {
	if clientErr, ok := err.(*ClientError); ok {
		switch clientErr.StatusCode {
		case http.StatusBadRequest:
			return status.Error(codes.InvalidArgument, clientErr.Message)
		case http.StatusNotFound:
			return status.Error(codes.NotFound, clientErr.Message)
		case http.StatusForbidden:
			return status.Error(codes.PermissionDenied, clientErr.Message)
		case http.StatusUnauthorized:
			return status.Error(codes.Unauthenticated, clientErr.Message)
		case http.StatusInternalServerError:
			return status.Error(codes.Internal, clientErr.Message)
		}
	}

	return status.Error(codes.Internal, err.Error())
}
