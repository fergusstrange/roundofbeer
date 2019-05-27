import React, { Fragment, useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import {
  Table,
  TableBody,
  TableCell,
  TableFooter,
  TableRow,
  Typography,
  Divider, Grid, CircularProgress,
} from '@material-ui/core';
import Fab from '@material-ui/core/Fab';
import { roundContext } from '../store/Store';
import { fetchRoundOrRedirect, nextRoundCandidate } from './roundLandingPageService';

export default function RoundLandingPage({ match, history }) {
  const [state, actions] = roundContext();
  const [roundLandingPage, updateRoundLandingPage] = useState({
    nextCandidateLoading: false,
  });

  useEffect(() => {
    updateRoundLandingPage({ pageLoading: !state.round });
  }, [state.round]);

  useEffect(() => fetchRoundOrRedirect(
    state.round,
    state.participatingRounds,
    actions,
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

  const CurrentCandidate = () => (
    <Typography variant="h3" component="h3">
      {state.round
        ? `${state.round.currentCandidate.name} buys`
        : undefined}
    </Typography>
  );

  const RoundDetails = () => (
    <Fragment>
      <Grid item xs={12}>
        <CurrentCandidate />
      </Grid>
      <Grid item xs={12}>
        <Divider variant="middle" component="h2" />
      </Grid>
      <Grid item xs={12}>
        <Table>
          <ParticipantRows />
          <TableFooter>
            <TableRow>
              <TableCell colSpan={1}>Total Rounds</TableCell>
              <TableCell colSpan={1}><RoundCount /></TableCell>
            </TableRow>
          </TableFooter>
        </Table>
      </Grid>
      <Grid item xs={12}>
        <Fab
          variant="extended"
          color="primary"
          aria-label="Add"
          disabled={!!roundLandingPage.nextCandidateLoading}
          onClick={() => nextRoundCandidate(state, actions, updateRoundLandingPage)}
        >
Next Buyer
        </Fab>
      </Grid>
    </Fragment>
  );

  return (
    <Fragment>
      <Grid container spacing={3} direction="column" alignItems="center" color="action">
        {state.round
          ? <RoundDetails />
          : <CircularProgress color="inherit" />}
      </Grid>
    </Fragment>
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
