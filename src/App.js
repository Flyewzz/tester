import React from "react";
import {Redirect, Route} from 'react-router';
import {withRouter, Switch} from "react-router-dom";

import MainPage from './MainPage';
import LoginPage from './LoginPage';
import SignUpPage from './SignUpPage';
import ProfilePage from './ProfilePage';

function App() {
  return (
    <div>
      <Switch>
        <Route path="/login" component={LoginPage} />
        <Route path="/signup" component={SignUpPage} />
        <Route path="/profile" component={ProfilePage} />
        <Route exact path="/:id" component={MainPage} />
        <Route exact path="/" component={() => <Redirect to="/1" />} />
      </Switch>
    </div>
  );
}

export default withRouter(App);
