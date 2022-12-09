package biz

import (
	"github.com/jinzhu/copier"
	"go_gin/dal/model"
	"go_gin/dal/query"
	"go_gin/dal/request"
)

func ListUser() (users []*model.User, err error) {
	users, err = query.User.Find()
	return
}

func Create(params request.UserReq) {
	user := new(model.User)
	copier.Copy(user, params)
	query.User.Create(user)

}
