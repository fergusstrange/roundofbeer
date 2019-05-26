const actions = {
  updateError: error => async state => ({
    ...state,
    error,
  }),
  clearError: () => async state => ({
    ...state,
    error: undefined,
  }),
  updateRoundToken: data => async state => ({
    ...state,
    roundToken: data.token,
    round: undefined,
    participatingRounds: Array.isArray(state.participatingRounds)
      ? [...state.participatingRounds.filter(pr => pr.roundUrl !== data.roundUrl),
        {
          roundUrl: data.roundUrl,
          roundToken: data.token,
          participants: data.participants,
        }]
      : [{
        roundUrl: data.roundUrl,
        roundToken: data.token,
        participants: data.participants,
      }],
  }),
  updateRound: round => async state => ({
    ...state,
    round,
  }),
};

export default actions;
