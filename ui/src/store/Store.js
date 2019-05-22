import { get, set } from 'local-storage';
import { createDakpan } from 'dakpan';
import actions from './Actions';

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

const [ContextProvider, roundContext] = createDakpan(localStorageStateOrDefaults())(actions);

export { ContextProvider, roundContext, updateLocalStore };
