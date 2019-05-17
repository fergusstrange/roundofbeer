import React from 'react';
import AddIcon from '@material-ui/icons/Add';
import { TextField } from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import PropType from 'prop-types';
import ApiClient from '../client/Client';
import { roundContext } from '../store/Store';

const client = new ApiClient();

export default function NewRoundPage({ history }) {
  const [state, actions] = roundContext();
  const participantRef = React.createRef();
  const participants = [];

  function validAndNotAlreadyExists(participant) {
    return participant
        && participants.find(element => element.toLowerCase()
        === participant.toLowerCase());
  }

  function addParticipant(e) {
    e.preventDefault();
    const participant = participantRef.current.value;
    if (validAndNotAlreadyExists(participant)) {
      participants.push(participant);
    }
  }

  function submitParticipants() {
    client.createRound(participants)
      .then(({ data }) => actions.updateRoundToken(data.token)
        .then(() => history.push(`/${data.roundUrl}`)));
  }

  return (
    <div>
      <form onSubmit={addParticipant}>
        <div>
          {participants.map(p => (
            <div key={`participantForm-div-${p}`}>
              <TextField
                key={`participantForm-textField-${p}`}
                value={p}
                disabled
              />
            </div>
          ))}
          <TextField label="name" value={state.participant} autoFocus inputRef={participantRef} />
        </div>
        <div>
          <Fab type="submit" color="primary" aria-label="Add" size="small">
            <AddIcon />
          </Fab>
          <Fab variant="extended" color="primary" aria-label="Add" onClick={submitParticipants}>Start Round</Fab>
        </div>
      </form>
    </div>
  );
}

NewRoundPage.propTypes = {
  history: PropType.shape().isRequired,
};
