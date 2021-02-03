import { chunk } from "lodash";
import { Button, Card, Col, Empty, Layout, Row, Spin, Typography } from "antd";
import {
  EyeOutlined,
  MessageOutlined,
  UserAddOutlined,
} from "@ant-design/icons";
import { createFriendshipRequest } from "../../store/actions/friendshipRequests";

const { Meta } = Card;
const { Paragraph } = Typography;
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
      {!isLoading && users.length === 0 && <Empty />}

      {isLoading && users.length === 0 && <Spin size="large" />}

      {users.length > 0 &&
        usersChunks.map((chunk, rowIdx) => (
          <Row key={rowIdx} style={{ marginBottom: "30px", justifyContent: "space-between" }}>
            {chunk.map((user, userIdx) => (
              <Col span={6}>
                <Card
                  key={userIdx}
                  hoverable
                  style={{ width: 240 }}
                  cover={
                    <img
                      alt="avatar"
                      src={`https://picsum.photos/seed/${Math.random()
                        .toString(36)
                        .substring(10)}/240/240?grayscale`}
                    />
                  }
                  actions={getActions(user.id, friendsIds)}
                >
                  <Meta
                    key={`meta_${userIdx}`}
                    title={`${user.first_name} ${user.last_name}, ${user.age}`}
                    description={
                      <Paragraph ellipsis={{ rows: 2 }} style={{minHeight: "44px"}}>
                        {user.location}
                      </Paragraph>
                    }
                  />
                </Card>
              </Col>
            ))}
          </Row>
        ))}

      {users.length > 0 && hasMore && (
        <Row style={{ justifyContent: "center" }}>
          <Button
            type="primary"
            block
            ghost
            style={{ width: "50%" }}
            onClick={(_) => loadMore()}
          >
            {isLoading ? <Spin /> : "Load More"}
          </Button>
        </Row>
      )}
    </Layout>
  );
}