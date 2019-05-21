import './App.css';

import React, { Fragment, useEffect } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { Grid } from '@material-ui/core';
import LandingPage from './landing/LandingPage';
import RoundLandingPage from './roundlanding/RoundLandingPage';
import NewRoundPage from './newround/NewRoundPage';
import ContextMessage from './contextmessage/ContextMessage';
import JoinRoundPage from './joinround/JoinRoundPage';
import { roundContext, updateLocalStore } from './store/Store';

function App() {
  const [state] = roundContext();
  useEffect(() => updateLocalStore(state), [state]);

  return (
    <Fragment>
      <Grid container spacing={24} justify="center" alignItems="center" alignContent="center" style={{ marginTop: '25px' }}>
        <Router>
          <Switch>
            <Route path="/new-round" component={NewRoundPage} />
            <Route path="/:roundUrl/join" component={JoinRoundPage} />
            <Route path="/:roundUrl" component={RoundLandingPage} />
            <Route path="/" component={LandingPage} />
          </Switch>
        </Router>
        <ContextMessage />
      </Grid>
    </Fragment>
  );
}

export default App;
