import axios from 'axios';

const defaultClient = axios.create();
const defaultHeaders = {
  Accept: 'application/json',
  'Content-Type': 'application/json',
};

export default class ApiClient {
  constructor(host) {
    this.withDefaultOptions = (furtherOptions, roundToken) => defaultClient(
      Object.assign({},
        {
          baseURL: host || 'https://api.roundof.beer',
          headers: roundToken
            ? Object.assign({}, defaultHeaders, { 'x-round-token': roundToken })
            : defaultHeaders,
        },
        furtherOptions),
    );
  }

  createRound(participants) {
    return this.withDefaultOptions({
      method: 'POST',
      url: '/round',
      data: { participants },
    });
  }

  joinRound(roundId, name) {
    return this.withDefaultOptions({
      method: 'POST',
      url: `/round/${roundId}`,
      data: { name },
    });
  }

  fetchRound(roundToken) {
    return this.withDefaultOptions({
      method: 'GET',
      url: '/round',
    }, roundToken);
  }

  nextRoundCandidate(roundToken) {
    return this.withDefaultOptions({
      method: 'PUT',
      url: '/round',
    }, roundToken);
  }
}
