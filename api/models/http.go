package models

type Status struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type Response struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

var (
	OK = Status{
		Code:        200,
		Status:      "OK",
		Description: "The request has succeeded",
	}
	Created = Status{
		Code:        201,
		Status:      "CREATED",
		Description: "The request has been fulfilled and has resulted in one or more new resources being created",
	}
	BadRequest = Status{
		Code:        400,
		Status:      "BAD_REQUEST",
		Description: "The server could not understand the request due to invalid syntax",
	}
	InvalidArgument = Status{
		Code:        400,
		Status:      "INVALID_ARGUMENT",
		Description: "Invalid argument value passed",
	}
	InternalServerError = Status{
		Code:        500,
		Status:      "INTERNAL_SERVER_ERROR",
		Description: "The server encountered an unexpected condition that prevented it from fulfilling the request",
	}
	GRPCError = Status{
		Code:        500,
		Status:      "GRPC_ERROR",
		Description: "The gRPC request failed",
	}
)
