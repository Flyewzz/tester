import React from 'react';


function Stricts(props) {
    return (
        <div>
        <br />
        <span>Ограничение по памяти: <strong>{props.mem}</strong></span
        ><br />
        <span>Ограничение по дисковой памяти: <strong>{props.hdd}</strong></span
        ><br />
        <span>Ограничение по времени выполнения: <strong>{props.time}</strong></span
        >
        </div>
    )
}

export default Stricts;