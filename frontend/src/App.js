import React from "react";

import { connect } from "react-redux";
import { Route, Router, Switch } from "react-router-dom";
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

const { Content, Footer } = Layout;

@connect((store) => ({ authStore: store.Auth }))
class App extends React.Component {
  componentDidMount() {
    const { isLoggedIn } = this.props.authStore;
    const token = localStorage.getItem("token");

    if (token && !isLoggedIn) {
      this.props.dispatch(fetchCurrentUser());
    }
  }

  render() {
    return (
      <Router history={history}>
        <Layout style={{ minHeight: "100vh" }}>
          {this.props.authStore.user && <Sidebar logOut={this.logOut} />}
          <Layout className="app_layout">
            <Content style={{ margin: "0 16px", paddingTop: "20px" }}>
              <Switch>
                <Route exact path={["/", "/home"]} component={HomePage} />
                <Route exact path="/login" component={LoginPage} />
                <Route exact path="/sign_up" component={SignUpPage} />
                <Route exact path="/users/:id" component={ProfilePage} />
                <Route exact path="/users" component={UsersPage} />
                <Route exact path="/friends" component={FriendsPage} />
              </Switch>
            </Content>
            <Footer style={{ textAlign: "center" }}>
              MyArts ©2021 Created by Pavel Petrov
            </Footer>
          </Layout>
        </Layout>
      </Router>
    );
  }
}

export default App;
