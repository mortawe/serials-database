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
	CodeUndefined     = 0
	CodeParseFailed   = 1001
	CodeRouteNotFound = 1004

	CodeUserNotFound         = 1100
	CodePassMismatch         = 1101
	CodeEmailTaken           = 1102
	CodeIncorrectCredentials = 1103
	CodeEmailInvalid         = 1104
	CodeIncorrectCode        = 1105
	CodeInvalidOTP           = 1106
	CodeInvalidPassword      = 1107
	CodeUserBlocked          = 1108

	CodeFileInTrash     = 1200
	CodeFileNotFound    = 1201
	CodeWrongFileType   = 1203
	CodeFilenameInvalid = 1204
	CodeDuplicateName   = 1205
	CodeRecursiveMoving = 1206

	CodeInvalidSMTPServer = 1300
	CodeInvalidSMTPPort   = 1301
	CodeInvalidSMTPUser   = 1302
	CodeInvalidSMTPPass   = 1303
	CodeEmptyDomains      = 1304

	CodeSelfAccess    = 1400
	CodeAccessDenied  = 1401
	CodeAccessToOwner = 1403
	CodeAccessExists  = 1404
)

var (
	ErrParseFailed   = &Error{CodeParseFailed, "Failed to parse request"}
	ErrRouteNotFound = &Error{CodeRouteNotFound, "No such route"}

	ErrUserNotFound         = &Error{CodeUserNotFound, "User not found"}
	ErrPassMismatch         = &Error{CodePassMismatch, "Passwords are different"}
	ErrEmailTaken           = &Error{CodeEmailTaken, "Email is already taken"}
	ErrIncorrectCredentials = &Error{CodeIncorrectCredentials, "Incorrect credentials"}
	ErrEmailInvalid         = &Error{CodeEmailInvalid, "Invalid email"}
	ErrIncorrectCode        = &Error{CodeIncorrectCode, "Incorrect code"}
	ErrInvalidOTP           = &Error{CodeInvalidOTP, "Invalid confirmation code"}
	ErrInvalidPassword      = &Error{CodeInvalidPassword, "Invalid password"}
	ErrUserBlocked          = &Error{CodeInvalidPassword, "User has been blocked"}

	ErrFileInTrash     = &Error{CodeFileInTrash, "File is already in trash"}
	ErrFileNotFound    = &Error{CodeFileNotFound, "File not found"}
	ErrWrongFileType   = &Error{CodeWrongFileType, "Wrong file type (file or directory)"}
	ErrFilenameInvalid = &Error{CodeFilenameInvalid, "File name is invalid"}
	ErrDuplicateName   = &Error{CodeDuplicateName, "Name is already taken"}
	ErrRecursiveMoving = &Error{CodeRecursiveMoving, "Recursive directory moving"}

	ErrInvalidSMTPServer = &Error{CodeInvalidSMTPServer, "SMTP server is invalid"}
	ErrInvalidSMTPPort   = &Error{CodeInvalidSMTPPort, "SMTP port is invalid"}
	ErrInvalidSMTPUser   = &Error{CodeInvalidSMTPUser, "SMTP username is invalid"}
	ErrInvalidSMTPPass   = &Error{CodeInvalidSMTPPass, "SMTP password is invalid"}
	ErrEmptyDomains      = &Error{CodeEmptyDomains, "Domains list is empty"}

	ErrSelfAccess    = &Error{CodeSelfAccess, "Attempt to grant access to oneself"}
	ErrAccessDenied  = &Error{CodeAccessDenied, "Access denied"}
	ErrAccessToOwner = &Error{CodeAccessToOwner, "Attempt to grant access to file owner"}
	ErrAccessExists  = &Error{CodeAccessExists, "Access has been already granted"}
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
