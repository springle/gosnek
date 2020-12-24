import { Callout, Icon, Intent } from "@blueprintjs/core";
import React from "react";
import { ServerGameState } from "../redux/types"
import { playerColor } from "../util/colors";

/**
 * Renders a 'playing' or 'not playing' message at the top of the left sidebar.
 * 
 * @param props latest game snapshot
 */
const AliveBadge = (props: {snapshot: ServerGameState}) => {
    var isAlive = false;
    for (let p of props.snapshot.players) {
        if (p.playerId === props.snapshot.playerId) {
            isAlive = true;
            break;
        }
    }

    const intent = isAlive ? Intent.SUCCESS : Intent.DANGER;
    const text = isAlive ? 'Playing' : 'Not Playing';
    const icon = <Icon icon='record' color={playerColor} iconSize={Icon.SIZE_LARGE}/>

    return (
        <Callout intent={intent} icon={icon}>
            <h4 className='bp3-heading'>{text}</h4>
        </Callout>
    )
}

export default AliveBadge;