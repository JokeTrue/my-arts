import React, { useEffect } from "react";

import { useDispatch, useSelector } from "react-redux";
import { Redirect, Route, Router, Switch } from "react-router-dom";
import { history } from "./helpers/history";
import { fetchCurrentUser } from "./store/actions/auth";

import HomePage from "./pages/HomePage";
import LoginPage from "./pages/LoginPage";
import SignUpPage from "./pages/SignUpPage";
import ProfilePage from "./pages/ProfilePage";
import UsersPage from "./pages/UsersPage";
import FriendsPage from "./pages/FriendsPage";

import { Layout } from "antd";
import Sidebar from "./components/Sidebar";
import FriendshipRequestsPage from "./pages/FriendshipRequestsPage";
import { PrivateRoute } from "./components/PrivateRoute";
import { CurrentUserContext } from "./helpers/currentUserContext";

const { Content, Footer } = Layout;

export default function App() {
  const dispatch = useDispatch();
  const { user } = useSelector((state) => state.Auth);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token !== null && token !== "" && user === null) {
      dispatch(fetchCurrentUser());
    }
  });

  return (
    <CurrentUserContext.Consumer>
      {(ctx) => (
        <CurrentUserContext.Provider value={{ user }}>
          <Router history={history}>
            <Layout style={{ minHeight: "100vh" }}>
              <Sidebar />
              <Layout className="app_layout">
                <Content style={{ margin: "0 16px", paddingTop: "20px" }}>
                  <Switch>
                    <Route exact path="/login" component={LoginPage} />
                    <Route exact path="/sign_up" component={SignUpPage} />
                    <PrivateRoute exact path="/home" component={HomePage} />
                    <PrivateRoute
                      exact
                      path="/users/:id"
                      component={ProfilePage}
                    />
                    <PrivateRoute exact path="/users" component={UsersPage} />
                    <PrivateRoute
                      exact
                      path="/friends"
                      component={FriendsPage}
                    />
                    <PrivateRoute
                      exact
                      path="/friendship_requests"
                      component={FriendshipRequestsPage}
                    />
                    <Redirect to="/login" />
                  </Switch>
                </Content>
                <Footer style={{ textAlign: "center" }}>
                  MyArts ©2021 Created by Pavel Petrov
                </Footer>
              </Layout>
            </Layout>
          </Router>
        </CurrentUserContext.Provider>
      )}
    </CurrentUserContext.Consumer>
  );
}
