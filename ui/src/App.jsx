import './App.css';

import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import LandingPage from './landing/LandingPage';
import RoundLandingPage from './roundlanding/RoundLandingPage';
import NewRoundPage from './newround/NewRoundPage';
import { Provider } from './store/Store';

function App() {
  return (
    <Provider>
      <Router>
        <Switch>
          <Route path="/new-round" component={NewRoundPage} />
          <Route path="/:roundUrl" component={RoundLandingPage} />
          <Route path="/" component={LandingPage} />
        </Switch>
      </Router>
    </Provider>
  );
}

export default App;
