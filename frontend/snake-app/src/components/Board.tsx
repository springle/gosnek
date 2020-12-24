import React, { CSSProperties } from 'react';
import { ServerGameState } from '../redux/types';

/**
 * Renders a single frame of the game.
 * 
 * This is re-drawn on each message from the server.
 * 
 * @param snapshot current game state from the server.
 */
export const Board = (snapshot: ServerGameState) => {

    // The 'occupied' set represents tiles of interest on the board.
    // Under normal circumstances, the board is sparsely populated, with at most a
    // couple players and a few pieces of food.
    const occupied = new Map<string, CSSProperties>();

    // Adds food tiles to the 'occupied' set
    snapshot.board.food.forEach(p => {
        const key = JSON.stringify([p.x, p.y]);
        const value: CSSProperties = { backgroundColor: 'orange' };
        occupied.set(key, value);
    })

    // Adds player tiles to the 'occupied' set.
    // The current player is assigned a protected color.
    snapshot.players.forEach(player => {
        player.occupies.forEach(p => {
            const key = JSON.stringify([p.x, p.y])
            const value: CSSProperties = { backgroundColor: 'purple' }
            occupied.set(key, value);
        })
    })

    // We define the grid style at runtime, as the size of the board is unknown until
    // a connection to the server has been established.
    const gridStyle: CSSProperties = {
        gridTemplateColumns: `repeat(${snapshot.board.width}, 1fr)`,
        gridTemplateRows: `repeat(${snapshot.board.height}, 1fr)`
    }

    // Constructs the board, adding a 'snake-game-cell' for each point on the board.
    // When points appear in the 'occupied' set, we set additional CSS properties.
    const grid: JSX.Element[] = [];
    for (let i = 0; i < snapshot.board.height; i++) {
        for (let j = 0; j < snapshot.board.width; j++) {
            const key = JSON.stringify([i, j]);
            const extraStyle = occupied.get(key) ? occupied.get(key) : {};
            grid.push(<div key={key} className='snake-game-cell' style={extraStyle}></div>)
        }
    }
    return (
        <div className='snake-game-container' style={gridStyle}>{grid}</div>
    );
}