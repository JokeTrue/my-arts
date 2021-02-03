import { chunk } from "lodash";

import { Card, Col, Empty, Layout, Row } from "antd";
import { CheckOutlined, CloseOutlined } from "@ant-design/icons";

const { Meta } = Card;

const FriendshipRequestsList = (props) => {
  const { requests, isLoading, onButtonAction } = props;
  const requestsChunks = chunk(requests, 4);

  return (
    <Layout style={{ padding: "0 24px" }}>
      {(requests.length === 0 || isLoading) && <Empty />}

      {requestsChunks.map((chunk, rowIdx) => (
        <Row key={rowIdx} style={{ marginBottom: "30px" }}>
          {chunk.map((request, requestIdx) => (
            <Col key={`col${rowIdx}_${requestIdx}`} span={6}>
              <Card
                key={requestIdx}
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
                actions={[
                  <CheckOutlined
                    onClick={() => onButtonAction("accept", request.id)}
                    key={`accept_${requestIdx}`}
                  />,
                  <CloseOutlined
                    onClick={() => onButtonAction("decline", request.id)}
                    key={`decline_${requestIdx}`}
                  />,
                ]}
              >
                <Meta
                  key={`meta_${requestIdx}`}
                  title={`${request.user.first_name} ${request.user.last_name}, ${request.user.age}`}
                  description={`${request.user.location}`}
                />
              </Card>
            </Col>
          ))}
        </Row>
      ))}
    </Layout>
  );
};

export default FriendshipRequestsList;
