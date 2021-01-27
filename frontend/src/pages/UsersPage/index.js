import React from "react";

import { chunk } from "lodash";
import debounce from "lodash/debounce";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";

import { fetchUsers } from "../../store/actions/users";
import HistoryBreadcrumbs from "../../components/Breadcrumbs";

import { Card, Col, Input, Row } from "antd";
import Layout from "antd/es/layout/layout";
import {
  AudioOutlined,
  EyeOutlined,
  MessageOutlined,
  UserAddOutlined,
} from "@ant-design/icons";

const { Search } = Input;
const { Meta } = Card;

@connect((store) => ({
  user: store.Auth.user,
  isLoggedIn: store.Auth.isLoggedIn,
  usersStore: store.Users,
}))
class UsersPage extends React.Component {
  constructor(props) {
    super(props);

    this.state = { query: "" };
    this.debouncedSearch = debounce(this.search, 500);
  }

  componentDidMount() {
    this.search();
  }

  search = () => {
    this.props.dispatch(fetchUsers(this.state.query));
  };

  onSearch = (value) => {
    this.setState({ query: value }, () => this.debouncedSearch());
  };

  onActionProfileClick = (userId) => {
    this.props.history.push("/users/" + userId);
  };

  onActionFriendClick = (userId) => {
    console.log(userId);
  };

  onActionMessageClick = (userId) => {
    console.log(userId);
  };

  render() {
    const routes = [
      {
        path: "home",
        breadcrumbName: "Home",
      },
      {
        path: "users",
        breadcrumbName: "Users",
      },
    ];

    const { user, isLoggedIn } = this.props;
    if (user && !isLoggedIn) {
      return <Redirect to="/login" />;
    }

    const { isLoading, users } = this.props.usersStore;
    const usersChunks = chunk(users, 4);

    return (
      <>
        <HistoryBreadcrumbs
          routes={routes}
          title="Users"
          subTitle="Search for Creators from all the world"
        />

        <Layout style={{ height: "100%" }}>
          <Search
            placeholder="Type First or Last Name"
            enterButton="Search"
            size="large"
            loading={isLoading}
            style={{ padding: "0 24px", marginTop: "30px" }}
            suffix={
              <AudioOutlined
                style={{
                  fontSize: 16,
                  color: "#1890ff",
                }}
              />
            }
            onSearch={this.onSearch}
          />
          {!isLoading && (
            <Layout style={{ padding: "0 24px", marginTop: "30px" }}>
              {usersChunks.map((chunk, rowIdx) => (
                <Row key={rowIdx} style={{ marginBottom: "30px" }}>
                  {chunk.map((user, userIdx) => (
                    <Col span={6}>
                      <Card
                        key={userIdx}
                        hoverable
                        style={{ width: 240 }}
                        cover={
                          <img
                            alt="avatar"
                            src="https://sun9-36.userapi.com/sun9-55/impf/c857520/v857520744/85ad6/CLDVzPGJEkQ.jpg?size=640x640&quality=96&proxy=1&sign=73231812f06bad04f453584f93a57c29&type=album"
                          />
                        }
                        actions={[
                          <EyeOutlined
                            onClick={() => this.onActionProfileClick(user.id)}
                            key="profile"
                          />,
                          <UserAddOutlined
                            onClick={() => this.onActionFriendClick(user.id)}
                            key="friend"
                          />,
                          <MessageOutlined
                            onClick={() => this.onActionMessageClick(user.id)}
                            key="message"
                          />,
                        ]}
                      >
                        <Meta
                          key={`meta_${userIdx}`}
                          title={`${user.first_name} ${user.last_name}, ${user.age}`}
                          description={`${user.location}`}
                        />
                      </Card>
                    </Col>
                  ))}
                </Row>
              ))}
            </Layout>
          )}
        </Layout>
      </>
    );
  }
}

export default UsersPage;
