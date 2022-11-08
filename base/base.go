package base

type Resp struct {
	Code int `json:"code" example:"200"`
	Msg string `json:"msg" example:"成功"`
	Success bool `json:"success" example:"false"`
	Data interface{} `json:"data" example:"nil"`
}