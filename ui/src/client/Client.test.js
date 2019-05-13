/* eslint-disable */
/**
 * @jest-environment node
 */
import Client from './Client';
import {Pact} from '@pact-foundation/pact';
import path from 'path';
import {eachLike, somethingLike, uuid} from "@pact-foundation/pact/dsl/matchers";

jest.setTimeout(30000);

const client = new Client({
    baseURL: 'http://localhost:8888',
    withCredentials: false,
});

describe('Tests the API Client', () => {
    const provider = new Pact({
        logLevel: "debug",
        port: 8888,
        consumer: 'ui',
        provider: 'api',
        cors: true,
        log: path.resolve(process.cwd(), 'logs', 'mockserver-integration.log'),
        dir: path.resolve(process.cwd(), 'pacts'),
        spec: 2
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
                    'Accept': 'application/json'
                },
                body: {
                    participants: eachLike('Beery Little Tom')
                }
            },
            willRespondWith: {
                status: 200,
                body: {
                    token: somethingLike('tom@beer.com'),
                    round: {
                        url: somethingLike('dsakdna'),
                        participants: eachLike({
                            uuid: uuid("ce118b6e-d8e1-11e7-9296-cec278b6b50a"),
                            name: somethingLike('Tom'),
                            round_count: somethingLike(5)
                        })
                    }
                }
            }
        }));

        it('Should create round', () => client
            .createRound(['lol', 'cat'])
            .then(res => {
                expect(res.status).toEqual(200);
                expect(res.data).toEqual({
                    token: "tom@beer.com",
                    round: {
                        url: 'dsakdna',
                        participants: [{
                            uuid: "ce118b6e-d8e1-11e7-9296-cec278b6b50a",
                            name: "Tom",
                            "round_count": 5
                        }]
                    }
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
                    'Accept': 'application/json'
                },
            },
            willRespondWith: {
                status: 200,
                body: {
                    token: somethingLike('tom@beer.com'),
                    round: {
                        url: somethingLike('dsakdna'),
                        participants: eachLike({
                            uuid: uuid("ce118b6e-d8e1-11e7-9296-cec278b6b50a"),
                            name: somethingLike('Tom'),
                            round_count: somethingLike(10)
                        })
                    }
                }
            }
        }));

        it('Should fetch a round', () => client
        .fetchRound('i<3b33r')
        .then(res => {
            expect(res.status).toEqual(200);
            expect(res.data).toEqual({
                token: "tom@beer.com",
                round: {
                    url: 'dsakdna',
                    participants: [{
                        uuid: "ce118b6e-d8e1-11e7-9296-cec278b6b50a",
                        name: "Tom",
                        "round_count": 10
                    }]
                }
            });
        }));
    })
});
