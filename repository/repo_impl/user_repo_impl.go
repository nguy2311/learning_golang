package repo_impl

import (
	"context"
	"learning_golang/banana"
	"learning_golang/db"
	"learning_golang/model"
	"learning_golang/model/req"
	"learning_golang/repository"
	"learning_golang/security"
	"time"

	"github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserRepoImpl struct {
	db *db.Database
}

func NewUserRepo(db *db.Database) repository.UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}

func (u *UserRepoImpl) SaveUser(ctx context.Context, user model.User) (model.User, error) {
	// Kiểm tra email đã tồn tại
	existingUser := model.User{}
	err := u.db.Db.Collection("users").FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		// Nếu tìm thấy người dùng có email trùng lặp, trả về lỗi xung đột
		return user, banana.UserConflict
	} else if err != mongo.ErrNoDocuments {
		// Nếu xảy ra lỗi khác ngoài việc không tìm thấy tài liệu, trả về lỗi
		return user, banana.SignUpFail
	}

	// Nếu email chưa tồn tại, tiếp tục thêm người dùng mới
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err = u.db.Db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		return user, banana.SignUpFail
	}
	return user, nil
}

func (u *UserRepoImpl) CheckLogin(ctx context.Context, loginReq req.ReqSignin) (model.User, error) {
	user := model.User{}
	err := u.db.Db.Collection("users").FindOne(ctx, bson.M{"email": loginReq.Email}).Decode(&user)
	if err != nil {
		return user, banana.UserNotFound
	}
	if !security.CheckPasswordHash(user.Password, []byte(loginReq.Password)) {
		return user, banana.PasswordErr
	}
	return user, nil
}
