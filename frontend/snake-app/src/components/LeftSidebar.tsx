import { Callout } from '@blueprintjs/core';
import React from 'react';
import { ServerGameState } from '../redux/types';
import AliveBadge from './AliveBadge';
import Standings from './Standings';


const LeftSideBar = (snapshot: ServerGameState) => {
    return (
        <Callout className='left-sidebar'>
            <AliveBadge snapshot={snapshot} />
            <Standings players={snapshot.players} myPlayerId={snapshot.playerId} />
        </Callout>
    );
}

export default LeftSideBar;