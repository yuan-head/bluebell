package models

// 定义请求参数的结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`                     // 用户名，必填，长度3-20
	Password   string `json:"password" binding:"required"`                     // 密码，必填，长度6-20
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 确认密码，必填，必须与Password相同
}
