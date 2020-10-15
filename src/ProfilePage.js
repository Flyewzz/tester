import React from "react";
import "./MainPage.css";
import "./ProfilePage.css";

import { Button, CircularProgress } from "@material-ui/core";
import PrimarySearchAppBar from "./PrimarySearchAppBar";
import { MuiThemeProvider } from "@material-ui/core/styles";
import { Pagination } from "@material-ui/lab";

import Theme from "./Theme";
import Alert from "@material-ui/lab/Alert";
import Cookie from "js-cookie";
import { Redirect } from "react-router-dom";
import jwt_decode from "jwt-decode";
import { DataGrid } from "@material-ui/data-grid";
import ProfileService from './ProfileService';

class MainPage extends React.Component {
  constructor(props) {
    super(props);

    this.token = Cookie.get("token");
    let user = null;
    if (this.token) {
      user = jwt_decode(this.token);
    }
    this.state = {
      items: {
        tasks: [],
        notification: null,
        user: user,
      },
    };
    this.columns = [
      { field: "id", headerName: "ID попытки", width: 150 },
      { field: "task_id", headerName: "Номер задания", width: 150 },
      { field: "status", headerName: "Статус решения", width: 150 },
      { field: "time", headerName: "Время отправки", width: 400 },
    ];
    this.profileService = new ProfileService();
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
        if (!result) {
          result = [];
        }
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
    }
  };

  componentDidMount() {
    this.profileService.getUserTasks(this.token).then(result => {
      if (!result) {
        result = [];
      }
      this.updateStatus('tasks', result);
    }).catch(error => {
      this.updateStatus('tasks', []);
      if (error === 401) {
        this.handleLogout();
      } else {
        alert(error);
      }
    });
    // this.updatePage(this.state.items.id);
  }

  updateStatus = (field, value) => {
    this.setState((prevState) => ({
      items: {
        ...prevState.items,
        [field]: value,
      },
    }));
  };

  handleChange = (newValue, e) => {
  };

  onSubmitClick = (e) => {
    e.preventDefault();
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
        <MuiThemeProvider theme={Theme}>
          <PrimarySearchAppBar
            avatar_text={this.state.items.user.name}
            onLogout={this.handleLogout}
          />
          </MuiThemeProvider>
          {this.state.items.notification && (
            <Alert severity={this.state.items.notification.type}>
              {this.state.items.notification.text}
            </Alert>
          )}
          <div style={{ height: 400, width: "100%" }}>
            <DataGrid
              className="task-progress-table"
              rows={this.state.items.tasks}
              columns={this.columns}
              autoPageSize
              pagination
            />
          </div>
      </>
    );
  }
}

export default MainPage;
