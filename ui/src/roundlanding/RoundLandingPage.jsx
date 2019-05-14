import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Table,
  TableBody,
  TableCell, TableFooter,
  TableRow,
} from '@material-ui/core';
import { useContext } from '../store/Store';

export default function RoundLandingPage({ match }) {
  const [state, actions] = useContext();

  useEffect(() => {
    actions.fetchRound();
  }, [actions, match.params.roundUrl]);

  const rows = state.round
    ? state.round.participants.map(p => (
      <TableRow key={p.uuid}>
        <TableCell colSpan={1}>{p.name}</TableCell>
        <TableCell colSpan={1}>{p.round_count}</TableCell>
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
        </TableBody>
        <TableFooter>
          <TableRow>
            <TableCell colSpan={1}>Total Rounds</TableCell>
            <TableCell colSpan={1}>{roundCount}</TableCell>
          </TableRow>
        </TableFooter>
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
