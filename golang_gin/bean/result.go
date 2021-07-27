package bean

type Result struct {
	Msg  string      `json:"msg"`
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}
