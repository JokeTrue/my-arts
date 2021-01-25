import React from "react";

import { Route, Switch } from "react-router-dom";

import HomePage from "../pages/HomePage";
import LoginPage from "../pages/LoginPage";
import SignUpPage from "../pages/SignUpPage";

const Routes = () => (
  <Switch>
    <Route exact path={["/", "/home"]} component={HomePage} />
    <Route exact path="/login" component={LoginPage} />
    <Route exact path="/sign_up" component={SignUpPage} />
  </Switch>
);

export default Routes;
