import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Table,
  TableBody,
  TableCell,
  TableFooter,
  TableRow, Typography,
} from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import { roundContext } from '../store/Store';
import { fetchRoundOrRedirect, nextRoundCandidate } from './roundLandingPageService';

export default function RoundLandingPage({ match, history }) {
  const [state, actions] = roundContext();

  useEffect(() => fetchRoundOrRedirect(
    state.round,
    state.participatingRounds,
    actions.updateRound,
    match.params.roundUrl,
    history,
  ), [
    state.round,
    state.participatingRounds,
    actions,
    history,
    match.params.roundUrl]);

  const ParticipantRows = () => (
    <TableBody>
      {state.round
        ? state.round.participants.map(p => (
          <TableRow key={p.uuid}>
            <TableCell colSpan={1}>{p.name}</TableCell>
            <TableCell colSpan={1}>{p.roundCount}</TableCell>
          </TableRow>
        ))
        : undefined
    }
    </TableBody>
  );

  const RoundCount = () => (
    <span>
      {state.round
        ? state.round
          .participants
          .map(p => p.roundCount)
          .reduce((a, b) => a + b)
        : 0}
    </span>
  );

  return (
    <div>
      <Typography variant="h1" component="h1">
        {`${state.round.currentCandidate.name} buys`}
      </Typography>
      <Table>
        <ParticipantRows />
        <TableFooter>
          <TableRow>
            <TableCell colSpan={1}>Total Rounds</TableCell>
            <TableCell colSpan={1}><RoundCount /></TableCell>
          </TableRow>
        </TableFooter>
      </Table>
      <Fab variant="extended" color="primary" aria-label="Add" onClick={() => nextRoundCandidate(state, actions)}>Next Buyer</Fab>
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
