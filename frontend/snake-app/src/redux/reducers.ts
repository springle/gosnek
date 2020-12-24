import { ConnectionStateAction, GameStateAction } from "./actions";
import { ConnectionState, InternalGameState, SET_CONNECTION_STATE, SET_SERVER_GAME_STATE } from "./types";

const initialGameState: InternalGameState = {
};

export const gameStateReducer = (state: InternalGameState = initialGameState, action: GameStateAction) => {
    switch(action.type) {
        case SET_SERVER_GAME_STATE:
            return {...state, snapshot: action.payload}
        default:
            return state;
    }
}

export const connectionStateReducer = (state: ConnectionState = ConnectionState.INIT, action: ConnectionStateAction) => {
    switch(action.type) {
        case SET_CONNECTION_STATE:
            return action.payload;
        default:
            return state;
    }
}
