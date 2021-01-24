import { Layout } from "antd";
import React, { Component } from "react";
import { connect } from "react-redux";
import { Router } from "react-router-dom";

import "antd/dist/antd.css";
import "./index.css";

import Sidebar from "./components/Sidebar";
import { history } from "./helpers/history";
import Routes from "./routes";
import { fetchCurrentUser, logout } from "./store/actions/auth";

const { Content, Footer, Header } = Layout;

@connect((store) => ({ authStore: store.Auth }))
class App extends Component {
  componentDidMount() {
    const { isLoggedIn } = this.props.authStore;
    const token = localStorage.getItem("token");

    if (token && !isLoggedIn) {
      this.props.dispatch(fetchCurrentUser());
    }
  }

  logOut() {
    this.props.dispatch(logout());
  }

  render() {
    return (
      <Router history={history}>
        <Layout style={{ minHeight: "100vh" }}>
          {this.props.authStore.user && <Sidebar />}
          <Layout className="app_layout">
            {this.props.authStore.user && (
              <Header style={{ background: "#fff", padding: 0 }} />
            )}
            <Content style={{ margin: "0 16px" }}>
              <Routes />
            </Content>
            <Footer style={{ textAlign: "center" }}>
              Ant Design ©2018 Created by Ant UED
            </Footer>
          </Layout>
        </Layout>
      </Router>
    );
  }
}

export default App;
