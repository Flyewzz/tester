import React from "react";
import {Route} from 'react-router';
import {withRouter, Switch} from "react-router-dom";

import MainPage from './MainPage';
import LoginPage from './LoginPage';
import SignUpPage from './SignUpPage';

function App() {
  return (
    <div>
      <Switch>
        <Route path="/login" component={LoginPage} />
        <Route path="/signup" component={SignUpPage} />
        <Route path="/:id" component={MainPage} />
      </Switch>
    </div>
  );
}

export default withRouter(App);
