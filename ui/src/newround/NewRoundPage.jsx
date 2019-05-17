import React from 'react';
import AddIcon from '@material-ui/icons/Add';
import { TextField } from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import PropType from 'prop-types';
import ApiClient from '../client/Client';
import { roundContext } from '../store/Store';

const client = new ApiClient();

export default function NewRoundPage({ history }) {
  const [, actions] = roundContext();
  const participantRef = React.createRef();
  const participants = [];

  function validAndNotAlreadyExists(participant) {
    return participant
        && participants.find(element => element.toLowerCase()
        === participant.toLowerCase());
  }

  function addParticipant() {
    const participant = participantRef.current.value;
    if (validAndNotAlreadyExists(participant)) {
      participants.push(participant);
      participantRef.current.value = undefined;
    }
  }

  function submitParticipants(e) {
    e.preventDefault();
    client.createRound(participants)
      .then(({ data }) => actions.updateRoundToken(data.token)
        .then(() => history.push(`/${data.roundUrl}`)))
      .then(() => actions.updateError('Unable to create round'));
  }

  return (
    <div>
      <form onSubmit={submitParticipants}>
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
          <TextField label="name" autoFocus inputRef={participantRef} />
        </div>
        <div>
          <Fab color="primary" aria-label="Add" size="small">
            <AddIcon onClick={addParticipant} />
          </Fab>
          <Fab type="submit" variant="extended" color="primary" aria-label="Add">Start Round</Fab>
        </div>
      </form>
    </div>
  );
}

NewRoundPage.propTypes = {
  history: PropType.shape().isRequired,
};
