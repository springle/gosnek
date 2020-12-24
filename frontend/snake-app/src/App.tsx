import React from 'react';
import './App.css';
import WelcomePage from './components/WelcomePage';
import { RootState } from '.';
import { connect, ConnectedProps } from 'react-redux'
import { ConnectionState } from './redux/types';
import GamePage from './components/GamePage';

const mapState = (root: RootState) => ({
  connectionState: root.connectionState
});

const connector = connect(mapState);

const App = (props: ConnectedProps<typeof connector>) => {
  if (props.connectionState === ConnectionState.INIT || props.connectionState === ConnectionState.LOADING) {
    return <WelcomePage />;
  }
  else return <GamePage/>;
}

export default connect(mapState)(App);
