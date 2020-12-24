import { ConnectionState, ServerGameState,  SET_CONNECTION_STATE,  SET_SERVER_GAME_STATE } from "./types"

export interface SetServerGameStateAction {
    type: typeof SET_SERVER_GAME_STATE;
    payload: ServerGameState;
}

export interface SetConnectionStateAction {
    type: typeof SET_CONNECTION_STATE;
    payload: ConnectionState;
}

export const setGameState = (state: ServerGameState): GameStateAction => {
    return {
        type: SET_SERVER_GAME_STATE,
        payload: state
    };
}

export const setConnectionState = (state: ConnectionState): ConnectionStateAction => {
    return {
        type: SET_CONNECTION_STATE,
        payload: state
    };
}

export type GameStateAction = SetServerGameStateAction;
export type ConnectionStateAction = SetConnectionStateAction;