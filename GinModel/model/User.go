package model

import (
	"fmt"
	"github.com/acmestack/gorm-plus/gplus"
)

func (User) TableName() string {
	return "users"
}

type User struct {
	UserName string // 用户名
	Password string // 密码
}

func Login(u *User) error {
	query, _ := gplus.NewQuery[User]()
	query.Eq("user_name", u.UserName)
	databaseUser, _ := gplus.SelectOne(query)
	if databaseUser.UserName == "" {
		insertResult := gplus.Insert(u)
		databaseUser = u
		fmt.Print("creat new user successfully", insertResult)
	}
	return nil
}

func ListUsers(pageNo int, pageSize int) gplus.Page[User] {
	query, _ := gplus.NewQuery[User]()
	page := gplus.NewPage[User](pageNo, pageSize)
	result, _ := gplus.SelectPage(page, query)
	return *result
}

func GetUserDetailByName(userName string) User {
	query, _ := gplus.NewQuery[User]()
	query.Eq("user_name", userName)
	reseult, _ := gplus.SelectOne(query)
	return *reseult
}
