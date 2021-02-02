import React, { useState } from "react";

import { Link } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";

import { logout } from "../../store/actions/auth";

import { Layout, Menu } from "antd";
import {
  HomeOutlined,
  LogoutOutlined,
  RadarChartOutlined,
  SearchOutlined,
  UserAddOutlined,
  UserOutlined,
} from "@ant-design/icons";
import "./index.css";

const { Sider } = Layout;

export default function Sidebar() {
  const dispatch = useDispatch();
  const [collapsed, setCollapsed] = useState(false);
  const { user } = useSelector((state) => state.Auth);

  if (!user) {
    return <div />;
  }

  return (
    <Sider collapsible collapsed={collapsed} onCollapse={setCollapsed}>
      <div className="logo" />
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
          <Link to={`/users/${user.id}`}>
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
        <Menu.Item key="friendship_requests">
          <Link to="/friendship_requests">
            <UserAddOutlined />
            <span>Requests</span>
          </Link>
        </Menu.Item>
        <Menu.Item key="logout" onClick={(e) => dispatch(logout())}>
          <LogoutOutlined />
          <span>Log out</span>
        </Menu.Item>
      </Menu>
    </Sider>
  );
}
