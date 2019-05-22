import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import { ContextProvider } from './store/Store';

it('renders without crashing', () => {
  // eslint-disable-next-line no-undef
  const div = document.createElement('div');
  ReactDOM.render(
    // eslint-disable-next-line react/jsx-filename-extension
    <ContextProvider>
      <App />
    </ContextProvider>, div,
  );
  ReactDOM.unmountComponentAtNode(div);
});
