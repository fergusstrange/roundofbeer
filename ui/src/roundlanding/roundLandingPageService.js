import ApiClient from '../client/Client';

const client = new ApiClient();

function redirectJoinRoundPage(history, roundUrl) {
  history.push(`/${roundUrl}/join`);
}

function participatingRoundsOrEmpty(participatingRounds) {
  return Array.isArray(participatingRounds)
    ? participatingRounds
    : [];
}

function filterForMatchingRound(pathRoundUrl) {
  return participatingRound => participatingRound.roundUrl === pathRoundUrl;
}

const fetchRoundOrRedirect = (round, participatingRounds, updateRound, pathRoundUrl, history) => {
  if (!round) {
    const existingRound = participatingRoundsOrEmpty(participatingRounds)
      .find(filterForMatchingRound(pathRoundUrl));
    if (existingRound) {
      client.fetchRound(existingRound.roundToken)
        .then(({ data }) => updateRound(data));
    } else {
      redirectJoinRoundPage(history, pathRoundUrl);
    }
  }
};

const nextRoundCandidate = (state, actions) => {
  client.nextRoundCandidate(state.roundToken)
    .then(({ data }) => actions.updateRound(data))
    .catch(() => actions.updateError('Unable to find next buyer... Rock, Paper, Scissors?'));
};

export { fetchRoundOrRedirect, nextRoundCandidate };
