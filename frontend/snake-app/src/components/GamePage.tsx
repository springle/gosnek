import React from 'react';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '..';
import { sample0 } from '../samples/sample0';
import { Board } from './Board';
import LeftSidebar from './LeftSidebar';

const mapState = (root: RootState) => ({
    gameState: root.gameState
})

const connector = connect(mapState);

const GamePage = (props: ConnectedProps<typeof connector>) => {
    //const board = props.gameState.snapshot? Board(props.gameState.snapshot) : null;
    const board = Board(sample0);
    const sidebar = LeftSidebar(sample0);
    return (
        <div className='game-page-container'>
            {sidebar}
            {board}
        </div>
    )
}

export default connect(mapState)(GamePage);
