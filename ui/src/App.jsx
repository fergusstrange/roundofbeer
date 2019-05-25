import './App.css';

import React, { Fragment, useEffect } from 'react';
import {
  BrowserRouter as Router, Link, Route, Switch,
} from 'react-router-dom';
import {
  createMuiTheme,
  AppBar,
  CssBaseline,
  Fab,
  Paper,
  Toolbar,
  makeStyles,
} from '@material-ui/core';
import { MuiThemeProvider } from '@material-ui/core/styles';
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
      main: '#fef5d8',
    },
    background: {
      default: '#f4cc37',
      paper: '#fef5d8',
    },
  },
});

const useStyles = makeStyles(t => ({
  appBar: {
    backgroundColor: '#f4cc37',
    backgroundImage: `url("${beerBackgroundImage}")`,
    backgroundRepeat: 'repeat-x',
    position: 'relative',
  },
  bottomAppBar: {
    top: 'auto',
    bottom: 0,
  },
  bottomAppBarIcon: {
    fontSize: '10px',
    position: 'absolute',
    zIndex: 1,
    top: -30,
    left: 0,
    right: 0,
    margin: '0 auto',
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
    marginTop: t.spacing(9),
    marginBottom: t.spacing(3),
    padding: t.spacing(3),
    [t.breakpoints.up(600 + t.spacing(3) * 2)]: {
      marginTop: t.spacing(9),
      marginBottom: t.spacing(3),
      padding: t.spacing(3),
    },
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
          <Toolbar />
        </AppBar>
        <Router>
          <main className={classes.layout}>
            <Paper className={classes.paper}>
              <Switch>
                <Route path="/new-round" component={NewRoundPage} />
                <Route path="/:roundUrl/join" component={JoinRoundPage} />
                <Route path="/:roundUrl" component={RoundLandingPage} />
                <Route path="/" component={LandingPage} />
              </Switch>
            </Paper>
            <AppBar position="fixed" color="secondary" className={classes.bottomAppBar}>
              <Toolbar>
                <Link to="/new-round">
                  <Fab color="primary" aria-label="New Round" className={classes.bottomAppBarIcon}>
                  New Round
                  </Fab>
                </Link>
              </Toolbar>
            </AppBar>
            <ContextMessage />
          </main>
        </Router>
      </MuiThemeProvider>
    </Fragment>
  );
}

export default App;
