import './App.css';

import React, { Fragment, useEffect } from 'react';
import {
  BrowserRouter as Router,
  Route,
  Switch,
} from 'react-router-dom';
import {
  createMuiTheme,
  AppBar,
  CssBaseline,
  Paper,
  Toolbar,
  makeStyles,
} from '@material-ui/core';
import { MuiThemeProvider } from '@material-ui/core/styles';
import LandingPage from './landing/LandingPage';
import RoundLandingPage from './roundlanding/RoundLandingPage';
import NewRoundPage from './newround/NewRoundPage';
import OtherRoundsPage from './otherrounds/OtherRoundsPage';
import ContextMessage from './contextmessage/ContextMessage';
import JoinRoundPage from './joinround/JoinRoundPage';
import BottomNavigationBar from './navigation/BottomNavigationBar';
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
    marginTop: t.spacing(6),
    marginBottom: t.spacing(3),
    padding: t.spacing(3),
    [t.breakpoints.up(600 + t.spacing(3) * 2)]: {
      marginTop: t.spacing(6),
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
                <Route path="/other-rounds" component={OtherRoundsPage} />
                <Route path="/:roundUrl/join" component={JoinRoundPage} />
                <Route path="/:roundUrl" component={RoundLandingPage} />
                <Route path="/" component={LandingPage} />
              </Switch>
            </Paper>
            <BottomNavigationBar />
            <ContextMessage />
          </main>
        </Router>
      </MuiThemeProvider>
    </Fragment>
  );
}

export default App;
