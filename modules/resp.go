package modules

type BaseResp struct {
	LogID     int64  `json:"log_id"`
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

