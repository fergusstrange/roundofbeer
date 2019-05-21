import React, { useState } from 'react';
import { TextField } from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import PropTypes from 'prop-types';
import Client from '../client/Client';
import { roundContext } from '../store/Store';

const client = new Client();

export default function JoinRoundPage({ match, history }) {
  const [, actions] = roundContext();
  const [joinRoundPage] = useState({
    nameRef: React.createRef(),
  });

  function submitJoinRound(event) {
    event.preventDefault();
    if (joinRoundPage.nameRef.current.value) {
      client.joinRound(match.params.roundUrl,
        joinRoundPage.nameRef.current.value)
        .then(({ data }) => actions.updateRoundToken(data.roundUrl, data.token)
          .then(() => history.push(`/${data.roundUrl}`)))
        .catch(() => actions.updateError('Unable to join that round'));
    }
  }

  return (
    <div>
      <form onSubmit={submitJoinRound}>
        <TextField label="Your name" autoFocus inputRef={joinRoundPage.nameRef} />
        <Fab type="submit" variant="extended" color="primary" aria-label="Add">
        Join Round
        </Fab>
      </form>
    </div>
  );
}

JoinRoundPage.propTypes = {
  history: PropTypes.shape().isRequired,
  match: PropTypes.shape({
    params: PropTypes.shape({
      roundUrl: PropTypes.string.isRequired,
    }).isRequired,
  }).isRequired,
};
