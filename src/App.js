import React from "react";
import "./App.css";
import Stricts from "./Stricts";
import ScriptTag from "react-script-tag";
import Results from "./Results";
import Examples from "./Examples";
import CodeContainer from "./CodeContainer";
import VerdictService from "./VerdictService";
import { Helmet } from "react-helmet";
import Rules from "./Rules";

class App extends React.Component {
  constructor(props) {
    super(props);
    this.code = "";
    this.state = {
      verdicts: [],
    };
    this.verdictService = new VerdictService();
  }

  updateVerdicts = (verdicts) => {
    this.setState({
      verdicts: verdicts,
    });
  };

  handleChange = (e) => {
    this.code = e.getValue();
  };

  onSubmitClick = (e) => {
    e.preventDefault();
    this.updateVerdicts(null);
    this.verdictService
      .getVerdicts(this.code)
      .then((result) => {
        this.updateVerdicts(result);
      })
      .catch((error) => {
        this.updateVerdicts([]);
        alert(error);
      });
  };

  render() {
    return (
      <>
        <Helmet>
          <link rel="stylesheet" href="/theme/dracula.css" />
          <link rel="stylesheet" href="/theme/monokai.css" />
        </Helmet>
        <ScriptTag
          isHydrating={true}
          type="text/javascript"
          src="/mode/javascript/javascript.js"
        />
        <ScriptTag
          isHydrating={true}
          type="text/javascript"
          src="/mode/clike/clike.js"
        />
        <ScriptTag
          isHydrating={true}
          type="text/javascript"
          src="/mode/python/python.js"
        />
        <div id="task">
          <blockquote>
            <span style={{ fontStyle: "italic" }}>
              <b>Задание</b><br/>
                Пользователь вводит два числа (a и b). Вычислите выражение: a<sup><small>3</small></sup>
                - b<sup><small>2</small> </sup>.
            </span>
          </blockquote>
          <Stricts mem="10 MB" hdd="1 MB" time="50ms" />
          <br />
          <br />
          <Examples />
          <br />
          <br />
          <Rules />
          
          
          
        </div>
        {/* <div id="title">
          Пользователь вводит два числа (a и b). Вычислите выражение: a^3 - b^2.
        </div> */}
        <br />
        <br />
        <CodeContainer
          onChange={this.handleChange}
          onClick={this.onSubmitClick}
          code={this.code}
          lang="python"
        />
        <Results verdicts={this.state.verdicts} />
      </>
    );
  }
}

export default App;
