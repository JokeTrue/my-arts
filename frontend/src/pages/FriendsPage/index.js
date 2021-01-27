import React from "react";

import { chunk } from "lodash";
import { connect } from "react-redux";

import HistoryBreadcrumbs from "../../components/Breadcrumbs";
import { fetchFriends } from "../../store/actions/friends";

import { Card, Col, Layout, Row } from "antd";
import { EyeOutlined, MessageOutlined } from "@ant-design/icons";

const { Meta } = Card;

@connect((store) => ({
  user: store.Auth.user,
  friendsStore: store.Friends,
  isLoggedIn: store.Auth.isLoggedIn,
}))
class FriendsPage extends React.Component {
  componentDidMount() {
    debugger;
    const { user } = this.props;
    if (user) {
      this.props.dispatch(fetchFriends(user.id));
    }
  }

  onActionProfileClick = (userId) => {
    this.props.history.push("/users/" + userId);
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
        path: "friends",
        breadcrumbName: "Friends",
      },
    ];

    const { isLoading, users } = this.props.friendsStore;
    const usersChunks = chunk(users, 4);

    return (
      <>
        <HistoryBreadcrumbs
          routes={routes}
          title="Friends"
          subTitle="Your closest friends"
        />
        <Layout style={{ height: "100%" }}>
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

export default FriendsPage;
