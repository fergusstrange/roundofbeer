package persistence

import (
	"github.com/fergusstrange/roundofbeer/api/dynamo"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"time"
)

const roundOfBeer = "roundofbeer"

type Round struct {
	Url              string        `dynamo:"url,hash"`
	Participants     []Participant `dynamo:"participants"`
	CurrentCandidate string        `dynamo:"current_candidate"`
	CreateDate       time.Time     `dynamo:"create_date"`
	UpdateDate       time.Time     `dynamo:"update_date"`
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

func CreateRound(round *Round) {
	err := dynamo.Client.
		Table(roundOfBeer).
		Put(round).
		Run()
	errors.LogFatal(err)
}

func FetchRound(roundId string) *Round {
	round := new(Round)
	err := dynamo.Client.
		Table(roundOfBeer).
		Get("url", roundId).
		One(round)
	if err != nil {
		return nil
	}
	return round
}

func UpdateParticipantsAndCurrentCandidate(roundUrl string, participants []Participant, currentCandidate string) *Round {
	round := new(Round)
	err := dynamo.Client.Table(roundOfBeer).
		Update("url", roundUrl).
		Set("participants", participants).
		Set("current_candidate", currentCandidate).
		Set("update_date", time.Now()).
		Value(round)
	errors.LogFatal(err)
	return round
}
