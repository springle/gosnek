import { Icon } from "@blueprintjs/core";
import React from "react";
import { Player } from "../redux/types";
import { getColor } from "../util/colors";

/**
 * Renders the player standings table on the left sidebar.
 * 
 * @param props current players, current player id
 */
const Standings = (props: { players: Player[], myPlayerId: number }) => {
    const tHead = (<thead>
        <tr>
            <th>Positition</th>
            <th>Player</th>
            <th>Size</th>
        </tr>
    </thead>
    );
    const tRows = props.players
        .sort((a, b) => b.occupies.length - a.occupies.length)
        .map((p, idx) => {
            const color = getColor(props.myPlayerId, p.playerId);
            return (<tr key={p.playerId}>
                <td>{1 + idx}</td>
                <td><Icon color={color} icon='record' />{p.name}</td>
                <td>{p.occupies.length}</td>
            </tr>);
        })
    return (
        <table className='bp3-html-table bp3-html-table-striped snake-game-standings'>
            {tHead}
            <tbody>
                {tRows}
            </tbody>
        </table>
    )
}

export default Standings;