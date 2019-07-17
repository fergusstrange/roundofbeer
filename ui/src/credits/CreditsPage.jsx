import React, { Fragment } from 'react';
import { Typography, List, ListItem } from '@material-ui/core';

export default function CreditsPage() {
  return (
    <Fragment>
      <Typography id="credits-title" variant="h6">Icons and Images</Typography>
      <List component="div">
        <ListItem component="div">
          <a href="https://www.flaticon.com/authors/ddara" title="dDara">dDara</a>
        </ListItem>
        <ListItem component="div">
          <a href="https://www.flaticon.com/authors/smashicons" title="Smashicons">Smashicons</a>
        </ListItem>
        <ListItem component="div">
          <a href="https://www.freepik.com/" title="Freepik">Freepik</a>
        </ListItem>
      </List>
    </Fragment>
  );
}
