const actions = {
  updateError: error => async state => ({
    ...state,
    error,
  }),
  clearError: () => async state => ({
    ...state,
    error: undefined,
  }),
  updateRoundToken: (roundUrl, roundToken) => async state => ({
    ...state,
    roundToken,
    round: undefined,
    participatingRounds: Array.isArray(state.participatingRounds)
      ? [...state.participatingRounds.filter(pr => pr.roundUrl !== roundUrl),
        { roundUrl, roundToken }]
      : [{ roundUrl, roundToken }],
  }),
  updateRound: round => async state => ({
    ...state,
    round,
  }),
};

export default actions;
