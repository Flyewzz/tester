import React from "react";
import {Route} from 'react-router';
import {withRouter, Switch} from "react-router-dom";

import MainPage from './MainPage';

function App() {
  return (
    <div>
      <Switch>
        <Route path="/:id" component={MainPage} />
      </Switch>
    </div>
  );
}

export default withRouter(App);
