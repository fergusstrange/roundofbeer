import React, { Fragment, useState } from 'react';
import {
  BottomNavigation,
  BottomNavigationAction, makeStyles,
} from '@material-ui/core';
import { Link } from 'react-router-dom';
import BottlesIcon from '../assets/bottles.svg';
import PintOfBeerIcon from '../assets/pint-of-beer.svg';

const useStyles = makeStyles(theme => ({
  bottomNavBarAction: {
    color: `${theme.palette.action.active} !important`,
    fontSize: '0.5rem !important',
  },
}));

export default function BottomNavigationBar() {
  const [currentNav, setNav] = useState(0);
  const classes = useStyles();

  return (
    <Fragment>
      <BottomNavigation
        value={currentNav}
        onChange={(_, newNav) => setNav(newNav)}
        showLabels
      >
        <BottomNavigationAction
          id="new-round-button"
          component={Link}
          to="/new-round"
          label="New Round"
          icon={<img src={PintOfBeerIcon} alt="New Round" width={24} height={24} />}
          className={classes.bottomNavBarAction}
        />
        <BottomNavigationAction
          id="other-rounds-button"
          component={Link}
          to="/other-rounds"
          label="Other Rounds"
          icon={<img src={BottlesIcon} alt="Other Rounds" width={24} height={24} />}
          className={classes.bottomNavBarAction}
        />
      </BottomNavigation>
    </Fragment>
  );
}
