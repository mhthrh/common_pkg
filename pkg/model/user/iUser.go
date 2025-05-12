package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"time"
)

type User struct {
	id          uuid.UUID
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	status      Status
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	createdAt   time.Time
	updatedAt   time.Time
	lastLogin   time.Time
}

type IUser interface {
	Create(ctx context.Context, user *User) *xErrors.Error
	GetByUserName(ctx context.Context, userName string) (User, *xErrors.Error)
	Update(ctx context.Context, user *User) *xErrors.Error
	Remove(ctx context.Context, userName string) *xErrors.Error
}

func UserInit(u *User) *User {
	u.id = uuid.New()
	u.createdAt = time.Now()
	return u
}
