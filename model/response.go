package model

// JsonData is api response body structure. HttpRes
type JsonData struct {
	Code    int         `json:"code" example:"0" description:"return status code, 0 is successful."`
	Message string      `json:"message" description:"return message"`
	Data    interface{} `json:"data" description:"return data"`
}

// JsonListData only use for swagger docs
type JsonListData struct {
	Code    int      `json:"code" description:"return code, 0 is successful."`
	Message string   `json:"message" description:"return message"`
	Data    []string `json:"data" description:"return data"`
}

// JsonMapData only use for swagger docs
type JsonMapData struct {
	Code    int    `json:"code" description:"return code, 0 is successful."`
	Message string `json:"message" description:"return message"`

	Data map[string]string `json:"data" description:"return data"`
}

