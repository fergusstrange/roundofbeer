import React from 'react';
import renderer from 'react-test-renderer';
import LandingPage from './LandingPage';
import { ContextProvider } from '../store/Store';

describe('Landing Page Tests', () => {
  it('Renders existing round join page', () => expect(renderer.create(
    <ContextProvider>
      <LandingPage
        history={{
          push: () => {},
        }}
      />
    </ContextProvider>,
  ).toJSON())
    .toMatchSnapshot());
});
