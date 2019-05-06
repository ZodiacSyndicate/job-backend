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

const cookieKey = "USER_ID"

type Base struct {
	ID        uuid.UUID  `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("ID", uuid.New()); err != nil {
		return err
	}
	return nil
}

// User table
type User struct {
	Base
	Username    string `gorm:"unique;NOT NULL" json:"username"`
	Password    string `gorm:"NOT NULL"`
	Type        string `gorm:"NOT NULL" json:"type"`
	Avatar      string `json:"avatar"`
	Description string `json:"desc"`
	Title       string `json:"title"`
	Company     string `json:"company"`
	Money       string `json:"money"`
}

type userResponse struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Type        string    `json:"type"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"desc"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Money       string    `json:"money"`
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
		e := db.Create(&user).Error
		if e == nil {
			ctx.SetCookie(&http.Cookie{Name: cookieKey, Value: user.Base.ID.String()})
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
		return e
	}
	if err != nil {
		return err
	}
	res := &response{
		Code: 1,
		Msg:  "User name has been taken",
	}
	return ctx.JSON(http.StatusBadRequest, res)
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login handle user login
func Login(ctx echo.Context) error {
	req := new(loginRequest)
	var (
		err  error
		user User
		res  userResponse
	)
	if err = ctx.Bind(req); err != nil {
		return err
	}
	err = db.Model(&User{}).First(&user, &User{Username: req.Username}).Scan(&res).Error
	if err == gorm.ErrRecordNotFound {
		return ctx.JSON(http.StatusBadRequest, &response{
			Code: 1,
			Msg:  "Username dose not exist",
		})
	}
	if err != nil {
		return err
	}
	if md5Pwd(req.Password) != user.Password {
		return ctx.JSON(http.StatusBadRequest, &response{
			Code: 1,
			Msg:  "Wrong password",
		})
	}
	ctx.SetCookie(&http.Cookie{Name: cookieKey, Value: user.Base.ID.String()})
	return ctx.JSON(http.StatusOK, &response{
		Code: 0,
		Msg:  "",
		Data: res,
	})
}

// Info handle get info
func Info(ctx echo.Context) error {
	var res userResponse
	cookie, err := ctx.Cookie(cookieKey)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response{Code: 1})
	}
	err = db.Model(&User{}).Where("id = ?", cookie.Value).Scan(&res).Error
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response{Code: 1})
	}
	return ctx.JSON(http.StatusOK, &response{
		Code: 0,
		Data: res,
	})
}

// List get user list
func List(ctx echo.Context) error {
	var res []*userResponse
	err := db.Model(&User{}).Where("type = ?", ctx.QueryParams().Get("type")).Scan(&res).Error
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response{Code: 1})
	}
	return ctx.JSON(http.StatusOK, &response{
		Code: 0,
		Data: res,
	})
}
