import { get, set } from 'local-storage';
import { createDakpan } from 'dakpan';

const localStorageKey = 'roundOfBeerLocalStorage';

function localStorageStateOrDefaults() {
  return get(localStorageKey)
    || {
      roundToken: undefined,
      round: undefined,
      participatingRounds: [],
    };
}

function updateLocalStore(state) {
  set(localStorageKey, state);
}

const [ContextProvider, roundContext] = createDakpan(localStorageStateOrDefaults())({
  updateError: error => async state => ({
    ...state,
    error,
  }),
  clearError: () => async state => ({
    ...state,
    error: undefined,
  }),
  updateRoundToken: (roundUrl, roundToken) => async state => ({
    ...state,
    roundToken,
    round: undefined,
    participatingRounds: Array.isArray(state.participatingRounds)
      ? [...state.participatingRounds.filter(pr => pr.roundUrl !== roundUrl),
        { roundUrl, roundToken }]
      : [{ roundUrl, roundToken }],
  }),
  updateRound: round => async state => ({
    ...state,
    round,
  }),
});

export { ContextProvider, roundContext, updateLocalStore };
