export const SET_SERVER_GAME_STATE = 'SET_SERVER_GAME_STATE';
export const SET_CONNECTION_STATE = 'SET_CONNECTION_STATE';

export enum ConnectionState {
    INIT,
    LOADING,
    CONNECTED
}

export interface InternalGameState {
    snapshot?: ServerGameState;
}

export interface ServerGameState {
    playerId: number;
    board: Board;
    players: Player[];
}

export interface Board {
    width: number;
    height: number;
    food: Point[];
};

export interface Player {
    playerId: number;
    name: string;
    occupies: Point[];
}

export interface Point {
    x: number;
    y: number;
}
