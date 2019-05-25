import './App.css';

import React, { Fragment, useEffect } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import {
  createMuiTheme,
  AppBar,
  CssBaseline,
  Paper,
  Toolbar,
} from '@material-ui/core';
import { makeStyles, MuiThemeProvider } from '@material-ui/core/styles';
import LandingPage from './landing/LandingPage';
import RoundLandingPage from './roundlanding/RoundLandingPage';
import NewRoundPage from './newround/NewRoundPage';
import ContextMessage from './contextmessage/ContextMessage';
import JoinRoundPage from './joinround/JoinRoundPage';
import { roundContext, updateLocalStore } from './store/Store';
import beerBackgroundImage from './assets/beer.png';

const theme = createMuiTheme({
  palette: {
    primary: {
      main: '#fff59d',
    },
    secondary: {
      main: '#ffca28',
    },
    background: {
      default: '#f4cc37',
      paper: '#fef5d8',
    },
  },
});

const useStyles = makeStyles(t => ({
  appBar: {
    backgroundImage: `url("${beerBackgroundImage}")`,
    backgroundRepeat: 'repeat-x',
    position: 'relative',
  },
  layout: {
    width: 'auto',
    marginLeft: t.spacing(2),
    marginRight: t.spacing(2),
    [t.breakpoints.up(600 + t.spacing(2) * 2)]: {
      width: 600,
      marginLeft: 'auto',
      marginRight: 'auto',
    },
  },
  paper: {
    marginTop: t.spacing(3),
    marginBottom: t.spacing(3),
    padding: t.spacing(3),
    [t.breakpoints.up(600 + t.spacing(3) * 2)]: {
      marginTop: t.spacing(6),
      marginBottom: t.spacing(6),
      padding: t.spacing(3),
    },
  },
  appBarTitle: {
    margin: '0 auto',
  },
}));

function App() {
  const [state] = roundContext();
  const classes = useStyles();

  useEffect(() => updateLocalStore(state), [state]);

  return (
    <Fragment>
      <MuiThemeProvider theme={theme}>
        <CssBaseline />
        <AppBar elevation={0} position="absolute" className={classes.appBar}>
          <Toolbar className={classes.appBarTitle} />
        </AppBar>
        <main className={classes.layout}>
          <Paper className={classes.paper}>
            <Router>
              <Switch>
                <Route path="/new-round" component={NewRoundPage} />
                <Route path="/:roundUrl/join" component={JoinRoundPage} />
                <Route path="/:roundUrl" component={RoundLandingPage} />
                <Route path="/" component={LandingPage} />
              </Switch>
            </Router>
          </Paper>
          <ContextMessage />
        </main>
      </MuiThemeProvider>
    </Fragment>
  );
}

export default App;
