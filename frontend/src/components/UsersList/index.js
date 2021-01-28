import { chunk } from "lodash";

import Layout from "antd/es/layout/layout";
import { Card, Col, Empty, Row } from "antd";
import {
  EyeOutlined,
  MessageOutlined,
  UserAddOutlined,
} from "@ant-design/icons";
import { createFriendshipRequest } from "../../store/actions/friendshipRequests";

const { Meta } = Card;

export default function UsersList(props) {
  const {
    users,
    history,
    isLoading,
    friendsIds,
    currentUserId,
    dispatch,
  } = props;

  const onActionProfileClick = (userId) => {
    history.push("/users/" + userId);
  };

  const onActionFriendClick = (userId) => {
    dispatch(createFriendshipRequest(userId));
  };

  const onActionMessageClick = (userId) => {
    console.log(userId);
  };

  const getActions = (userId, friendIds) => {
    let actions = [
      <EyeOutlined
        onClick={() => onActionProfileClick(userId)}
        key="profile"
      />,
      <UserAddOutlined
        onClick={() => onActionFriendClick(userId)}
        key="friend"
      />,
      <MessageOutlined
        onClick={() => onActionMessageClick(userId)}
        key="message"
      />,
    ];
    if (friendIds.includes(userId)) {
      actions.splice(1, 1);
    }
    return actions;
  };

  const usersList = users.filter((item) => item.id !== currentUserId);
  const usersChunks = chunk(usersList, 4);

  return (
    <Layout style={{ padding: "0 24px", marginTop: "30px" }}>
      {(users.length === 0 || isLoading) && <Empty />}

      {users.length > 0 &&
        usersChunks.map((chunk, rowIdx) => (
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
                  actions={getActions(user.id, friendsIds)}
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
  );
}