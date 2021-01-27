import React from "react";

import { connect } from "react-redux";
import { Link, Redirect } from "react-router-dom";

import { fetchToken } from "../../store/actions/auth";

import { Button, Col, Divider, Form, Input, Row, Space } from "antd";
import { GithubOutlined, GoogleOutlined } from "@ant-design/icons";
import "./indes.css";

@connect((store) => ({
  user: store.Auth.user,
  isLoggedIn: store.Auth.isLoggedIn,
}))
class LoginPage extends React.Component {
  onFinish = (values) => {
    const { username, password } = values;
    this.props.dispatch(fetchToken(username, password));
  };

  render() {
    if (this.props.isLoggedIn) {
      return <Redirect to="/home" />;
    }

    return (
      <Row
        type="flex"
        justify="center"
        align="middle"
        style={{ minHeight: "100vh" }}
      >
        <Col span={8}>
          <div className="login_form__header">
            <div className="site-logo" />
          </div>
          <div className="login_form">
            <div className="login_form__social">
              <div className="sign_in_text">Sign In</div>
              <Space
                style={{
                  justifyContent: "center",
                  width: "100%",
                  margin: "15px 0",
                }}
              >
                <Button icon={<GithubOutlined />}>GITHUB</Button>
                <Button icon={<GoogleOutlined />}>GOOGLE</Button>
              </Space>
              <Divider>or</Divider>
            </div>

            <Form name="basic" layout="vertical" onFinish={this.onFinish}>
              <Form.Item
                label="Username"
                name="username"
                rules={[
                  {
                    required: true,
                    message: "Please input your username!",
                  },
                ]}
              >
                <Input size="large" placeholder="Username" />
              </Form.Item>

              <Form.Item
                label="Password"
                name="password"
                rules={[
                  {
                    required: true,
                    message: "Please input your password!",
                  },
                ]}
              >
                <Input.Password size="large" placeholder="Password" />
              </Form.Item>

              <Form.Item>
                <Button
                  block
                  type="primary"
                  htmlType="submit"
                  className="login_form__submit"
                >
                  Submit
                </Button>
              </Form.Item>
            </Form>
          </div>
          <div className="login_form__footer">
            <Link to="/sign_up">Create </Link> <span>a new account</span>
          </div>
        </Col>
      </Row>
    );
  }
}

export default LoginPage;
