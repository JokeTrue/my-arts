import React from "react";

import { connect } from "react-redux";
import { Router } from "react-router-dom";

import Routes from "./routes";
import { history } from "./helpers/history";
import { fetchCurrentUser } from "./store/actions/auth";

import { Layout } from "antd";
import Sidebar from "./components/Sidebar";

import "antd/dist/antd.css";
import "./index.css";

const { Content, Footer, Header } = Layout;

@connect((store) => ({ authStore: store.Auth }))
class App extends React.Component {
  constructor(props) {
    super(props);
  }

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
            {this.props.authStore.user && (
              <Header style={{ background: "#fff", padding: 0 }} />
            )}
            <Content style={{ margin: "0 16px" }}>
              <Routes />
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
