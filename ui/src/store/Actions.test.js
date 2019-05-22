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
      () => expect(actions.updateRoundToken('abcdef', 'aToken')({
        anotherProp: 12345,
        round: {},
      })).resolves.toEqual({
        anotherProp: 12345,
        round: undefined,
        roundToken: 'aToken',
        participatingRounds: [{
          roundToken: 'aToken',
          roundUrl: 'abcdef',
        }],
      }));

    it('updating participatingRounds with new round',
      () => expect(actions.updateRoundToken('abcdef', 'aToken')({
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
          },
        ],
      }));

    it('updating participatingRounds replacing existing round',
      () => expect(actions.updateRoundToken('abcdef', 'anUpdatedToken')({
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
          },
        ],
      }));
  });
});
