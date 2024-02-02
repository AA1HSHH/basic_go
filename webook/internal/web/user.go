package web

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/service"
	"errors"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (u *UserHandler) SignUp(ctx *gin.Context) {

	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	isEmail, err := u.emailRexExp.MatchString(req.Email)
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
	isPasswd, err := u.passwordRexExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isPasswd {
		ctx.String(http.StatusOK, "密码太简单，需要数字、字母和特殊字符，并且总长度要大于8")
		return
	}

	err = u.svc.Signup(ctx, domain.User{
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

func (u *UserHandler) Edit(ctx *gin.Context) {

}
func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
