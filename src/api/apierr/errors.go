package apierr

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// New returns new API error with given API code and message
func New(code int, msg string) *Error {
	return &Error{code, msg}
}

// NewUndefined returns new API error with undefined (0) code and given message
func NewUndefined(msg string) *Error {
	return &Error{CodeUndefined, msg}
}

const (
	CodeUndefined   = 0
	CodeParseFailed = 1001
)

var (
	ErrParseFailed = &Error{CodeParseFailed, "Failed to parse request"}
)

// Response responses with status code and writes given API error to body
func Response(c *gin.Context, code int, err *Error) {
	c.AbortWithStatusJSON(code, gin.H{
		"error": err,
	})
}

// ResponseMsg responses with status code and writes API error with
// 0 code (undefined) and given message to body
func ResponseMsg(c *gin.Context, code int, msg string) {
	Response(c, code, NewUndefined(msg))
}

// ResponseErr responses with status code and writes API error with
// 0 code (undefined) and given error as message to body
func ResponseErr(c *gin.Context, code int, err error) {
	ResponseMsg(c, code, err.Error())
}
