package persistence

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/fergusstrange/roundofbeer/api/errors"
	"github.com/guregu/dynamo"
	"time"
)

const roundOfBeer = "roundofbeer"

type Persistence interface {
	CreateRoundTable()
	CreateRound(round *Round)
	FetchRound(roundId string) *Round
	UpdateParticipantsAndCurrentCandidate(updatedRound *Round) *Round
}

type DynamoDBPersistence struct {
	client *dynamo.DB
}

func NewDynamoDBPersistence() *DynamoDBPersistence {
	return &DynamoDBPersistence{
		client: newClient(),
	}
}

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

func (db *DynamoDBPersistence) CreateRoundTable() {
	tables, err := db.client.ListTables().All()
	errors.LogFatal(err)
	for _, table := range tables {
		if table == roundOfBeer {
			return
		}
	}
	err = db.client.
		CreateTable(roundOfBeer, Round{}).
		Provision(3, 2).
		Run()
	errors.LogFatal(err)
}

func (db *DynamoDBPersistence) CreateRound(round *Round) {
	err := db.client.
		Table(roundOfBeer).
		Put(round).
		Run()
	errors.LogFatal(err)
}

func (db *DynamoDBPersistence) FetchRound(roundId string) *Round {
	round := new(Round)
	err := db.client.
		Table(roundOfBeer).
		Get("url", roundId).
		One(round)
	if err != nil {
		return nil
	}
	return round
}

func (db *DynamoDBPersistence) UpdateParticipantsAndCurrentCandidate(updatedRound *Round) *Round {
	persistedRound := new(Round)
	err := db.client.Table(roundOfBeer).
		Update("url", updatedRound.Url).
		Set("participants", updatedRound.Participants).
		Set("current_candidate", updatedRound.CurrentCandidate).
		Set("update_date", time.Now()).
		Value(persistedRound)
	errors.LogFatal(err)
	return persistedRound
}

func newClient() *dynamo.DB {
	newSession, err := session.NewSession()
	errors.LogFatal(err)
	return dynamo.New(newSession)
}
