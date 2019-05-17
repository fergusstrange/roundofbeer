import React from 'react';
import { Snackbar } from '@material-ui/core';
import { roundContext } from '../store/Store';

export default function ContextMessage() {
  const [state, actions] = roundContext();

  return (
    <div>
      <Snackbar
        open={!!state.error}
        anchorOrigin={{
          vertical: 'bottom',
          horizontal: 'center',
        }}
        autoHideDuration={6000}
        onClose={actions.clearError}
        message={<span>{state.error}</span>}
      />
    </div>
  );
}
