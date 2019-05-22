import React from 'react';
import renderer from 'react-test-renderer';
import NewRoundPage from './NewRoundPage';
import { ContextProvider } from '../store/Store';

describe('New Round Page Tests', () => {
  it('renders correctly', () => expect(renderer.create(
    <ContextProvider>
      <NewRoundPage history={{}} />
    </ContextProvider>,
  ).toJSON())
    .toMatchSnapshot());
});
