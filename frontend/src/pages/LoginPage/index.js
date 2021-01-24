import { GithubOutlined, GoogleOutlined } from "@ant-design/icons";
import { Button, Col, Form, Input, Row, Space } from "antd";
import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";

import "./indes.css";
import { fetchToken } from "../../store/actions/auth";

@connect((store) => ({ authStore: store.Auth }))
class LoginPage extends Component {
  onFinish = (values) => {
    const { username, password } = values;
    this.props.dispatch(fetchToken(username, password));
  };

  render() {
    if (this.props.authStore.isLoggedIn) {
      return <Redirect to="/home" />;
    }

    return (
      <Row
        type="flex"
        justify="center"
        align="middle"
        style={{ minHeight: "100vh" }}
      >
        <Col span={8} className="login_form">
          <div className="login_form_social">
            <div className="sign_in_text">Sign in with</div>
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
            <hr style={{ margin: "35px 0" }} />
            <div className="sign_in_text">Or sign in with credentials</div>
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
                className="login_form_submit"
              >
                Submit
              </Button>
            </Form.Item>
          </Form>
        </Col>
      </Row>
    );
  }
}

export default LoginPage;
