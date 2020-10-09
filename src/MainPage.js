import React from "react";
import "./MainPage.css";
import Stricts from "./Stricts";
import ScriptTag from "react-script-tag";
import Results from "./Results";
import Examples from "./Examples";
import VerdictService from "./VerdictService";
import TaskService from "./TaskService";
import { Helmet } from "react-helmet";
import Rules from "./Rules";
import AceEditor from "react-ace";

import "ace-builds/src-noconflict/mode-c_cpp";
import "ace-builds/src-noconflict/mode-python";
import "ace-builds/src-noconflict/theme-monokai";
import "ace-builds/src-noconflict/ext-language_tools";

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
    const editor = this.ace.editor;
    editor.setFontSize(16);
    editor.setOptions({
      enableBasicAutocompletion: true,
      enableLiveAutocompletion: true,
      enableSnippets: true,
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
  
  handleChange = (newValue, e) => {
    const editor = this.ace.editor; // The editor object is from Ace's API
    this.code = editor.getValue()
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
        {/* <ScriptTag
          isHydrating={true}
          type="text/javascript"
          src="/mode/clike/clike.js"
        />
        <ScriptTag
          isHydrating={true}
          type="text/javascript"
          src="/mode/python/python.js"  
        /> */}
        <div className="task">
          <blockquote>
            <div className={"task-title"}>
              Задание #{this.id}
              <br />
            </div>
            <div className={"task-text"}>{this.state.items.taskInfo.text}</div>
          </blockquote>
          <div className={"task-info"}>
            <Stricts
              mem={this.state.items.taskInfo.ram}
              hdd={this.state.items.taskInfo.hdd}
              time={this.state.items.taskInfo.time}
              limitations={this.state.items.taskInfo.limitations}
            />
          </div>
          <div className={"task-info"}>
            <Examples samples={this.state.items.taskInfo.samples} />
          </div>
          <br />
          <br />
          <Rules />
        </div>
        <br />
        <br />
        <div className="code-results">
          <div>
            <AceEditor
              mode="c_cpp"
              theme="monokai"
              onChange={this.handleChange}
              name="UNIQUE_ID_OF_DIV"
              style={{ height: "500px", width: "auto" }}
              ref={(instance) => {
                this.ace = instance;
              }} // Let's put things into scope
            />

            <form encType="multipart/form-data">
              <input
                className={"input"}
                type="submit"
                name="preview-form-submit"
                id="preview-form-submit"
                value="Submit"
                onClick={this.onSubmitClick}
              />
            </form>
          </div>
          <div>
            <Results verdicts={this.state.items.verdicts} />
          </div>
        </div>
      </>
    );
  }
}

export default MainPage;
