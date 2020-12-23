import React from 'react';
import './App.css';
import { WelcomePage } from './components/WelcomePage';
import { RootState } from '.';
import { connect, ConnectedProps } from 'react-redux'
import { ConnectionState } from './redux/types';

const mapState = (state: RootState) => ({
  game: state.game
})

const connector = connect(mapState);

const App = (props: ConnectedProps<typeof connector>) => {
  if (props.game.connectionState === ConnectionState.INIT) {
    return <WelcomePage />;
  }
  else return null;
}


export default connect(mapState)(App);
