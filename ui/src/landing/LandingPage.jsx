import React from 'react';
import PropType from 'prop-types';
import { useContext } from '../store/Store';

export default function LandingPage({ history }) {
  const [state] = useContext();

  if (state.round) {
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
  history: PropType.shape().isRequired,
};
