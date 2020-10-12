import React from "react";
import "./MainPage.css";
import Stricts from "./Stricts";
import ScriptTag from "react-script-tag";
import Results from "./Results";
import Examples from "./Examples";
import VerdictService from "./VerdictService";
import TaskService from "./TaskService";
import Rules from "./Rules";
import AceEditor from "react-ace";

import "ace-builds/src-noconflict/mode-c_cpp";
import "ace-builds/src-noconflict/mode-python";
import "ace-builds/src-noconflict/theme-monokai";
import "ace-builds/src-noconflict/ext-language_tools";

import { Button, CircularProgress } from "@material-ui/core";
import PrimarySearchAppBar from "./PrimarySearchAppBar";
import { MuiThemeProvider } from "@material-ui/core/styles";

import Theme from "./Theme";
import Alert from "@material-ui/lab/Alert";
import Cookie from "js-cookie";
import { Redirect } from "react-router-dom";
import jwt_decode from "jwt-decode";

class MainPage extends React.Component {
  constructor(props) {
    super(props);
    this.code = "";

    this.token = Cookie.get("token");
    let user = null;
    if (this.token) {
      user = jwt_decode(this.token);
    }
    this.state = {
      items: {
        verdicts: [],
        taskInfo: {},
        notification: null,
        user: user,
      },
    };
    this.verdictService = new VerdictService();
    this.taskService = new TaskService();
    this.id = null;
  }

  isLogged = () => {
    return this.state.items.user;
  };

  handleLogout = () => {
    Cookie.remove("token");
    this.updateStatus("user", null);
  };

  componentDidMount() {
    this.id = this.props.match.params.id;

    this.taskService
      .getInfo(this.id, this.token)
      .then((result) => {
        this.updateInfo(result);
      })
      .catch((error) => {
        this.updateInfo({});
        if (error === 401) {
          this.handleLogout();
        } else {
          alert(error);
        }
      });
    if (this.isLogged()) {
      const editor = this.ace.editor;
      editor.setFontSize(16);
      editor.setOptions({
        enableBasicAutocompletion: true,
        enableLiveAutocompletion: true,
        enableSnippets: true,
      });
    }
  }

  updateInfo = (info) => {
    this.setState((prevState) => ({
      items: {
        ...prevState.items,
        taskInfo: info,
      },
    }));
  };

  updateNotification = (notification) => {
    this.setState((prevState) => ({
      items: {
        ...prevState.items,
        notification: notification,
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

  updateStatus = (field, value) => {
    this.setState((prevState) => ({
      items: {
        ...prevState.items,
        [field]: value,
      },
    }));
  };

  handleChange = (newValue, e) => {
    this.updateNotification(null);
    const editor = this.ace.editor; // The editor object is from Ace's API
    this.code = editor.getValue();
  };

  onSubmitClick = (e) => {
    e.preventDefault();
    if (this.code === "") {
      this.showWarning("Невозможно отправить пустой текст!");
      return;
    }
    this.updateVerdicts(null);
    this.verdictService
      .getVerdicts(this.id, this.code, this.token)
      .then((result) => {
        this.updateVerdicts(result);
      })
      .catch((error) => {
        this.updateVerdicts([]);
        if (error === 401) {
          this.handleLogout();
        } else {
          alert(error);
        }
      });
  };

  showWarning = (text) => {
    this.updateNotification({
      type: "warning",
      text: text,
    });
  };

  render() {
    if (!this.isLogged()) {
      return <Redirect to="/login"></Redirect>;
    }
    return (
      <>
        <MuiThemeProvider theme={Theme}>
          <PrimarySearchAppBar
            avatar_text={this.state.items.user.name}
            onLogout={this.handleLogout}
          />
          <div className="task">
            <div className={"task-title"}>Задание #{this.id}</div>
            <div className={"task-text"}>{this.state.items.taskInfo.text}</div>
            <div className={"task-info"}>
              <Stricts
                mem={this.state.items.taskInfo.ram}
                hdd={this.state.items.taskInfo.hdd}
                time={this.state.items.taskInfo.time}
                limitations={this.state.items.taskInfo.limitations}
              />
            </div>
            <Examples samples={this.state.items.taskInfo.samples} />
            <br />
            <br />
            <Rules />
          </div>
          <br />
          <br />
          {this.state.items.notification && (
            <Alert severity={this.state.items.notification.type}>
              {this.state.items.notification.text}
            </Alert>
          )}
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
                }}
              />
              <form encType="multipart/form-data">
                <Button
                  variant="contained"
                  color="primary"
                  style={{
                    margin: "10px 0px 20px",
                    backgroundColor: "#398A96",
                  }}
                  onClick={this.onSubmitClick}
                >
                  Submit
                </Button>
              </form>
            </div>
            <div>
              <Results verdicts={this.state.items.verdicts} />
            </div>
          </div>
        </MuiThemeProvider>
      </>
    );
  }
}

export default MainPage;
