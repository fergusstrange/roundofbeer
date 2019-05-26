import actions from './Actions';

describe('Tests actions', () => {
  it('updates error', () => expect(actions.updateError('Uh oh!')({
    anotherProp: 12345,
  })).resolves.toEqual({
    anotherProp: 12345,
    error: 'Uh oh!',
  }));

  it('clears error', () => expect(actions.clearError()({
    error: 'Shit went wrong',
  })).resolves.toEqual({
    error: undefined,
  }));


  describe('updates round token', () => {
    it('removing round and adding participatingRounds',
      () => expect(actions.updateRoundToken({
        roundUrl: 'abcdef',
        token: 'aToken',
        participants: ['Bob'],
      })({
        anotherProp: 12345,
        round: {},
      })).resolves.toEqual({
        anotherProp: 12345,
        round: undefined,
        roundToken: 'aToken',
        participatingRounds: [{
          roundToken: 'aToken',
          roundUrl: 'abcdef',
          participants: ['Bob'],
        }],
      }));

    it('updating participatingRounds with new round',
      () => expect(actions.updateRoundToken({
        roundUrl: 'abcdef',
        token: 'aToken',
        participants: ['Bob'],
      })({
        round: {},
        anotherProp: 12345,
        participatingRounds: [{
          roundToken: 'anotherToken',
          roundUrl: 'fedcba',
        }],
      })).resolves.toEqual({
        anotherProp: 12345,
        roundToken: 'aToken',
        participatingRounds: [
          {
            roundToken: 'anotherToken',
            roundUrl: 'fedcba',
          },
          {
            roundToken: 'aToken',
            roundUrl: 'abcdef',
            participants: ['Bob'],
          },
        ],
      }));

    it('updating participatingRounds replacing existing round',
      () => expect(actions.updateRoundToken({
        roundUrl: 'abcdef',
        token: 'anUpdatedToken',
        participants: ['Bob'],
      })({
        round: {},
        anotherProp: 12345,
        participatingRounds: [
          {
            roundToken: 'aToken',
            roundUrl: 'abcdef',
          },
          {
            roundToken: 'anotherToken',
            roundUrl: 'fedcba',
          },
        ],
      })).resolves.toEqual({
        anotherProp: 12345,
        roundToken: 'anUpdatedToken',
        participatingRounds: [
          {
            roundToken: 'anotherToken',
            roundUrl: 'fedcba',
          },
          {
            roundToken: 'anUpdatedToken',
            roundUrl: 'abcdef',
            participants: ['Bob'],
          },
        ],
      }));
  });

  it('updates round', () => expect(actions.updateRound({
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
  })({
    anotherProp: 12345,
    round: {
      url: '12345',
    },
  })).resolves.toEqual({
    anotherProp: 12345,
    round: {
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
    },
  }));
});
