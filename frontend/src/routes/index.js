import React from "react";
import { Route, Switch } from "react-router-dom";

import HomePage from "../pages/HomePage";
import LoginPage from "../pages/LoginPage";

const Routes = () => (
  <Switch>
    <Route exact path={["/", "/home"]} component={HomePage} />
    <Route exact path="/login" component={LoginPage} />
  </Switch>
);

export default Routes;
