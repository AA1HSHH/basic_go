package web

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/service"
	"errors"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"unicode/utf8"
)

const (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:            svc,
	}
}
func (u *UserHandler) RegisterRouters(server *gin.Engine) {
	server.POST("/users/signup", u.SignUp)
	server.POST("/users/login", u.LogIn)
	server.POST("/users/edit", u.Edit)
	server.GET("/users/profile", u.Profile)
}

func (h *UserHandler) SignUp(ctx *gin.Context) {

	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	isEmail, err := h.emailRexExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "邮箱格式不对")
		return
	}
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次密码不相同")
		return
	}
	isPasswd, err := h.passwordRexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isPasswd {
		ctx.String(http.StatusOK, "密码太简单，需要数字、字母和特殊字符，并且总长度要大于8")
		return
	}

	err = h.svc.Signup(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	switch {
	case errors.Is(err, service.ErrUserDuplicateEmail):
		ctx.String(http.StatusOK, "邮箱重复")
	case err == nil:
		ctx.String(http.StatusOK, "注册成功")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}

}

func (h *UserHandler) LogIn(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := h.svc.Login(ctx, req.Email, req.Password)

	switch err {
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或密码不对")
	case nil:

		sess := sessions.Default(ctx)
		sess.Set("userId", user.Id)
		sess.Options(sessions.Options{MaxAge: 900})
		err = sess.Save()
		if err != nil {
			ctx.String(http.StatusOK, "服务器异常")
			return
		}

		ctx.String(http.StatusOK, "登录成功")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}

}

func (h *UserHandler) Edit(ctx *gin.Context) {
	type Req struct {
		AboutMe  string `json:"aboutMe"`
		Birthday string `json:"birthday"`
		Nickname string `json:"nickname"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// 字段必须大写
	type Res struct {
		Code int64  `json:"code"`
		Msg  string `json:"msg"`
	}

	if utf8.RuneCountInString(req.Nickname) >= 15 {
		ctx.JSON(http.StatusForbidden, Res{
			Code: 1,
			Msg:  "昵称过长，需要少于16个中英文字符",
		})
		return
	}
	if utf8.RuneCountInString(req.AboutMe) >= 25 {
		ctx.JSON(http.StatusForbidden, Res{
			Code: 1,
			Msg:  "简介过长，需要少于26个中英文字符",
		})
		return
	}
	sess := sessions.Default(ctx)
	userId := sess.Get("userId")

	if err := h.svc.Edit(ctx, userId.(int64), req.AboutMe, req.Birthday, req.Nickname); err != nil {
		ctx.JSON(http.StatusOK, Res{
			Code: 1,
			Msg:  "系统错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, Res{
		Code: 0,
		Msg:  "更新成功",
	})
}
func (h *UserHandler) Profile(ctx *gin.Context) {

	sess := sessions.Default(ctx)
	userId := sess.Get("userId")
	u, err := h.svc.Profile(ctx, userId.(int64))
	// 字段必须大写

	type Res struct {
		Id       int64
		Email    string
		Nickname string
		Birthday string
		AboutMe  string
	}
	switch err {
	case service.ErrUserNotFound:
		ctx.String(http.StatusOK, "用户不存在")
	case nil:
		ctx.JSON(http.StatusOK, Res{Id: u.Id, Email: u.Email, Nickname: u.Nickname, Birthday: u.Birthday, AboutMe: u.AboutMe})
	default:
		ctx.String(http.StatusOK, "系统错误")

	}

}
