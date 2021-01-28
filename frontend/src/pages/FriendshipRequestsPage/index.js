import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";

import HistoryBreadcrumbs from "../../components/Breadcrumbs";
import FriendshipRequestsList from "./friendshipRequestsList";

import {
  actionFriendshipRequest,
  fetchFriendshipRequests,
} from "../../store/actions/friendshipRequests";
import { useCurrentUser } from "../../helpers/currentUserContext";

export default function FriendshipRequestsPage() {
  const routes = [
    {
      path: "home",
      breadcrumbName: "Home",
    },
    {
      path: "friendship_requests",
      breadcrumbName: "Requests",
    },
  ];

  const dispatch = useDispatch();
  const { user: currentUser } = useCurrentUser();

  const { requests, isLoading } = useSelector(
    (state) => state.FriendshipRequests
  );

  const onButtonAction = (action, requestId) => {
    dispatch(actionFriendshipRequest(action, currentUser.id, requestId));
  };

  useEffect(() => {
    if (currentUser) {
      dispatch(fetchFriendshipRequests(currentUser.id));
    }
  }, [dispatch, currentUser]);

  return (
    <>
      <HistoryBreadcrumbs
        routes={routes}
        title="Friendship Requests"
        subTitle="Explore friendship request from other users"
      />
      <FriendshipRequestsList
        requests={requests}
        isLoading={isLoading}
        onButtonAction={onButtonAction}
      />
    </>
  );
}
