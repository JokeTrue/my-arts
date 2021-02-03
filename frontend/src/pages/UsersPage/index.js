import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchUsers } from "../../store/actions/users";
import HistoryBreadcrumbs from "../../components/Breadcrumbs";

import { Input } from "antd";
import Layout from "antd/es/layout/layout";
import { AudioOutlined } from "@ant-design/icons";
import UsersList from "../../components/UsersList";
import { fetchFriends } from "../../store/actions/friends";
import { useCurrentUser } from "../../helpers/currentUserContext";

const { Search } = Input;

export default function UsersPage(props) {
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

  const dispatch = useDispatch();
  const { user } = useCurrentUser();
  const [query, setQuery] = useState("");

  const {
    users,
    isLoading: isUsersLoading,
    offset,
    limit,
    hasMore,
  } = useSelector((state) => state.Users);
  const {
    users: friends,
    isLoading: isFriendsLoading,
  } = useSelector((state) => state.Friends);
  const loadMore = () => dispatch(fetchUsers(user.id, offset, limit));

  useEffect(() => {
    if (user) {
      dispatch(fetchFriends(user.id, friends.length, 10_000));
    }
  }, [dispatch, user]);

  useEffect(() => {
    dispatch(fetchUsers(query, offset, limit));
  }, [dispatch, query]);

  const isLoading = isUsersLoading || isFriendsLoading;
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
          onSearch={setQuery}
        />
        <UsersList
          users={users}
          dispatch={dispatch}
          isLoading={isLoading}
          currentUserId={user?.id}
          history={props.history}
          friendsIds={friends.map((item) => item.id)}
          hasMore={hasMore}
          loadMore={loadMore}
        />
      </Layout>
    </>
  );
}
