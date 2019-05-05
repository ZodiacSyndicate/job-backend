package controllers

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Base struct {
	ID        uuid.UUID  `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (b *User) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", uuid.New())
	if err != nil {
		return err
	}
	return nil
}

// User table
type User struct {
	Base
	Username string `gorm:"unique;NOT NULL" json:"username"`
	Password string `gorm:"NOT NULL"`
	Type     string `gorm:"NOT NULL" json:"type"`
	Avatar   string `json:"avatar"`
	Desc     string `json:"desc"`
	Title    string `json:"title"`
	Company  string `json:"company"`
	Money    string `json:"money"`
}

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type registerResponse struct {
	Username string    `json:"username"`
	Type     string    `json:"type"`
	ID       uuid.UUID `json:"id"`
}

func md5Pwd(pwd string) string {
	data := []byte(pwd + "aksjhdI*(E*YDSYD&@IUhiu9(E*E")
	return fmt.Sprintf("%x", md5.Sum(data))
}

// Register handle user register
func Register(ctx echo.Context) error {
	req := new(registerRequest)
	if err := ctx.Bind(req); err != nil {
		return err
	}
	user := User{Username: req.Username, Password: md5Pwd(req.Password), Type: req.Type}
	err := db.Model(&User{}).First(&user, &User{Username: req.Username}).Error
	if err == gorm.ErrRecordNotFound {
		db.Create(&user)
		return ctx.JSON(http.StatusOK, &response{
			Code: 0,
			Msg:  "",
			Data: registerResponse{
				Username: user.Username,
				Type:     user.Type,
				ID:       user.Base.ID,
			},
		})
	}
	res := &response{
		Code: 1,
		Msg:  "User name has been taken",
		Data: nil,
	}
	return ctx.JSON(http.StatusBadRequest, res)
}
