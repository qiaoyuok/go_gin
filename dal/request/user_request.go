package request

type UserReq struct {
	Name        string `json:"name"`
	Age         int32  `json:"age"`
	Description string `json:"description"`
	State       int32  `json:"state"`
}

// UserLoginReq 用户登录
type UserLoginReq struct {
	UserName string `json:"user_name" binding:"required,min=1"`
	Passwd   string `json:"passwd" binding:"required,min=1"`
}
