import React from "react";

import { Link } from "react-router-dom";
import { connect } from "react-redux";

import { logout } from "../../store/actions/auth";

import { Layout, Menu } from "antd";
import {
  HomeOutlined,
  LogoutOutlined,
  RadarChartOutlined,
  SearchOutlined,
  UserOutlined,
} from "@ant-design/icons";
import "./index.css";

const { Sider } = Layout;

@connect((store) => ({ authStore: store.Auth }))
class Sidebar extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      collapsed: false,
    };
  }

  onCollapse = (collapsed) => {
    this.setState({ collapsed });
  };

  render() {
    const { id } = this.props.authStore.user;
    return (
      <Sider
        collapsible
        collapsed={this.state.collapsed}
        onCollapse={this.onCollapse}
      >
        <div className="logo" />
        {id && (
          <Menu theme="dark" mode="inline">
            <Menu.Item key="home">
              <Link to="/home">
                <HomeOutlined />
                <span>Home</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="users">
              <Link to="/users">
                <SearchOutlined />
                <span>Users</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="profile">
              <Link to={`/users/${id}`}>
                <UserOutlined />
                <span>Profile</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="friends">
              <Link to="/friends">
                <RadarChartOutlined />
                <span>Friends</span>
              </Link>
            </Menu.Item>
            <Menu.Item
              key="logout"
              onClick={(e) => this.props.dispatch(logout())}
            >
              <LogoutOutlined />
              <span>Log out</span>
            </Menu.Item>
          </Menu>
        )}
      </Sider>
    );
  }
}

export default Sidebar;
