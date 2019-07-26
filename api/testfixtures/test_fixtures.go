package testfixtures

import (
	"github.com/fergusstrange/roundofbeer/api/persistence"
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
	mp.mockStore[round.URL] = round
}

func (mp MockPersistence) FetchRound(roundID string) *persistence.Round {
	return mp.mockStore[roundID]
}

func (mp MockPersistence) UpdateParticipantsAndCurrentCandidate(updatedRound *persistence.Round) *persistence.Round {
	mp.mockStore[updatedRound.URL] = updatedRound
	return mp.mockStore[updatedRound.URL]
}

func (mp MockPersistence) NewTestRound(id string) *persistence.Round {
	round := NewTestRound(id)
	mp.CreateRound(round)
	return round
}

func NewTestRound(id string) *persistence.Round {
	return &persistence.Round{
		Participants: []persistence.Participant{
			{
				Name:       "Tom",
				UUID:       "50d4993a-183b-48ff-b175-31644a45021b",
				RoundCount: 0,
			},
			{
				Name:       "John",
				UUID:       "c547d1c9-ca1a-4ff6-83fc-e7df627ef45f",
				RoundCount: 1,
			},
		},
		URL:              id,
		CreateDate:       time.Now(),
		UpdateDate:       time.Now(),
		CurrentCandidate: "50d4993a-183b-48ff-b175-31644a45021b",
	}
}
