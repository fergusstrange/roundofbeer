import { get, set } from 'local-storage';
import { createDakpan } from 'dakpan';
import ApiClient from '../client/Client';

const localStorageKey = 'roundOfBeerLocalStorage';
const client = new ApiClient();

function localStorageStateOrDefaults() {
  return get(localStorageKey)
    || {
      token: undefined,
      round: undefined,
      participant: '',
      participants: [],
    };
}

function updateLocalStore(state) {
  set(localStorageKey, state);
}

const [ContextProvider, useContext] = createDakpan(localStorageStateOrDefaults())({
  clearError: () => async state => ({
    ...state,
    error: undefined,
  }),
  addParticipant: () => ({ participant, ...state }) => ({
    ...state,
    participants: [
      ...state.participants,
      participant,
    ],
    participant: '',
  }),
  updateParticipant: participant => state => ({
    ...state,
    participant,
  }),
  submitParticipants: () => async (state) => {
    const { error } = await client.createRound(state.participants);
    return !error
      ? ({
        ...state,
        participant: '',
        participants: [],
      })
      : ({
        ...state,
        error,
      });
  },
  fetchRound: roundToken => async (state) => {
    const { error, data } = await client.fetchRound(roundToken);
    return error
      ? ({
        ...state,
        error,
      })
      : ({
        ...state,
        round: data.round,
        roundToken: data.token,
      });
  },
});

export { ContextProvider, useContext, updateLocalStore };
