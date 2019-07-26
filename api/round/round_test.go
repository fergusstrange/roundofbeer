package round

import (
	"github.com/fergusstrange/roundofbeer/api/jwt"
	"github.com/fergusstrange/roundofbeer/api/testfixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformRound(t *testing.T) {
	round := TransformRound(testfixtures.NewTestRound("abcdef"))

	assert.Equal(t, &Round{
		URL: "abcdef",
		Participants: []Participant{
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
		CurrentCandidate: Participant{
			Name:       "Tom",
			UUID:       "50d4993a-183b-48ff-b175-31644a45021b",
			RoundCount: 0,
		},
	}, round)
}

func TestEncodeRoundToken(t *testing.T) {
	roundToken := EncodeRoundToken("abcdef")

	token, err := jwt.NewHelper().Decode(roundToken)

	assert.Nil(t, err)
	assert.Equal(t, "abcdef", token.RoundURL)
}

func TestParticipantsFromRound(t *testing.T) {
	participants := ParticipantsFromRound(testfixtures.NewTestRound("abcdef"))

	assert.Equal(t, []string{"Tom", "John"}, participants)
}


