import './App.css';

import React, { Fragment, useEffect } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { Grid } from '@material-ui/core';
import LandingPage from './landing/LandingPage';
import RoundLandingPage from './roundlanding/RoundLandingPage';
import NewRoundPage from './newround/NewRoundPage';
import ContextMessage from './contextmessage/ContextMessage';
import { useContext, updateLocalStore } from './store/Store';

function App() {
  const [state] = useContext();
  useEffect(() => updateLocalStore(state), [state]);

  return (
    <Fragment>
      <Grid container spacing={24} justify="center" alignItems="center" alignContent="center">
        <Router>
          <Switch>
            <Route path="/new-round" component={NewRoundPage} />
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
