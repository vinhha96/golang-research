package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vinhha96/golang-research/models"
	"math/rand"
	"net/http"
	"strconv"
)

func ShowProfilePage(ctx *gin.Context) {

}

func LoginPage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title": "Login",
	}, "login.html")
}

func PerformLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if models.IsUserValid(username, password) {
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

func generateUserToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func ShowRegistrationPage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title": "Register User",
	}, "register.html")
}

func Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	_, err := models.RegisterNewUser(username, password)
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

func LogoutPage(ctx *gin.Context) {
	// Clear the cookie
	ctx.SetCookie("token", "", -1, "", "", false, true)
	ctx.Set("is_logged_in", false)

	// Redirect to the home page
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
