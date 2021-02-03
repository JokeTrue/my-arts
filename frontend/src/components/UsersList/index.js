import { chunk } from "lodash";

import Layout from "antd/es/layout/layout";
import { Button, Card, Col, Empty, Row } from "antd";
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
    hasMore,
    loadMore,
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
                  cover={<img alt="avatar" src="https://picsum.photos/200" />}
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
      {!isLoading && hasMore && (
        <Row style={{ justifyContent: "center" }}>
          <Button
            type="primary"
            block
            ghost
            style={{ width: "50%" }}
            onClick={(_) => loadMore()}
          >
            Load More
          </Button>
        </Row>
      )}
    </Layout>
  );
}