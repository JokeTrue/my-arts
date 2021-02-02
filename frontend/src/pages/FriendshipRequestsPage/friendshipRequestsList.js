import { chunk } from "lodash";

import { Card, Col, Empty, Layout, Row } from "antd";
import { CheckOutlined, CloseOutlined } from "@ant-design/icons";

const { Meta } = Card;

const FriendshipRequestsList = (props) => {
  const { requests, isLoading, onButtonAction } = props;
  const requestsChunks = chunk(requests, 4);

  return (
    <Layout style={{ padding: "0 24px", marginTop: "30px" }}>
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
                    src="https://sun9-36.userapi.com/sun9-55/impf/c857520/v857520744/85ad6/CLDVzPGJEkQ.jpg?size=640x640&quality=96&proxy=1&sign=73231812f06bad04f453584f93a57c29&type=album"
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
