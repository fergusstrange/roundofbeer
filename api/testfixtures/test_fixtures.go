package testfixtures

import (
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"time"
)

type MockPersistence struct {
}

func (mp MockPersistence) CreateRoundTable() {}

func (mp MockPersistence) CreateRound(round *persistence.Round) {}

func (mp MockPersistence) FetchRound(roundId string) *persistence.Round {
	return aFakeRound()
}

func (mp MockPersistence) UpdateParticipantsAndCurrentCandidate(updatedRound *persistence.Round) *persistence.Round {
	return aFakeRound()
}

func aFakeRound() *persistence.Round {
	return &persistence.Round{
		Url:              "abc",
		CurrentCandidate: aFakeParticipant().UUID,
		Participants:     []persistence.Participant{aFakeParticipant()},
		UpdateDate:       time.Now(),
		CreateDate:       time.Now(),
	}
}

func aFakeParticipant() persistence.Participant {
	return persistence.Participant{
		Name:       "John",
		UUID:       "54de9ad7-7ae0-47f1-bb47-28b14d17c941",
		RoundCount: 1,
	}
}
