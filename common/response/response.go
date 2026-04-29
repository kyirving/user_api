package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Output(w http.ResponseWriter, Code int, Msg interface{}, Data interface{}) {
	resp := &Resp{
		Code: Code,
		Data: Data,
	}
	defaultMsg, ok := HttpCodeMap[Code]
	if !ok {
		defaultMsg = "unknown"
	}

	if Msg != nil {
		resp.Msg = Msg.(string)
	} else {
		resp.Msg = defaultMsg
	}

	httpx.OkJson(w, resp)
}
