package persistence

import (
	"github.com/fergusstrange/roundofbeer/api/dynamo"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/google/uuid"
	"time"
)

const roundOfBeer = "roundofbeer"

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
		if table == roundOfBeer {
			return
		}
	}
	err = dynamo.Client.
		CreateTable(roundOfBeer, Round{}).
		Provision(3, 2).
		Run()
	errors.LogFatal(err)
}

func CreateRound(url string, participants []string) {
	var participantList []Participant
	for _, participant := range participants {
		participantList = append(participantList, Participant{
			UUID:       uuid.New().String(),
			Name:       participant,
			RoundCount: 0,
		})
	}
	err := dynamo.Client.
		Table(roundOfBeer).
		Put(Round{
			Url:          url,
			CreateDate:   time.Now(),
			UpdateDate:   time.Now(),
			Participants: participantList,
		}).
		Run()
	errors.LogFatal(err)
}

func FetchRound(roundId string) Round {
	var round Round
	err := dynamo.Client.
		Table(roundOfBeer).
		Get("url", roundId).
		One(&round)
	errors.LogFatal(err)
	return round
}

func UpdateRound(round Round) {
	err := dynamo.Client.Table(roundOfBeer).
		Update(round.Url, round).
		Run()
	errors.LogFatal(err)
}
