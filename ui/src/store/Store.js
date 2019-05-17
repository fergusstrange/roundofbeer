import { get, set } from 'local-storage';
import { createDakpan } from 'dakpan';

const localStorageKey = 'roundOfBeerLocalStorage';

function localStorageStateOrDefaults() {
  return get(localStorageKey)
    || {
      token: undefined,
      round: undefined,
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
  updateRoundToken: roundToken => async state => ({
    ...state,
    roundToken,
    participant: '',
    participants: [],
  }),
  updateRound: round => async state => ({
    ...state,
    round,
  }),
});

export { ContextProvider, roundContext, updateLocalStore };
