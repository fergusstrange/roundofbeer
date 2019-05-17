import React from 'react';
import PropTypes from 'prop-types';
import { roundContext } from '../store/Store';

export default function LandingPage({ history }) {
  const [state] = roundContext();

  if (state.round && state.roundToken) {
    history.push(`/${state.round.url}`);
  } else {
    history.push('/new-round');
  }

  return (
    <div>
      You need to enable JavaScript to run this app.
    </div>
  );
}

LandingPage.propTypes = {
  history: PropTypes.shape().isRequired,
};
