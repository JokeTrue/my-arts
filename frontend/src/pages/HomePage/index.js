import React, { useEffect } from "react";

import { useDispatch, useSelector } from "react-redux";
import { fetchUsersCount } from "../../store/actions/home";

import { Card, Col, Row, Statistic } from "antd";
import { ArrowDownOutlined, ArrowUpOutlined } from "@ant-design/icons";
import { useCurrentUser } from "../../helpers/currentUserContext";
import HistoryBreadcrumbs from "../../components/Breadcrumbs";

export default function HomePage() {
  const routes = [
    {
      path: "home",
      breadcrumbName: "Home",
    },
  ];

  const { user } = useCurrentUser();
  const dispatch = useDispatch();

  useEffect(() => dispatch(fetchUsersCount()), [user?.id]);

  const { usersCount, isLoading } = useSelector((state) => state.Home);
  return (
    <>
      <HistoryBreadcrumbs
        routes={routes}
        title="Home"
        subTitle="Current statistics"
      />
      <Row style={{ justifyContent: "space-between" }}>
        <Col span={7}>
          <Card>
            <Statistic
              title="Active Users"
              loading={isLoading}
              value={usersCount}
            />
          </Card>
        </Col>
        <Col span={7}>
          <Card>
            <Statistic
              title="Active"
              loading={isLoading}
              value={11.28}
              precision={2}
              valueStyle={{ color: "#3f8600" }}
              prefix={<ArrowUpOutlined />}
              suffix="%"
            />
          </Card>
        </Col>
        <Col span={7}>
          <Card>
            <Statistic
              title="Idle"
              loading={isLoading}
              value={9.3}
              precision={2}
              valueStyle={{ color: "#cf1322" }}
              prefix={<ArrowDownOutlined />}
              suffix="%"
            />
          </Card>
        </Col>
      </Row>
    </>
  );
}
