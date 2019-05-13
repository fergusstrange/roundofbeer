/* eslint-disable */
import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'
import { ContextProvider } from './store/Store'

it('renders without crashing', () => {
  const div = document.createElement('div')
  ReactDOM.render(
    <ContextProvider>
      <App/>
    </ContextProvider>, div)
  ReactDOM.unmountComponentAtNode(div)
})
