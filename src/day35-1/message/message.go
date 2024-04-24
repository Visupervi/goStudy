package message

const (
	LoginMsgType = "LoginMsg"
	LoginResType = "LoginResult"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserId   int    `json:"userId"`
	Pwd      string `json:"pwd"`
	UserName string `json:"userName"`
}

type LoginResult struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
