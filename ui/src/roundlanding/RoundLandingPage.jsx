import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Table,
  TableBody,
  TableCell,
  TableFooter,
  TableRow,
} from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import ApiClient from '../client/Client';
import { roundContext } from '../store/Store';

const client = new ApiClient();

export default function RoundLandingPage({ match, history }) {
  const [state, actions] = roundContext();

  useEffect(() => {
    function redirectJoinRoundPage() {
      history.push(`/${match.params.roundUrl}/join`);
    }

    if (!state.round && state.roundToken) {
      client.fetchRound(state.roundToken)
        .then(({ data }) => actions.updateRound(data))
        .catch(() => actions.updateError('Please rejoin that round')
          .then(redirectJoinRoundPage));
    } else if ((state.round && state.roundToken && state.round.url !== match.params.roundUrl)
      || (!state.round && !state.roundToken)) {
      redirectJoinRoundPage();
    }
  }, [actions, history, match.params.roundUrl, state.round, state.roundToken]);

  function nextRoundCandidate() {
    client.nextRoundCandidate(state.roundToken)
      .then(({ data }) => actions.updateRound(data))
      .catch(() => actions.updateError('Unable to find next buyer... Rock, Paper, Scissors?'));
  }

  const rows = state.round
    ? state.round.participants.map(p => (
      <TableRow key={p.uuid}>
        <TableCell colSpan={1}>{p.name}</TableCell>
        <TableCell colSpan={1}>{p.roundCount}</TableCell>
      </TableRow>
    ))
    : undefined;

  const roundCount = state.round
    ? state.round
      .participants
      .map(p => p.roundCount)
      .reduce((a, b) => a + b)
    : 0;

  return (
    <div>
      <Table>
        <TableBody>
          {rows}
        </TableBody>
        <TableFooter>
          <TableRow>
            <TableCell colSpan={1}>Total Rounds</TableCell>
            <TableCell colSpan={1}>{roundCount}</TableCell>
          </TableRow>
        </TableFooter>
      </Table>
      <Fab variant="extended" color="primary" aria-label="Add" onClick={nextRoundCandidate}>Next Buyer</Fab>
    </div>
  );
}

RoundLandingPage.propTypes = {
  history: PropTypes.shape().isRequired,
  match: PropTypes.shape({
    params: PropTypes.shape({
      roundUrl: PropTypes.string.isRequired,
    }).isRequired,
  }).isRequired,
};
