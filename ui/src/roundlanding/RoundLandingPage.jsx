import React from 'react';
import PropTypes from 'prop-types';
import { connect } from '../store/Store';

const RoundLandingPage = ({ match }) => (
  <div>
    {match.params.roundUrl}
  </div>
);

RoundLandingPage.propTypes = {
  match: PropTypes.shape().isRequired,
};

export default connect(state => state)(RoundLandingPage);
