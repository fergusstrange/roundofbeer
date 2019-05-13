import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Table,
  TableBody,
  TableCell,
  TableRow,
} from '@material-ui/core';
import { useContext } from '../store/Store';

export default function RoundLandingPage({ match }) {
  const [state, actions] = useContext();

  useEffect(() => {
    actions.fetchRound(state.roundToken);
  }, [actions, match.params.roundUrl, state.roundToken]);

  const rows = state.round
    ? state.round.participants.map(p => (
      <TableRow key={p.uuid}>
        <TableCell colSpan={1}>p.name</TableCell>
        <TableCell colSpan={1}>p.round_count</TableCell>
      </TableRow>
    ))
    : undefined;

  const roundCount = state.round
    ? state.round
      .participants
      .map(p => p.round_count)
      .reduce((a, b) => a + b)
    : 0;

  return (
    <div>
      <Table>
        <TableBody>
          {rows}
          <TableRow>
            <TableCell colSpan={2}>Total Rounds</TableCell>
            <TableCell colSpan={1}>{roundCount}</TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>
  );
}

RoundLandingPage.propTypes = {
  match: PropTypes.shape({
    params: PropTypes.shape({
      roundUrl: PropTypes.string.isRequired,
    }).isRequired,
  }).isRequired,
};
