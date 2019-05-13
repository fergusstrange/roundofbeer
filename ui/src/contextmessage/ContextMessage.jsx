import React from 'react';
import { Snackbar } from '@material-ui/core';
import { useContext } from '../store/Store';

export default function ContextMessage() {
  const [state, actions] = useContext();

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
