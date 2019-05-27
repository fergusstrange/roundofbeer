import React, { Fragment, useState } from 'react';
import AddIcon from '@material-ui/icons/Add';
import { Container, Grid, TextField } from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import PropType from 'prop-types';
import ApiClient from '../client/Client';
import { roundContext } from '../store/Store';

const client = new ApiClient();

export default function NewRoundPage({ history }) {
  const [, actions] = roundContext();
  const [newRoundPage, updateNewRoundPage] = useState({
    createLoading: false,
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

  function updateCreateLoading() {
    updateNewRoundPage({
      ...newRoundPage,
      createLoading: true,
    });
  }

  function updateCreateLoadingFinished() {
    updateNewRoundPage({
      ...newRoundPage,
      createLoading: false,
    });
  }

  function submitParticipants(e) {
    e.preventDefault();
    updateCreateLoading();
    const participants = allParticipants();
    if (participants.length > 1) {
      client.createRound(participants)
        .then(({ data }) => actions.updateRoundToken(data)
          .then(() => updateCreateLoadingFinished())
          .then(() => history.push(`/${data.roundUrl}`)))
        .catch(() => actions.updateError('Unable to create round')
          .then(() => updateCreateLoadingFinished()));
    } else {
      actions.updateError('Need at least one friend!')
        .then(() => updateCreateLoadingFinished());
    }
  }

  return (
    <Fragment>
      <form onSubmit={submitParticipants}>
        <Grid container spacing={3} direction="column" alignItems="center">
          <Grid item xs={12}>
            <Container>
              {newRoundPage.participants.map(p => (
                <div key={`participantForm-div-${p}`}>
                  <TextField
                    component="h3"
                    key={`participantForm-textField-${p}`}
                    value={p}
                    disabled
                  />
                </div>
              ))}
            </Container>
            <Container>
              <TextField label="name" value={newRoundPage.participant} onChange={updateParticipant} autoFocus inputRef={newRoundPage.participantRef} />
              <Fab size="small" color="primary" aria-label="Add" onClick={addParticipant}>
                <AddIcon />
              </Fab>
            </Container>
          </Grid>
          <Grid item xs={12}>
            <Fab
              type="submit"
              disabled={newRoundPage.createLoading}
              variant="extended"
              color="primary"
              aria-label="Add"
            >
              Start Round
            </Fab>
          </Grid>
        </Grid>
      </form>
    </Fragment>
  );
}

NewRoundPage.propTypes = {
  history: PropType.shape().isRequired,
};
