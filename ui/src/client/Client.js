import axios from 'axios';

const defaultClient = axios.create();
const defaultHeaders = {
  Accept: 'application/json',
  'Content-Type': 'application/json',
};

function defaultOptions(roundToken, overrides) {
  return Object.assign({},
    {
      baseURL: process.env.REACT_APP_API_URL || 'https://api.roundof.beer',
      withCredentials: true,
      headers: roundToken
        ? Object.assign({}, defaultHeaders, { 'x-round-token': roundToken })
        : defaultHeaders,
    },
    overrides);
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
    });
  }

  async joinRound(roundId, name) {
    return this.withDefaultOptions({
      method: 'POST',
      url: `/round/${roundId}`,
      data: { name },
    });
  }

  async fetchRound(roundToken) {
    return this.withDefaultOptions({
      method: 'GET',
      url: '/round',
    }, roundToken);
  }

  async nextRoundCandidate(roundToken) {
    return this.withDefaultOptions({
      method: 'PUT',
      url: '/round',
    }, roundToken);
  }
}
