package request

type UserReq struct {
	Name        string `json:"name"`
	Age         int32  `json:"age"`
	Description string `json:"description"`
	State       int32  `json:"state"`
}
