import { Action } from "./actions";
import { ConnectionState, InternalGameState, SET_SERVER_GAME_STATE } from "./types";

const initialState: InternalGameState = {
    connectionState: ConnectionState.INIT
};

export const reducer = (state: InternalGameState = initialState, action: Action) => {
    switch(action.type) {
        case SET_SERVER_GAME_STATE:
            // Compute the next game state
            return {...state, connectionState: ConnectionState.CONNECTED};
        default:
            return state;
    }
}
