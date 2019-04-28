package handlers

import (
	"github.com/fergusstrange/roundofbeer/api/jwt"
	"github.com/gin-gonic/gin"
)

func CreateRound(ctx *gin.Context) {
	token := ctx.GetHeader("x-roundofbeer-auth")
	jwt.NewHelper().Decode(token)
	ctx.Status(200)
}

func GetRound(ctx *gin.Context) {
	ctx.Status(200)
}

func GetRoundsByParticipant(ctx *gin.Context) {
	ctx.Status(200)
}

func IncrementRoundParticipant(ctx *gin.Context) {
	ctx.Status(200)
}
