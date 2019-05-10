import React from 'react';
import AddIcon from '@material-ui/icons/Add';
import { TextField } from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import PropTypes from 'prop-types';
import { actions, connect } from '../store/Store';

const NewRoundPage = ({ participant, participants }) => {
  function updateParticipant(e) {
    if (e.target.value) {
      actions.updateParticipant(e.target.value);
    }
  }

  function validAndNotAlreadyExists() {
    return participant
        && !participants.find(element => element.toLowerCase() === participant.toLowerCase());
  }

  function addParticipant(e) {
    e.preventDefault();
    if (validAndNotAlreadyExists()) {
      actions.addParticipant();
    }
  }

  function submitParticipants() {
    actions.submitParticipants();
  }

  return (
    <div>
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
            <TextField label="name" value={participant} autoFocus onChange={updateParticipant} />
          </div>
          <div>
            <Fab type="submit" color="primary" aria-label="Add" size="small">
              <AddIcon />
            </Fab>
            <Fab variant="extended" color="primary" aria-label="Add" onClick={submitParticipants}>Start Round</Fab>
          </div>
        </form>
      </div>
    </div>
  );
};

NewRoundPage.propTypes = {
  participant: PropTypes.string.isRequired,
  participants: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default connect(state => state)(NewRoundPage);
