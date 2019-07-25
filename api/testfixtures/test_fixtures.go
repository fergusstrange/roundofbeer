package testfixtures

import (
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/google/uuid"
	"time"
)

type MockPersistence struct {
	mockStore map[string]*persistence.Round
}

func NewMockPersistence() MockPersistence {
	return MockPersistence{
		mockStore: make(map[string]*persistence.Round),
	}
}

func (mp MockPersistence) CreateRoundTable() {}

func (mp MockPersistence) CreateRound(round *persistence.Round) {
	mp.mockStore[round.Url] = round
}

func (mp MockPersistence) FetchRound(roundId string) *persistence.Round {
	return mp.mockStore[roundId]
}

func (mp MockPersistence) UpdateParticipantsAndCurrentCandidate(updatedRound *persistence.Round) *persistence.Round {
	mp.mockStore[updatedRound.Url] = updatedRound
	return mp.mockStore[updatedRound.Url]
}

func (mp MockPersistence) NewTestRound(id string) *persistence.Round {
	tomsUUID := uuid.New().String()
	round := &persistence.Round{
		Participants: []persistence.Participant{
			{
				Name:       "Tom",
				UUID:       tomsUUID,
				RoundCount: 0,
			},
			{
				Name:       "John",
				UUID:       uuid.New().String(),
				RoundCount: 1,
			},
		},
		Url:              id,
		CreateDate:       time.Now(),
		UpdateDate:       time.Now(),
		CurrentCandidate: tomsUUID,
	}
	mp.CreateRound(round)
	return round
}
