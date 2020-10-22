import React from "react";
import "./MainPage.css";
import Stricts from "../Stricts";
import ScriptTag from "react-script-tag";
import Results from "../components/results/Results";
import Examples from "../components/examples/Examples";
import VerdictService from "../services/VerdictService";
import TaskService from "../services/TaskService";
import Rules from "../components/rules/Rules";
import AceEditor from "react-ace";

import "ace-builds/src-noconflict/mode-c_cpp";
import "ace-builds/src-noconflict/mode-python";
import "ace-builds/src-noconflict/theme-monokai";
import "ace-builds/src-noconflict/ext-language_tools";

import { Button, CircularProgress } from "@material-ui/core";
import PrimarySearchAppBar from "../components/app_bar/PrimarySearchAppBar";
import { MuiThemeProvider } from "@material-ui/core/styles";
import { Pagination } from "@material-ui/lab";

import Theme from "../Theme";
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
        id: "1",
        // id: this.props.match.params.id,
      },
    };
    this.verdictService = new VerdictService();
    this.taskService = new TaskService();
  }

  isLogged = () => {
    return this.state.items.user;
  };

  handleLogout = () => {
    Cookie.remove("token");
    this.updateStatus("user", null);
  };

  updatePage = (id) => {
    this.taskService
      .getInfo(id, this.token)
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
      let code = localStorage.getItem(`code ${this.state.items.id}`);
      if (code) {
        this.code = code;
        editor.setValue(this.code);
      }
    }
  };

  componentDidMount() {
    this.updatePage(this.state.items.id);
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
    localStorage.setItem(`code ${this.state.items.id}`, this.code);
  };

  onSubmitClick = (e) => {
    e.preventDefault();
    if (this.code === "") {
      this.showWarning("Невозможно отправить пустой текст!");
      return;
    }
    this.updateVerdicts(null);
    this.verdictService
      .getVerdicts(this.state.items.id, this.code, this.token)
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

  changePage = (e, value) => {
    this.updateStatus("id", value);
    this.updatePage(value);
  };

  render() {
    if (!this.isLogged()) {
      return <Redirect to="/login"></Redirect>;
    }
    return (
      <>
        {/* <MuiThemeProvider theme={Theme}> */}
          <PrimarySearchAppBar
            avatar_text={this.state.items.user.name}
            onLogout={this.handleLogout}
          />
        {/* </MuiThemeProvider> */}
        <Pagination
          count={this.state.items.taskInfo.task_count}
          value={this.state.items.id}
          boundaryCount={3}
          onChange={this.changePage}
          className="task-paginator"
        />
        {/* <MuiThemeProvider theme={Theme}> */}
          <div className="task">
            <div className={"task-title"}>Задание #{this.state.items.id}</div>
            <div className={"task-text"}>
              <section
                className="not-found-controller"
                dangerouslySetInnerHTML={{
                  __html: this.state.items.taskInfo.text,
                }}
              />
            </div>
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
        {/* </MuiThemeProvider> */}
      </>
    );
  }
}

export default MainPage;
