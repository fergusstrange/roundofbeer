import ApiClient from '../client/Client';

const client = new ApiClient();

function participatingRoundsOrEmpty(participatingRounds) {
  return Array.isArray(participatingRounds)
    ? participatingRounds
    : [];
}

function filterForMatchingRound(pathRoundUrl) {
  return participatingRound => participatingRound.roundUrl === pathRoundUrl;
}

function fetchExistingRound(participatingRounds, pathRoundUrl) {
  return participatingRoundsOrEmpty(participatingRounds).find(filterForMatchingRound(pathRoundUrl));
}

function updateNextCandidateFinished(updateRoundLandingPage) {
  return updateRoundLandingPage({ nextCandidateLoading: false });
}

const nextRoundCandidate = (state, actions, updateRoundLandingPage) => {
  updateRoundLandingPage({ nextCandidateLoading: true });
  client.nextRoundCandidate(state.roundToken)
    .then(({ data }) => actions.updateRound(data))
    .then(() => updateNextCandidateFinished(updateRoundLandingPage))
    .catch(() => actions.updateError('Unable to find next buyer... Rock, Paper, Scissors?')
      .then(() => updateNextCandidateFinished(updateRoundLandingPage)));
};

export {
  nextRoundCandidate,
  fetchExistingRound,
};
