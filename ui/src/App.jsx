import './App.css';

import React, { Fragment, useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  BrowserRouter as Router, Link,
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
  Grid,
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
import GithubIcon from './assets/github.png';
import ClapIcon from './assets/clap.png';
import CreditsPage from './credits/CreditsPage';

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
  root: {
    display: 'flex',
    flexDirection: 'column',
    minHeight: '100vH',
  },
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
  footer: {
    padding: t.spacing(2),
    margin: '0 auto',
    marginTop: 'auto',
  },
}));

const ComponentPages = () => (
  <Switch>
    <Route path="/new-round" component={NewRoundPage} />
    <Route path="/other-rounds" component={OtherRoundsPage} />
    <Route path="/credits" component={CreditsPage} />
    <Route path="/:roundUrl/join" component={JoinRoundPage} />
    <Route path="/:roundUrl" component={RoundLandingPage} />
    <Route path="/" component={LandingPage} />
  </Switch>
);

const Footer = ({ classes }) => (
  <footer className={classes.footer}>
    <Grid container spacing={3}>
      <Grid item xs={3}>
        <a href="https://github.com/fergusstrange/roundofbeer">
          <img src={GithubIcon} alt="Github" width={16} height={16} />
        </a>
      </Grid>
      <Grid item xs={3}>
        <Link id="credits-link" to="/credits">
          <img src={ClapIcon} alt="Credits" width={16} height={16} />
        </Link>
      </Grid>
    </Grid>
  </footer>
);

Footer.propTypes = {
  classes: PropTypes.shape().isRequired,
};

function App() {
  const [state] = roundContext();
  const classes = useStyles();

  useEffect(() => updateLocalStore(state), [state]);

  return (
    <Fragment>
      <MuiThemeProvider theme={theme}>
        <CssBaseline />
        <div className={classes.root}>
          <Router>
            <AppBar elevation={0} position="absolute" className={classes.appBar}>
              <Toolbar />
            </AppBar>
            <main className={classes.layout}>
              <Paper className={classes.paper}>
                <ComponentPages />
              </Paper>
              <BottomNavigationBar />
              <ContextMessage />
            </main>
            <Footer classes={classes} />
          </Router>
        </div>
      </MuiThemeProvider>
    </Fragment>
  );
}

export default App;
