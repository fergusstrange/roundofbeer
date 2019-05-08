import createStore from 'react-waterfall';
import ApiClient from '../client/Client';

const client = new ApiClient();

const config = {
  initialState: {
    participant: '',
    participants: [],
  },
  actionsCreators: {
    addParticipant: async ({ participants, participant }) => ({
      participants: [
        ...participants,
        participant,
      ],
      participant: '',
    }),
    updateParticipant: async (_x, _e, participant) => ({ participant }),
    submitParticipants: async ({ participants }) => {
      await client.createRound(participants);
      return ({
        participant: '',
        participants: [],
      });
    },
  },
};

export const {
  Provider,
  Consumer,
  actions,
  getState,
  connect,
  subscribe,
} = createStore(config);
