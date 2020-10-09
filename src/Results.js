import React from "react";
import './Results.css';


class Results extends React.Component {

  getVerdicts = (verdicts) => {
    console.log(`results ${verdicts.length}`)
    return verdicts.map((verdict) => {
      return (
        <tr>
          <td>{verdict.name}</td>
          {verdict.status === "OK" ? (
            <td style={{color: 'green'}}>{verdict.status}</td>
          ) : (
            <td style={{color: 'red'}}>{verdict.status}</td>
          )}
          <td style={{padding: "10px"}}>{verdict.message}</td>
        </tr>
      );
    });
  };

  render() {
    return this.props.verdicts == null ? (
      <span>Обработка...</span>
    ) : (
      <table className='table-results'>
        <thead>
          <th>Название</th>
          <th>Статус</th>
          <th>Сообщение</th>
        </thead>
        <tbody>{this.getVerdicts(this.props.verdicts)}</tbody>
      </table>
    );
  }
}

export default Results;
