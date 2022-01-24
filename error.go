package connect

import "fmt"

type Err struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func (e Err) Error() string {
	return fmt.Sprintf("Connect API Error: error code [%d] error message [%s]", e.ErrorCode, e.Message)
}

func (e Err) IsRebalacing() bool {
	if e.ErrorCode == 409 {
		return true
	}

	return false
}
