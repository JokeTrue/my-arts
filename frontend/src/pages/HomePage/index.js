import React from "react";

import { connect } from "react-redux";
import { Redirect } from "react-router-dom";

@connect((store) => ({
  user: store.Auth.user,
  isLoggedIn: store.Auth.isLoggedIn,
}))
class HomePage extends React.Component {
  render() {
    const { user, isLoggedIn } = this.props;
    if (user && !isLoggedIn) {
      return <Redirect to="/login" />;
    }

    return <div>123</div>;
  }
}

export default HomePage;
