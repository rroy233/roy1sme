package roy1sme

import "errors"

/*
HTTP error codes
*/
const (
	ResErrNotAuth            = -1001
	ResErrParameterNotValid  = -1002
	ResErrOther              = -1003
	ResErrPlanLimitExceeded  = -1004
	ResErrApiKeyExpired      = -1005
	ResErrPermissionDenied   = -1006
	ResErrCustomTokenInvalid = -1007
)

var errorMap = map[int]error{
	ResErrNotAuth:            errors.New("ResErrNotAuth"),
	ResErrParameterNotValid:  errors.New("ResErrParameterNotValid"),
	ResErrOther:              errors.New("ResErrOther"),
	ResErrPlanLimitExceeded:  errors.New("ResErrPlanLimitExceeded"),
	ResErrApiKeyExpired:      errors.New("ResErrApiKeyExpired"),
	ResErrPermissionDenied:   errors.New("ResErrPermissionDenied"),
	ResErrCustomTokenInvalid: errors.New("ResErrCustomTokenInvalid"),
}

/*
AuthType Authentication credential types
*/
type AuthType int

const (
	AuthTypeWebLogin = AuthType(iota)
	AuthTypeApiToken
)

var AuthTypeMap = map[AuthType]string{
	0: "Web",
	1: "ApiKey",
}

/*
UrlLife URL expiration time
*/
type UrlLife int

const (
	ExpireOneDay = UrlLife(iota)
	ExpireOneWeek
	ExpireOneMonth
	ExpireNever
)

var urlLifeSet []int64 = []int64{
	24 * 3600,
	7 * 24 * 3600,
	30 * 24 * 3600,
	-1,
}

// ReqCreate Request body for creating a short URL
type ReqCreate struct {
	Url         string  `json:"url" binding:"required"`
	CustomToken string  `json:"custom_token"`
	ExpireID    UrlLife `json:"expire_id"`
}

// RespCreate Response body for creating a short URL
type RespCreate struct {
	Status int      `json:"status"`
	Msg    string   `json:"msg"`
	Data   ShortUrl `json:"data"`
}

// ShortUrl Generated short URL
type ShortUrl struct {
	ShortUrl string `json:"short_url"`
	ExpireAt string `json:"expire_at"`
}

// RespUserHistory Response body for user creation history
type RespUserHistory struct {
	Status int               `json:"status"`
	Msg    string            `json:"msg"`
	Data   []UserHistoryItem `json:"data"`
}

// UserHistoryItem User creation history item
type UserHistoryItem struct {
	ID             int    `json:"id"`
	Token          string `json:"token"`
	ShortUrl       string `json:"short_url"`
	Url            string `json:"url"`
	CreatedThrough string `json:"created_through"`
	ExpireAt       string `json:"expire_at"`
	Icon           string `json:"icon"`
}
