package controller

import (
	"context"

	"github.com/Penetration-Platform-Go/Auth-Service/lib"
	user "github.com/Penetration-Platform-Go/gRPC-Files/User-Service"
	"github.com/gin-gonic/gin"
)

// ValidationRequest contains request for request user information
type ValidationRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Challenge string `json:"geetest_challenge"`
	Validate  string `json:"geetest_validate"`
	Seccode   string `json:"geetest_seccode"`
	Status    int8   `json:"geetest_status"`
}

// ValidationResponse contains response for validating user
type ValidationResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

// LogInHandler handler login event
func LogInHandler(ctx *gin.Context) {

	if !lib.VerifyUsernameFormat(ctx.PostForm("username")) {
		ctx.Status(406)
		return
	}

	// valid user password or email from user service
	uclient := user.NewUserClient(UserGrpcClient)
	userInformation, err := uclient.GetInformationByUsername(context.Background(), &user.Username{
		Username: ctx.PostForm("username"),
	})

	if err != nil {
		ctx.Status(400)
	} else if userInformation.Password != lib.StringToMd5(ctx.PostForm("password")) {
		ctx.Status(400)
	} else {
		ctx.JSON(200, userInformation)
	}
}
