package handlers

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Round struct {
	Url          string        `json:"url"`
	Participants []Participant `json:"participants"`
	CreateDate   time.Time     `json:"create_date"`
	UpdateDate   time.Time     `json:"update_date"`
}

type Participant struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	RoundCount int    `json:"round_count"`
}

func IncrementRoundParticipant(ctx *gin.Context) {
	ctx.Status(200)
}
