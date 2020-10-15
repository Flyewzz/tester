import React from "react";
import "./Results.css";
import { CircularProgress } from "@material-ui/core";
import Theme from './Theme';
import { MuiThemeProvider } from "@material-ui/core/styles";


class Results extends React.Component {
  getVerdicts = (verdicts) => {
    return verdicts.map((verdict) => {
      return (
        <tr>
          <td>{verdict.name}</td>
          <td>
            <b
              className={
                verdict.status === "OK" ? "positive-answer" : "negative-answer"
              }
            >
              {verdict.status}
            </b>
          </td>
          <td style={{ padding: "10px" }}>{verdict.message}</td>
        </tr>
      );
    });
  };

  render() {
    return this.props.verdicts ? (
      <table className="table-results">
        <thead>
          <th>Название</th>
          <th>Статус</th>
          <th>Сообщение</th>
        </thead>
        <tbody>{this.getVerdicts(this.props.verdicts)}</tbody>
      </table>
    ) : (
      <MuiThemeProvider theme={Theme}>
        <CircularProgress color="secondary" />
      </MuiThemeProvider>
    );
  }
}

export default Results;
