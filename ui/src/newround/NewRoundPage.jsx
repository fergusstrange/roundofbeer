import React, { useState } from 'react';
import AddIcon from '@material-ui/icons/Add';
import { TextField } from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import PropType from 'prop-types';
import ApiClient from '../client/Client';
import { roundContext } from '../store/Store';

const client = new ApiClient();

export default function NewRoundPage({ history }) {
  const [, actions] = roundContext();
  const [newRoundPage, updateNewRoundPage] = useState({
    participantRef: React.createRef(),
    participants: [],
    participant: '',
  });

  function validAndNotAlreadyExists(participant) {
    return participant
        && !newRoundPage.participants
          .find(element => element.toLowerCase() === participant.toLowerCase());
  }

  function updateParticipant() {
    updateNewRoundPage({
      ...newRoundPage,
      participant: newRoundPage.participantRef.current.value,
    });
  }

  function addParticipant() {
    if (validAndNotAlreadyExists(newRoundPage.participant)) {
      updateNewRoundPage({
        ...newRoundPage,
        participants: [...newRoundPage.participants, newRoundPage.participant],
        participant: '',
      });
      newRoundPage.participantRef.current.focus();
    }
  }

  function allParticipants() {
    if (newRoundPage.participant && validAndNotAlreadyExists(newRoundPage.participant)) {
      return [...newRoundPage.participants, newRoundPage.participant];
    }
    return newRoundPage.participants;
  }

  function submitParticipants(e) {
    e.preventDefault();
    const participants = allParticipants();
    if (participants.length > 1) {
      client.createRound(participants)
        .then(({ data }) => actions.updateRoundToken(data.roundUrl, data.token)
          .then(() => history.push(`/${data.roundUrl}`)))
        .catch(() => actions.updateError('Unable to create round'));
    } else {
      actions.updateError('Need at least one friend!');
    }
  }

  return (
    <div>
      <form onSubmit={submitParticipants}>
        <div>
          {newRoundPage.participants.map(p => (
            <div key={`participantForm-div-${p}`}>
              <TextField
                key={`participantForm-textField-${p}`}
                value={p}
                disabled
              />
            </div>
          ))}
          <TextField label="name" value={newRoundPage.participant} onChange={updateParticipant} autoFocus inputRef={newRoundPage.participantRef} />
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
