import React from "react";

import { connect } from "react-redux";
import { Layout, Menu } from "antd";

import { logout } from "../../store/actions/auth";

import {
  DesktopOutlined,
  FileOutlined,
  LogoutOutlined,
  PieChartFilled,
  TeamOutlined,
  UserOutlined,
} from "@ant-design/icons";
import "./index.css";

const { Sider } = Layout;
const { SubMenu } = Menu;

@connect((store) => ({ authStore: store.Auth }))
class Sidebar extends React.Component {
  state = {
    collapsed: false,
  };

  constructor(props) {
    super(props);
  }

  onCollapse = (collapsed) => {
    this.setState({ collapsed });
  };

  render() {
    return (
      <Sider
        collapsible
        collapsed={this.state.collapsed}
        onCollapse={this.onCollapse}
      >
        <div className="logo" />
        <Menu theme="dark" defaultSelectedKeys={["1"]} mode="inline">
          <Menu.Item key="1">
            <PieChartFilled />
            <span>Option 1</span>
          </Menu.Item>
          <Menu.Item key="2">
            <DesktopOutlined />
            <span>Option 2</span>
          </Menu.Item>
          <SubMenu
            key="sub1"
            title={
              <span>
                <UserOutlined />
                <span>User</span>
              </span>
            }
          >
            <Menu.Item key="3">Tom</Menu.Item>
            <Menu.Item key="4">Bill</Menu.Item>
            <Menu.Item key="5">Alex</Menu.Item>
          </SubMenu>
          <SubMenu
            key="sub2"
            title={
              <span>
                <TeamOutlined />
                <span>Team</span>
              </span>
            }
          >
            <Menu.Item key="6">Team 1</Menu.Item>
            <Menu.Item key="8">Team 2</Menu.Item>
          </SubMenu>
          <Menu.Item key="9">
            <FileOutlined />
            <span>File</span>
          </Menu.Item>
          <Menu.Item
            key="logout"
            onClick={(e) => this.props.dispatch(logout())}
          >
            <LogoutOutlined />
            <span>Log out</span>
          </Menu.Item>
        </Menu>
      </Sider>
    );
  }
}

export default Sidebar;
