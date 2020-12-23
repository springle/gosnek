import { ServerGameState, SET_SERVER_GAME_STATE } from "./types"

export interface SetServerGameStateAction {
    type: typeof SET_SERVER_GAME_STATE;
    payload: ServerGameState;
}

export type Action = SetServerGameStateAction

export const setGameState = (state: ServerGameState): Action => {
    return {
        type: SET_SERVER_GAME_STATE,
        payload: state
    };
}