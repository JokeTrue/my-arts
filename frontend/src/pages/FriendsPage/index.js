import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

import HistoryBreadcrumbs from "../../components/Breadcrumbs";
import { fetchFriends } from "../../store/actions/friends";
import UsersList from "../../components/UsersList";
import { useCurrentUser } from "../../helpers/currentUserContext";

export default function FriendsPage(props) {
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

  const dispatch = useDispatch();
  const { user } = useCurrentUser();

  const { users: friends, isLoading, offset, limit, hasMore } = useSelector(
    (state) => state.Friends
  );
  const loadMore = () => dispatch(fetchFriends(user.id, offset, limit));

  useEffect(() => {
    if (user !== null) {
      dispatch(fetchFriends(user.id, offset, limit));
    }
  }, [dispatch, user]);

  return (
    <>
      <HistoryBreadcrumbs
        routes={routes}
        title="Friends"
        subTitle="Your closest friends"
      />
      <UsersList
        users={friends}
        isLoading={isLoading}
        currentUserId={user?.id}
        history={props.history}
        friendsIds={friends.map((item) => item.id)}
        hasMore={hasMore}
        loadMore={loadMore}
      />
    </>
  );
}
