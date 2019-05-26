import React, { Fragment } from 'react';
import {
  Grid, List, ListItem, ListItemText,
} from '@material-ui/core';
import { Link } from 'react-router-dom';
import { roundContext } from '../store/Store';

export default function OtherRoundsPage() {
  const [state] = roundContext();

  return (
    <Fragment>
      <Grid container spacing={3} direction="column" alignItems="center">
        <Grid item xs={12}>
          <List>
            {state.participatingRounds.map(pr => (
              <Link to={`/${pr.roundUrl}`}>
                <ListItem button>
                  <ListItemText primary={pr.roundUrl} />
                  {pr.participants
                    ? pr.participants.map(p => <ListItemText inset primary={p} />)
                    : undefined}
                </ListItem>
              </Link>
            ))}
          </List>
        </Grid>
      </Grid>
    </Fragment>
  );
}
