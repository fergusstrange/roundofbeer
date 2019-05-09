import React from 'react';
import AddIcon from '@material-ui/icons/Add';
import { TextField } from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import PropTypes from 'prop-types';
import {
  actions, connect,
} from '../store/Store';

const NewRoundPage = ({ participant, participants }) => {
  const addParticipant = (e) => {
    e.preventDefault();
    if (participants.filter(p => p.toLowerCase().equals(participant.toLowerCase())).length === 0) {
      actions.addParticipant();
    }
  };

  const updateParticipant = (e) => {
    if (e.target.value) {
      actions.updateParticipant(e.target.value);
    }
  };

  return (
    <div>
      <div>
        {participants.map(name => (<div key={name}>{name}</div>))}
      </div>
      <div>
        <form onSubmit={addParticipant}>
          <div>
            <TextField label="name" value={participant} onChange={updateParticipant} />
            <Fab type="submit" color="primary" aria-label="Add" size="small">
              <AddIcon />
            </Fab>
          </div>
          <div>
            <Fab variant="extended" color="primary" aria-label="Add" onClick={actions.submitParticipants}>Start Round</Fab>
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
