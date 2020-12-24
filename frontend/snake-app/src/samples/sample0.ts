import { ServerGameState } from "../redux/types"

export const sample0 = {
    board: {
        width: 25,
        height: 25,
        food: [
            { x: 12, y: 12 },
            { x: 100, y: 80 }
        ]
    },
    players: [
        { name: 'OBAMA', occupies: [{ x: 1, y: 1 }, { x: 1, y: 2 }, { x: 1, y: 3 }] },
        { name: 'GWB', occupies: [{ x: 10, y: 1 }, { x: 10, y: 2 }, { x: 10, y: 3 }] }
    ]
} as ServerGameState;