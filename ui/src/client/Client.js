import axios from 'axios';

const defaultClient = axios.create();
const defaultHeaders = {
  Accept: 'application/json',
  'Content-Type': 'application/json',
};

function defaultOptions(roundToken, overrides) {
  return Object.assign({},
    {
      baseURL: 'https://api.roundof.beer',
      withCredentials: true,
      headers: roundToken
        ? Object.assign({}, defaultHeaders, { 'x-round-token': roundToken })
        : defaultHeaders,
    },
    overrides);
}

function errorWithMessage(message) {
  return () => ({
    error: message,
  });
}

export default class ApiClient {
  constructor(overrides) {
    this.withDefaultOptions = (furtherOptions, roundToken) => defaultClient(
      Object.assign({},
        defaultOptions(roundToken, overrides),
        furtherOptions),
    );
  }

  async createRound(participants) {
    return this.withDefaultOptions({
      method: 'POST',
      url: '/round',
      data: { participants },
    })
      .catch(errorWithMessage('Unable to create a new round'));
  }

  async joinRound(roundId, name) {
    return this.withDefaultOptions({
      method: 'POST',
      url: `/round/${roundId}`,
      data: { name },
    })
      .catch(errorWithMessage('Unable to join that round'));
  }

  async fetchRound(roundToken) {
    return this.withDefaultOptions({
      method: 'GET',
      url: '/round',
    }, roundToken)
      .catch(errorWithMessage('Unable to load that round'));
  }

  async nextRoundCandidate(roundToken) {
    return this.withDefaultOptions({
      method: 'PUT',
      url: '/round',
    }, roundToken)
      .catch(errorWithMessage('Unable to find next buyer... Rock, Paper, Scissors?'));
  }
}
