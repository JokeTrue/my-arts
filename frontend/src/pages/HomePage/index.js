import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";

@connect((store) => ({ authStore: store.Auth }))
class HomePage extends Component {
  render() {
    const { isLoggedIn } = this.props.authStore;

    if (!isLoggedIn) {
      return <Redirect to="/login" />;
    }
  }
}

export default HomePage;
