package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vinhha96/golang-research/models"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (userHandler *UserHandler) ShowProfilePage(ctx *gin.Context) {

}

func (userHandler *UserHandler) LoginPage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title": "Login",
	}, "login.html")
}

func (userHandler *UserHandler) PerformLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if userHandler.isUserValid(username, password) {
		token := generateUserToken()
		ctx.SetCookie("token", token, 3600, "", "", false, true)
		ctx.Set("is_logged_in", true)

		render(ctx, gin.H{
			"title": "Login successfully",
		}, "login-successful.html")
	} else {
		ctx.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided",
		})
	}
}

func (userHandler *UserHandler) ShowRegistrationPage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title": "Register User",
	}, "register.html")
}

func (userHandler *UserHandler) Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	_, err := userHandler.registerNewUser(username, password)
	if err == nil {
		token := generateUserToken()
		ctx.SetCookie("token", token, 3600, "", "", false, true)
		ctx.Set("is_logged_in", true)

		render(ctx, gin.H{
			"title": "Successful registration & Login",
		}, "login-successful.html")
	} else {
		ctx.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}
}

func (userHandler *UserHandler) LogoutPage(ctx *gin.Context) {
	// Clear the cookie
	ctx.SetCookie("token", "", -1, "", "", false, true)
	ctx.Set("is_logged_in", false)

	// Redirect to the home page
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}

func generateUserToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func (userHandler *UserHandler) isUserValid(username, password string) bool {
	var userList []models.User

	userHandler.DB.Find(&userList)

	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func (userHandler *UserHandler) registerNewUser(username, password string) (*models.User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("Password can't not empty")
	} else if !userHandler.isUserNameAvailable(username) {
		return nil, errors.New("Username is not available")
	}

	newUser := models.User{Username: username, Password: password}
	userHandler.DB.Save(&newUser)

	return &newUser, nil
}

func (userHandler *UserHandler) isUserNameAvailable(username string) bool {
	var userList []models.User

	userHandler.DB.Find(&userList)

	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}
