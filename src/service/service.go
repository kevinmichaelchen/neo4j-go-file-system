package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"net/http"
)

const UserIDContextKey = "userID"

func GetUserID(ctx context.Context) (int, *Error) {
	userIDPayload := ctx.Value(UserIDContextKey)
	if userIDPayload == nil {
		return 0, &Error{
			HttpCode:     http.StatusUnauthorized,
			GrpcCode:     codes.Unauthenticated,
			ErrorMessage: "User ID not found",
			Error:        nil,
		}
	}
	userID, ok := userIDPayload.(int)
	if !ok {
		return 0, &Error{
			HttpCode:     http.StatusBadRequest,
			GrpcCode:     codes.InvalidArgument,
			ErrorMessage: "User ID is invalid. Should be an int.",
			Error:        nil,
		}
	}
	return userID, nil
}

func CreateUserContext(userID int) context.Context {
	return context.WithValue(context.Background(), UserIDContextKey, userID)
}

type Error struct {
	HttpCode     int
	GrpcCode     codes.Code
	ErrorMessage string
	Error        error
}

func NewError(httpCode int, errorMessage string, err error) *Error {
	return &Error{
		HttpCode:     httpCode,
		ErrorMessage: errorMessage,
		Error:        err,
	}
}

func Internal(err error) *Error {
	return &Error{
		HttpCode:     http.StatusInternalServerError,
		GrpcCode:     codes.Internal,
		ErrorMessage: err.Error(),
		Error:        err,
	}
}

func Unimplemented() *Error {
	return &Error{
		HttpCode:     http.StatusNotImplemented,
		GrpcCode:     codes.Unimplemented,
		ErrorMessage: "Unimplemented",
		Error:        nil,
	}
}
