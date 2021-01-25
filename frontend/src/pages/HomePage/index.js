import React from "react";

import { connect } from "react-redux";
import { Redirect } from "react-router-dom";

@connect((store) => ({ authStore: store.Auth }))
class HomePage extends React.Component {
  render() {
    const { isLoggedIn } = this.props.authStore;

    if (!isLoggedIn) {
      return <Redirect to="/login" />;
    }

    return <div>123</div>;
  }
}

export default HomePage;
