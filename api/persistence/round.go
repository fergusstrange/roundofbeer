package persistence

import (
	"github.com/fergusstrange/roundofbeer/api/dynamo"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/google/uuid"
	"time"
)

const tableName = "roundofbeer"

type Round struct {
	Url          string        `dynamo:"url,hash"`
	Participants []Participant `dynamo:"participants"`
	CreateDate   time.Time     `dynamo:"create_date"`
	UpdateDate   time.Time     `dynamo:"update_date"`
}

type Participant struct {
	UUID       string `dynamo:"uuid"`
	Name       string `dynamo:"name"`
	RoundCount int    `dynamo:"round_count"`
}

func CreateRoundTable() {
	tables, err := dynamo.Client.ListTables().All()
	errors.LogFatal(err)
	for _, table := range tables {
		if table == tableName {
			return
		}
	}
	err = dynamo.Client.
		CreateTable(tableName, Round{}).
		Provision(3, 2).
		Run()
	errors.LogFatal(err)
}

func CreateRound(url string, participants []string) {
	participantList := make([]Participant, len(participants))
	for _, participant := range participants {
		participantList = append(participantList, Participant{
			UUID:       uuid.New().String(),
			Name:       participant,
			RoundCount: 0,
		})
	}
	err := dynamo.Client.
		Table(tableName).
		Put(Round{
			Url:          url,
			CreateDate:   time.Now(),
			UpdateDate:   time.Now(),
			Participants: participantList,
		}).
		Run()
	errors.LogFatal(err)
}
