package message


//
const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string  `json:"type"`// 消息的类型
	Data string `json:"data"` // 消息的内容
}

// 先定义两个消息。后面需要再增加

type LoginMes struct {
	UserId int `json:"userId"` // 用户id
	UserPwd string `json:"userPwd"` // 用户密码
	UserName string `json:"userName"` // 用户名
}

type LoginResMes struct {
	Code int `json:"code"` // 返回状态码 500 表示用户未注册 200 表示用户登录成功
	Error string `json:"error"` // 返回错误信息
}

type RegisterMes struct {

}