import React from 'react';
import AddIcon from '@material-ui/icons/Add';
import { TextField } from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import PropType from 'prop-types';
import { useContext } from '../store/Store';

export default function NewRoundPage({ history }) {
  const [state, actions] = useContext();

  function updateParticipant(e) {
    if (e.target.value) {
      actions.updateParticipant(e.target.value);
    }
  }

  function validAndNotAlreadyExists() {
    return state.participant
        && !state.participants.find(element => element.toLowerCase()
        === state.participant.toLowerCase());
  }

  function addParticipant(e) {
    e.preventDefault();
    if (validAndNotAlreadyExists()) {
      actions.addParticipant();
    }
  }

  function submitParticipants() {
    actions.submitParticipants()
      .then(() => history.push(`/${state.round.url}`));
  }

  return (
    <div>
      <form onSubmit={addParticipant}>
        <div>
          {state.participants.map(p => (
            <div key={`participantForm-div-${p}`}>
              <TextField
                key={`participantForm-textField-${p}`}
                value={p}
                disabled
              />
            </div>
          ))}
          <TextField label="name" value={state.participant} autoFocus onChange={updateParticipant} />
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
