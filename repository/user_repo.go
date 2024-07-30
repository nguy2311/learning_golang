package repository

import (
	"context"
	"learning_golang/model"
	"learning_golang/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginReq req.ReqSignin) (model.User, error)
}
