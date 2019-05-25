import React from 'react';
import renderer from 'react-test-renderer';
import JoinRoundPage from './JoinRoundPage';
import { ContextProvider } from '../store/Store';

describe('Join Round Page Tests', () => {
  it('Renders existing round join page', () => expect(renderer.create(
    <ContextProvider>
      <JoinRoundPage
        history={{}}
        match={{
          params: {
            roundUrl: 'abcdefg',
          },
        }}
      />
    </ContextProvider>,
  ).toJSON())
    .toMatchSnapshot());
});
