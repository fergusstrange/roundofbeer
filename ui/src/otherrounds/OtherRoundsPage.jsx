import React, { Fragment } from 'react';
import PropTypes from 'prop-types';
import {
  Grid, List, ListItem, ListItemText,
} from '@material-ui/core';
import { Link } from 'react-router-dom';
import { roundContext } from '../store/Store';

export default function OtherRoundsPage({ history }) {
  const [state, actions] = roundContext();

  function updateRoundToken(e, participatingRound) {
    e.preventDefault();
    actions.clearRoundAndToken()
      .then(() => history.push(`/${participatingRound.roundUrl}`));
  }

  return (
    <Fragment>
      <Grid container spacing={3} direction="column" alignItems="flex-start">
        <Grid item xs={12}>
          <List>
            {state.participatingRounds.map(pr => (
              <Fragment key={pr.roundUrl}>
                <Link to={`/${pr.roundUrl}`} onClick={e => updateRoundToken(e, pr)}>
                  <ListItem button>
                    <ListItemText primary={pr.roundUrl} />
                  </ListItem>
                </Link>
                {pr.participants
                  ? pr.participants.map(p => (
                    <ListItem key={`${pr.roundUrl}-${p}`}>
                      <ListItemText inset primary={p} />
                    </ListItem>
                  ))
                  : undefined}
              </Fragment>
            ))}
          </List>
        </Grid>
      </Grid>
    </Fragment>
  );
}

OtherRoundsPage.propTypes = {
  history: PropTypes.shape().isRequired,
};
