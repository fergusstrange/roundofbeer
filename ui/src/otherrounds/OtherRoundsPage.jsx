import React, { Fragment } from 'react';
import PropTypes from 'prop-types';
import {
  Grid,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Typography,
} from '@material-ui/core';
import RemoveCircleIcon from '@material-ui/icons/RemoveCircle';
import ArrowForwardIcon from '@material-ui/icons/ArrowForward';
import { roundContext } from '../store/Store';

export default function OtherRoundsPage({ history }) {
  const [state, actions] = roundContext();

  function updateRoundToken(e, participatingRound) {
    e.preventDefault();
    actions.clearRoundAndUpdateToken(participatingRound.roundToken)
      .then(() => history.push(`/${participatingRound.roundUrl}`));
  }

  function removeParticipatingRound(e, pr) {
    e.preventDefault();
    actions.removeParticipatingRound(pr.roundUrl);
  }

  const NoOtherRounds = () => (
    <Grid container spacing={3} direction="column" alignItems="center">
      <Grid item xs={12}>
        <Typography variant="h5" className="other-rounds-page-no-rounds">No other rounds</Typography>
      </Grid>
    </Grid>
  );

  const OtherRoundsList = () => (
    <Grid container spacing={3} direction="column" alignItems="flex-start">
      <Grid item xs={12}>
        <List>
          {state.participatingRounds.map(pr => (
            <Fragment key={pr.roundUrl}>
              <ListItem>
                <ListItem button onClick={e => updateRoundToken(e, pr)}>
                  <ListItemIcon>
                    <ArrowForwardIcon />
                  </ListItemIcon>
                  <ListItemText primary={pr.roundUrl} />
                </ListItem>
                <ListItem button onClick={e => removeParticipatingRound(e, pr)}>
                  <RemoveCircleIcon className="other-rounds-remove-round" />
                </ListItem>
              </ListItem>
              {pr.participants
                ? pr.participants.map(p => (
                  <ListItem key={`${pr.roundUrl}-${p}`}>
                    <ListItemText
                      className="other-rounds-participant"
                      inset
                      primary={p}
                    />
                  </ListItem>
                ))
                : undefined}
            </Fragment>
          ))}
        </List>
      </Grid>
    </Grid>
  );

  return (
    <Fragment>
      {state.participatingRounds && state.participatingRounds.length > 0
        ? <OtherRoundsList />
        : <NoOtherRounds />
      }
    </Fragment>
  );
}

OtherRoundsPage.propTypes = {
  history: PropTypes.shape().isRequired,
};
