package handler

const (
	Success = iota
	ParamsError
	AuthError
	LogicError
	InnerError
)

var msgMAP = map[int]string{
	Success:     "",
	ParamsError: "参数错误！",
	AuthError:   "认证错误！",
	LogicError:  "逻辑错误！",
	InnerError:  "服务器错误！",
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CodeMsg struct {
	Code     int
	ExtraMsg string
}

func (resp *Response) SetMsg(msg CodeMsg) {
	Msg, ok := msgMAP[msg.Code]
	if ok {
		resp.Msg = Msg + "  " + msg.ExtraMsg
	}
	resp.Code = msg.Code
}
