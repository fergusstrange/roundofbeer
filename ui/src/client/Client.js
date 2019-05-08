import axios from 'axios';

const defaultClient = axios.create();
const defaultHeaders = {
  Accept: 'application/json',
  'Content-Type': 'application/json',
};

export default class ApiClient {
  constructor() {
    this.withDefaultOptions = (furtherOptions, roundToken) => defaultClient(
      Object.assign({},
        {
          baseURL: 'https://api.roundof.beer',
          headers: roundToken
            ? Object.assign({}, defaultHeaders, { 'x-round-token': roundToken })
            : defaultHeaders,
          withCredentials: true,
        },
        furtherOptions),
    );
  }

  createRound(participants) {
    this.withDefaultOptions({
      method: 'POST',
      path: '/round',
      data: { participants },
    });
  }

  joinRound(roundId, name) {
    this.withDefaultOptions({
      method: 'POST',
      path: `/round/${roundId}`,
      data: { name },
    });
  }

  fetchRound(roundToken) {
    this.withDefaultOptions({
      method: 'GET',
      path: '/round',
    }, roundToken);
  }

  nextRoundCandidate(roundToken) {
    this.withDefaultOptions({
      method: 'PUT',
      path: '/round',
    }, roundToken);
  }
}
