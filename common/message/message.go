package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` // 消息類型
	Data string `json:"data"` // 消息的內容
}
type LoginMes struct {
	UserId   int    `json:"userId"`   // 用戶ID
	UserPwd  string `json:"userPwd"`  // 用戶密碼
	UserName string `json:"userName"` // 用戶名稱
}
type LoginResMes struct {
	Code  int    `json:"code"`  // 返回狀態碼 500 表示該用戶未註冊 200 表示登入成功
	Error string `json:"error"` // 返回錯誤訊息
}
