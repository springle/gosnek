import { Divider } from '@blueprintjs/core';
import React from 'react';
import { Player, ServerGameState } from '../redux/types';

const Standings = (props: {players: Player[]}) => {
    console.log(props)
    const elems = props.players
        .sort((a, b) => a.occupies.length - b.occupies.length)
        .map((p, idx) => {
            return <li key={idx}>{p.name}, size={p.occupies.length}</li>
        })
    return (
        <ol>{elems}</ol>
    )
}

const LeftSideBar = (snapshot: ServerGameState) => {
    return (
        <div>
            <h1 className='bp3-heading'>Players</h1>
            <Divider />
            <Standings players={snapshot.players}/>
        </div>
    );
}

export default LeftSideBar;