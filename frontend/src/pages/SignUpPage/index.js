import React from "react";

import { Redirect } from "react-router-dom";
import { connect } from "react-redux";

import { signUp } from "../../store/actions/auth";

import {
  Button,
  Col,
  Form,
  Input,
  InputNumber,
  Row,
  Select,
  Space,
} from "antd";
import { GithubOutlined, GoogleOutlined } from "@ant-design/icons";
import "./indes.css";

const { TextArea } = Input;

@connect((store) => ({ authStore: store.Auth }))
class SignUpPage extends React.Component {
  onFinish = (values) => {
    this.props.dispatch(signUp(values));
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
        <Col span={8}>
          <div className="signup_form__header">
            <div className="site-logo" />
          </div>
          <div className="signup_form">
            <div className="signup_form__social">
              <div className="sign_in_text">Sign Up</div>
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
              <div className="line">
                <div className="or">or</div>
              </div>
            </div>
            <Form name="basic" layout="vertical" onFinish={this.onFinish}>
              <Form.Item
                name="email"
                label="E-mail"
                rules={[
                  {
                    type: "email",
                    message: "The input is not valid E-mail!",
                  },
                  {
                    required: true,
                    message: "Please input your E-mail!",
                  },
                ]}
              >
                <Input />
              </Form.Item>

              <Form.Item
                name="password1"
                label="Password"
                rules={[
                  {
                    required: true,
                    message: "Please input your password!",
                  },
                ]}
                hasFeedback
              >
                <Input.Password />
              </Form.Item>

              <Form.Item
                name="password2"
                label="Confirm Password"
                dependencies={["password1"]}
                hasFeedback
                rules={[
                  {
                    required: true,
                    message: "Please confirm your password!",
                  },
                  ({ getFieldValue }) => ({
                    validator(_, value) {
                      if (!value || getFieldValue("password1") === value) {
                        return Promise.resolve();
                      }

                      return Promise.reject(
                        "The two passwords that you entered do not match!"
                      );
                    },
                  }),
                ]}
              >
                <Input.Password />
              </Form.Item>

              <Form.Item
                name="first_name"
                label="First Name"
                rules={[
                  {
                    required: true,
                    message: "Please input your first name!",
                    whitespace: true,
                  },
                ]}
              >
                <Input />
              </Form.Item>

              <Form.Item
                name="last_name"
                label="Last name"
                rules={[
                  {
                    required: true,
                    message: "Please input your last name!",
                    whitespace: true,
                  },
                ]}
              >
                <Input />
              </Form.Item>

              <Form.Item
                name="age"
                label="Age"
                rules={[
                  {
                    type: "number",
                    min: 18,
                    max: 99,
                  },
                ]}
              >
                <InputNumber />
              </Form.Item>

              <Form.Item
                name="gender"
                label="Gender"
                rules={[
                  {
                    required: true,
                  },
                ]}
              >
                <Select placeholder="Select your gender" allowClear>
                  <Select.Option value="M">Male</Select.Option>
                  <Select.Option value="F">Female</Select.Option>
                  <Select.Option value="O">Other</Select.Option>
                </Select>
              </Form.Item>

              <Form.Item
                name="location"
                label="Location"
                rules={[
                  {
                    required: true,
                    message: "Please input your location!",
                    whitespace: true,
                  },
                ]}
              >
                <Input />
              </Form.Item>

              <Form.Item name="biography" label="Biography">
                <TextArea rows={4} />
              </Form.Item>

              <Form.Item>
                <Button
                  block
                  type="primary"
                  htmlType="submit"
                  className="signup_form_submit"
                >
                  Submit
                </Button>
              </Form.Item>
            </Form>
          </div>
        </Col>
      </Row>
    );
  }
}

export default SignUpPage;
