import React from "react";
import "./MainPage.css";
import Stricts from "./Stricts";
import ScriptTag from "react-script-tag";
import Results from "./Results";
import Examples from "./Examples";
import CodeContainer from "./CodeContainer";
import VerdictService from "./VerdictService";
import TaskService from "./TaskService";
import { Helmet } from "react-helmet";
import Rules from "./Rules";

class MainPage extends React.Component {
  constructor(props) {
    super(props);
    this.code = "";
    this.state = {
      items: {
        verdicts: [],
        taskInfo: {},
      },
    };
    this.verdictService = new VerdictService();
    this.taskService = new TaskService();
    this.state.items.taskInfo = {};
    this.id = null;
  }

  componentDidMount() {
    this.id = this.props.match.params.id;
    this.taskService
      .getInfo(this.id)
      .then((result) => {
        this.updateInfo(result);
      })
      .catch((error) => {
        this.updateInfo({});
        alert(error);
      });
  }

  updateInfo = (info) => {
    this.setState((prevState) => ({
      items: {
        ...prevState.items,
        taskInfo: info,
      },
    }));
  };

  updateVerdicts = (verdicts) => {
    this.setState((prevState) => ({
      items: {
        ...prevState.items,
        verdicts: verdicts,
      },
    }));
  };

  handleChange = (e) => {
    this.code = e.getValue();
  };

  onSubmitClick = (e) => {
    e.preventDefault();
    this.updateVerdicts(null);
    this.verdictService
      .getVerdicts(this.id, this.code)
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
              <b>Задание</b>
              <br />
              {this.state.items.taskInfo.text}
            </span>
          </blockquote>
          <Stricts
            mem={this.state.items.taskInfo.ram}
            hdd={this.state.items.taskInfo.hdd}
            time={this.state.items.taskInfo.time}
            limitations={this.state.items.taskInfo.limitations}
          />
          <br />
          <br />
          <Examples samples={this.state.items.taskInfo.samples} />
          <br />
          <br />
          <Rules />
        </div>
        <br />
        <br />
        <CodeContainer
          onChange={this.handleChange}
          onClick={this.onSubmitClick}
          code={this.code}
          lang="c++"
        />
        <Results verdicts={this.state.items.verdicts} />
      </>
    );
  }
}

export default MainPage;
