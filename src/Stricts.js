import React from "react";

function Stricts(props) {
  return (
    <>
      <div>
        <br />
        <span>
          Ограничение по памяти: <strong>{props.mem} MB</strong>
        </span>
        <br />
        <span>
          Ограничение по дисковой памяти: <strong>{props.hdd} MB</strong>
        </span>
        <br />
        <span>
          Ограничение по времени выполнения: <strong>{props.time} ms</strong>
        </span>
      </div>
      <div>
        <span>Прочие ограничения: </span>
        <span>
          <strong>{props.limitations}</strong>
        </span>
      </div>
    </>
  );
}

export default Stricts;
