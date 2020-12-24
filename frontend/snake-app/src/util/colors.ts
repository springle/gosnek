// https://blueprintjs.com/docs/#core/colors.qualitative-color-schemes
export const colors = [
    '#2965CC',
    '#29A634',
    //'#D99E0B', too similar to food
    '#D13913',
    '#8F398F',
    '#00B3A4',
    '#DB2C6F',
    '#9BBF30',
    '#96622D'
    //'#7157D9' // too similar to player
];

export const playerColor = 'purple';

/**
 * 
 * Chooses a color, given a player id.
 * 
 * For best results (to prevent color re-use), the caller should provide increasing ids
 * for each enemy player.
 */
export const getColor = (myPlayerId: number, playerId: number) => {
    if (myPlayerId === playerId) {
        return playerColor;
    }
    else {
        const idx = playerId % colors.length;
        return colors[idx];
    }
}