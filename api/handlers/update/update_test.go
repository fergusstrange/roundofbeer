package update

import (
	"github.com/fergusstrange/roundofbeer/api/persistence"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_selectNextCandidate_OneParticipant_ExpectJesus(t *testing.T) {
	candidate := selectNextCandidate(persistence.Round{
		Participants: []persistence.Participant{
			{
				UUID:       "6b9f56c7-2035-43d3-b65a-c7cfc9c32b86",
				Name:       "Jesus",
				RoundCount: 1,
			},
		},
	})

	assert.Equal(t,
		persistence.Participant{
			UUID:       "6b9f56c7-2035-43d3-b65a-c7cfc9c32b86",
			Name:       "Jesus",
			RoundCount: 1,
		},
		candidate)
}

func Test_selectNextCandidate_TwoParticipants_ExpectJohn(t *testing.T) {
	candidate := selectNextCandidate(persistence.Round{
		Participants: []persistence.Participant{
			{
				UUID:       "6b9f56c7-2035-43d3-b65a-c7cfc9c32b86",
				Name:       "John",
				RoundCount: 0,
			},
			{
				UUID:       "6184606d-e9b8-4a76-98fa-7ed56bb85239",
				Name:       "James",
				RoundCount: 1,
			},
		},
	})

	assert.Equal(t,
		persistence.Participant{
			UUID:       "6b9f56c7-2035-43d3-b65a-c7cfc9c32b86",
			Name:       "John",
			RoundCount: 0,
		},
		candidate)
}

func Test_selectNextCandidate_TwoParticipants_ExpectAny(t *testing.T) {
	candidate := selectNextCandidate(persistence.Round{
		Participants: []persistence.Participant{
			{
				UUID:       "6b9f56c7-2035-43d3-b65a-c7cfc9c32b86",
				Name:       "John",
				RoundCount: 1,
			},
			{
				UUID:       "6184606d-e9b8-4a76-98fa-7ed56bb85239",
				Name:       "James",
				RoundCount: 1,
			},
		},
	})

	assert.Condition(t, func() (success bool) {
		return assert.ObjectsAreEqual(persistence.Participant{
			UUID:       "6b9f56c7-2035-43d3-b65a-c7cfc9c32b86",
			Name:       "John",
			RoundCount: 1,
		}, candidate) ||
			assert.ObjectsAreEqual(persistence.Participant{
				UUID:       "6184606d-e9b8-4a76-98fa-7ed56bb85239",
				Name:       "James",
				RoundCount: 1,
			}, candidate)
	})
}

func Test_selectNextCandidate_ThreeParticipants_ExpectJim(t *testing.T) {
	candidate := selectNextCandidate(persistence.Round{
		Participants: []persistence.Participant{
			{
				UUID:       "6b9f56c7-2035-43d3-b65a-c7cfc9c32b86",
				Name:       "John",
				RoundCount: 3,
			},
			{
				UUID:       "676163cb-2357-4f7e-b0cd-af942a83ada3",
				Name:       "Jim",
				RoundCount: 1,
			},
			{
				UUID:       "6184606d-e9b8-4a76-98fa-7ed56bb85239",
				Name:       "James",
				RoundCount: 2,
			},
		},
	})

	assert.Equal(t,
		persistence.Participant{
			UUID:       "676163cb-2357-4f7e-b0cd-af942a83ada3",
			Name:       "Jim",
			RoundCount: 1,
		},
		candidate)
}

func Test_selectNextCandidate_ThreeParticipants_ExpectJose(t *testing.T) {
	candidate := selectNextCandidate(persistence.Round{
		Participants: []persistence.Participant{
			{
				UUID:       "6b9f56c7-2035-43d3-b65a-c7cfc9c32b86",
				Name:       "John",
				RoundCount: 3,
			},
			{
				UUID:       "676163cb-2357-4f7e-b0cd-af942a83ada3",
				Name:       "Jim",
				RoundCount: 3,
			},
			{
				UUID:       "6184606d-e9b8-4a76-98fa-7ed56bb85239",
				Name:       "Jose",
				RoundCount: 0,
			},
		},
	})

	assert.Equal(t,
		persistence.Participant{
			UUID:       "6184606d-e9b8-4a76-98fa-7ed56bb85239",
			Name:       "Jose",
			RoundCount: 0,
		},
		candidate)
}

func Test_randomIndexOrFirstWhenOnlyOneCandidate_Expect0(t *testing.T) {
	index := randomIndexOrFirstWhenOnlyOneCandidate([]persistence.Participant{
		{
			UUID:       "6b9f56c7-2035-43d3-b65a-c7cfc9c32b86",
			Name:       "John",
			RoundCount: 1,
		},
	})

	assert.Equal(t, 0, index)
}

func Test_randomIndexOrFirstWhenOnlyOneCandidate_Expect0Or1(t *testing.T) {
	index := randomIndexOrFirstWhenOnlyOneCandidate([]persistence.Participant{
		{
			UUID:       "6b9f56c7-2035-43d3-b65a-c7cfc9c32b86",
			Name:       "John",
			RoundCount: 1,
		},
		{
			UUID:       "6184606d-e9b8-4a76-98fa-7ed56bb85239",
			Name:       "James",
			RoundCount: 1,
		},
	})

	assert.Condition(t, func() (success bool) {
		return index == 0 || index == 1
	})
}
