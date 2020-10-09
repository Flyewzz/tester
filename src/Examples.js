import React from "react";
import './Examples.css';


function Examples(props) {
  const toNewLines = (string) => {
    return string.split("\n").map((item, i) => {
      return <p style={{ margin: "1px" }} key={i}>{item}</p>;
    });
  }
  const getSamples = (samples) => {
    samples = samples.split("|");
    let tags = [];
    for (let i = 0; i < samples.length; i += 2) {
      tags.push(
        <tr>
          <td>
            { toNewLines(samples[i]) }
          </td>
          <td>
          { toNewLines(samples[i + 1]) }
          </td>
        </tr>
      );
    }
    return tags;
  };

  return (
    <table className="table-examples">
      <caption style={{ marginBottom: "10px" }}>Примеры</caption>
      <thead>
        <th>Входные данные</th>
        <th>Выходные данные</th>
      </thead>
      <tbody>{props.samples && getSamples(props.samples)}</tbody>
    </table>
  );
}

export default Examples;
