/**
 * @jest-environment node
 */
import { Pact } from '@pact-foundation/pact';
import path from 'path';
import { eachLike, somethingLike, uuid } from '@pact-foundation/pact/dsl/matchers';
import Client from './Client';

jest.setTimeout(30000);

const client = new Client({
  baseURL: 'http://localhost:8888',
  withCredentials: false,
});

describe('Tests the API Client', () => {
  const provider = new Pact({
    logLevel: 'debug',
    port: 8888,
    consumer: 'ui',
    provider: 'api',
    cors: true,
    log: path.resolve(process.cwd(), 'logs', 'mockserver-integration.log'),
    dir: path.resolve(process.cwd(), 'pacts'),
    spec: 2,
  });

  beforeAll(() => provider.setup());

  afterAll(() => provider.finalize());

  describe('Create Round', () => {
    beforeEach(() => provider.addInteraction({
      state: 'round does not exist',
      uponReceiving: 'a valid create round request',
      withRequest: {
        method: 'POST',
        path: '/round',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
        },
        body: {
          participants: eachLike('Beery Little Tom'),
        },
      },
      willRespondWith: {
        status: 200,
        body: {
          token: somethingLike('tom@beer.com'),
          roundUrl: somethingLike('aUrl'),
          participants: eachLike('Bob'),
        },
      },
    }));

    it('Should create round', () => client
      .createRound(['lol', 'cat'])
      .then((res) => {
        expect(res.status).toEqual(200);
        expect(res.data).toEqual({
          token: 'tom@beer.com',
          roundUrl: 'aUrl',
          participants: ['Bob'],
        });
      }));
  });

  describe('Join Round', () => {
    beforeEach(() => provider.addInteraction({
      state: 'round exists',
      uponReceiving: 'a request to join the round',
      withRequest: {
        method: 'POST',
        path: '/round/das8db',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
        },
        body: {
          name: somethingLike('Tom'),
        },
      },
      willRespondWith: {
        status: 200,
        body: {
          token: somethingLike('tom@beer.com'),
          roundUrl: somethingLike('aUrl'),
          participants: eachLike('Bob'),
        },
      },
    }));

    it('Should join round', () => client
      .joinRound('das8db', 'Tom')
      .then((res) => {
        expect(res.status).toEqual(200);
        expect(res.data).toEqual({
          token: 'tom@beer.com',
          roundUrl: 'aUrl',
          participants: ['Bob'],
        });
      }));
  });

  describe('Fetch Round', () => {
    beforeEach(() => provider.addInteraction({
      state: 'round exists',
      uponReceiving: 'a valid roundToken header',
      withRequest: {
        method: 'GET',
        path: '/round',
        headers: {
          'x-round-token': somethingLike('i<3b33r'),
          Accept: 'application/json',
        },
      },
      willRespondWith: {
        status: 200,
        body: {
          url: somethingLike('dsakdna'),
          participants: eachLike({
            uuid: uuid('ce118b6e-d8e1-11e7-9296-cec278b6b50a'),
            name: somethingLike('Tom'),
            roundCount: somethingLike(10),
          }),
          currentCandidate: {
            uuid: uuid('ce118b6e-d8e1-11e7-9296-cec278b6b50a'),
            name: somethingLike('Tom'),
            roundCount: somethingLike(10),
          },

        },
      },
    }));

    it('Should fetch a round', () => client
      .fetchRound('i<3b33r')
      .then((res) => {
        expect(res.status).toEqual(200);
        expect(res.data).toEqual({
          url: 'dsakdna',
          participants: [{
            uuid: 'ce118b6e-d8e1-11e7-9296-cec278b6b50a',
            name: 'Tom',
            roundCount: 10,
          }],
          currentCandidate: {
            uuid: 'ce118b6e-d8e1-11e7-9296-cec278b6b50a',
            name: 'Tom',
            roundCount: 10,
          },
        });
      }));
  });

  describe('Gets Next Round', () => {
    beforeEach(() => provider.addInteraction({
      state: 'An inflight round exists',
      uponReceiving: 'a valid roundToken header',
      withRequest: {
        method: 'PUT',
        path: '/round',
        headers: {
          'x-round-token': somethingLike('i<3b33r'),
          Accept: 'application/json',
        },
      },
      willRespondWith: {
        status: 200,
        body: {
          url: somethingLike('dsakdna'),
          participants: eachLike({
            uuid: uuid('5559be5c-2d73-446b-a3f8-da14d7c7f5a6'),
            name: somethingLike('Geoff'),
            roundCount: somethingLike(11),
          }),
          currentCandidate: {
            uuid: uuid('5559be5c-2d73-446b-a3f8-da14d7c7f5a6'),
            name: somethingLike('Geoff'),
            roundCount: somethingLike(11),
          },

        },
      },
    }));

    it('Should update current candidate and fetch updated round', () => client
      .nextRoundCandidate('i<3b33r')
      .then((res) => {
        expect(res.status).toEqual(200);
        expect(res.data).toEqual({
          url: 'dsakdna',
          participants: [
            {
              uuid: '5559be5c-2d73-446b-a3f8-da14d7c7f5a6',
              name: 'Geoff',
              roundCount: 11,
            }],
          currentCandidate: {
            uuid: '5559be5c-2d73-446b-a3f8-da14d7c7f5a6',
            name: 'Geoff',
            roundCount: 11,
          },
        });
      }));
  });
});
