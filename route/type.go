package route

type (
	baseMsg struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	sendEmailReq struct {
		Time    string `json:"time"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}
	sendEmailResp struct {
		BaseMsg baseMsg `json:"base_msg"`
		Data    struct {
			IsSuccess bool `json:"is_success"`
		} `json:"data"`
	}
)
