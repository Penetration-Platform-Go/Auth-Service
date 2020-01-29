package controller

import (
	"context"

	"github.com/Penetration-Platform-Go/Auth-Service/lib"
	user "github.com/Penetration-Platform-Go/gRPC-Files/User-Service"
	"github.com/gin-gonic/gin"
)

// LogInHandler handler login event
func LogInHandler(ctx *gin.Context) {
	username := ctx.PostForm("username")
	if !lib.VerifyUsernameFormat(username) {
		ctx.Status(406)
		return
	}

	// valid user password or email from user service
	uclient := user.NewUserClient(UserGrpcClient)
	userInformation, err := uclient.GetInformationByUsername(context.Background(), &user.Username{
		Username: username,
	})

	if err != nil {
		ctx.Status(400)
	} else if userInformation.Password != lib.StringToMd5(ctx.PostForm("password")) {
		ctx.Status(400)
	} else {
		token, err := lib.GenerateJWT(username)
		if err != nil {
			ctx.Status(500)
		} else {
			ctx.Header("Authenticate", token)
			ctx.JSON(200, userInformation)
		}

	}
}
