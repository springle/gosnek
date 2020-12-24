import React, { CSSProperties } from 'react';
import { ServerGameState } from '../redux/types';

export const Board = (state: ServerGameState) => {

    var occupied = new Map<string, CSSProperties>();

    state.board.food.forEach(p => {
        const key = JSON.stringify([p.x, p.y]);
        const value: CSSProperties = {backgroundColor: 'orange'};
        occupied.set(key, value);
    })

    // TODO change color depending on player ID
    state.players.forEach(player => {
        player.occupies.forEach(p => {
            const key = JSON.stringify([p.x, p.y])
            const value: CSSProperties = {backgroundColor: 'purple'}
            occupied.set(key, value);
        })
    })

    const gridStyle: CSSProperties = {
        gridTemplateColumns: `repeat(${state.board.width}, 1fr)`,
        gridTemplateRows: `repeat(${state.board.height}, 1fr)`
    }

    let grid: JSX.Element[] = [];

    for (let i = 0; i < state.board.height; i++) {
        for (let j = 0; j < state.board.width; j++) {
            const key = JSON.stringify([i, j]);
            const extraStyle = occupied.get(key) ? occupied.get(key) : {};
            grid.push(<div className='snake-game-cell' style={extraStyle}>({i}, {j})</div>)
        }
    }
    return (<div className='snake-game-container' style={gridStyle}>{grid}</div>)
}