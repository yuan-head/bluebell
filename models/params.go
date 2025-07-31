package models

// 定义请求参数的结构体
type ParamSignUp struct {
	Username   string `json:"username"`    // 用户名，必填，长度3-20
	Password   string `json:"password"`    // 密码，必填，长度6-20
	RePassword string `json:"re_password"` // 确认密码，必填，必须与Password相同
}
