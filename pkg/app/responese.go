package app

import (
	"Meow-backend/pkg/errcode"
	"Meow-backend/pkg/httpstatus"
	"Meow-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// Response defines a standard API response structure
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Details []string    `json:"details,omitempty"`
}

// NewResponse creates a new Response instance
func NewResponse() *Response {
	return &Response{}
}

// SuccessResponse sends a success response
func SuccessResponse(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(http.StatusOK, Response{
		Code:    errcode.Success.Code(),
		Message: errcode.Success.Message(),
		Data:    data,
	})
}

func SuccessResponseWithDetailed(c *gin.Context, data interface{}, detailMsg []string) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(http.StatusOK, Response{
		Code:    errcode.Success.Code(),
		Message: errcode.Success.Message(),
		Data:    data,
		Details: detailMsg,
	})
}

// Error sends an error response

// handleGRPCError handles gRPC errors
func handleGRPCError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		// If it's not a gRPC error, treat it as an internal server error
		c.JSON(http.StatusInternalServerError, Response{
			Code:    int(codes.Internal),
			Message: "Internal Server Error",
			Data:    gin.H{},
		})
		return
	}

	response := Response{
		Code:    int(st.Code()),
		Message: st.Message(),
		Data:    gin.H{},
		Details: make([]string, 0, len(st.Details())),
	}

	for _, detail := range st.Details() {
		response.Details = append(response.Details, cast.ToString(detail))
	}

	c.JSON(httpstatus.HTTPStatusFromCode(st.Code()), response)
}

func ErrorResponse(c *gin.Context, err error) {
	if err == nil {
		SuccessResponse(c, nil)
		return
	}

	switch typed := err.(type) {
	case *errcode.CustomError:
		response := Response{
			Code:    typed.Code(),
			Message: typed.Message(),
			Data:    gin.H{},
			Details: typed.Details(),
		}
		c.JSON(errcode.ToHTTPStatusCode(typed.Code()), response)

	case *errcode.Error:
		reponse := Response{
			Code:    typed.Code,
			Message: typed.Message,
			Data:    gin.H{},
		}
		c.JSON(errcode.ToHTTPStatusCode(typed.Code), reponse)

	default:
		handleGRPCError(c, err)
	}
}

func ErrorResponseWithData(c *gin.Context, data interface{}) {
	customError := errcode.ErrCustomError
	response := Response{
		Code: customError.Code(),
		Data: data,
	}

	c.JSON(errcode.ToHTTPStatusCode(customError.Code()), response)

}

func ErrorResponseWithMessage(c *gin.Context, message string) {
	customError := errcode.ErrCustomError
	response := Response{
		Code:    customError.Code(),
		Message: message,
	}

	c.JSON(errcode.ToHTTPStatusCode(customError.Code()), response)

}

func ErrorResponseWithDetailed(c *gin.Context, data interface{}, details []string) {
	customError := errcode.ErrCustomError

	response := Response{
		Code:    customError.Code(),
		Data:    data,
		Details: details,
	}

	c.JSON(errcode.ToHTTPStatusCode(customError.Code()), response)
}

// RouteNotFound handles 404 errors
func RouteNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "The route was not found")
}

// HealthCheckResponse defines the response structure for health checks
type HealthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HostnameHealthCheck performs a health check
func HostnameHealthCheck(c *gin.Context) {
	SuccessResponse(c, HealthCheckResponse{
		Status:   "UP",
		Hostname: utils.GetHostname(),
	})
}
