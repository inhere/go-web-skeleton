package errcode

import "errors"

// Predefined const error codes.
// message please see /resource/lang/*.ini
const (
	OK  = 0
	ERR = 500

	// error codes
	ErrSvrNotAvailable = 2100
	ErrServer          = 2101
	ErrNoPermission    = 2102
	ErrNotFound        = 2104
	ErrNoResource      = 2105
	ErrParams          = 2106
	ErrMissingParams   = 2107
	ErrNoRecord        = 2108
	ErrOpFail          = 2112
	ErrInvalidReq      = 2113
	ErrParam           = 2114
	InvalidParam       = 2114

	ErrReqData     = 2201
	ErrRepeatOp    = 2202
	ErrDatabase    = 2210
	ErrInvalidData = 2211
	ErrCheckFail   = 2212

	ErrDupRows    = 2406
	ErrInsertFail = 2401
	ErrUpdateFail = 2403
	ErrDeleteFail = 2404
)

var (
	InvalidParamErr = errors.New("invalid param")
)
