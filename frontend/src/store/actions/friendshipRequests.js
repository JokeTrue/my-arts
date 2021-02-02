export const FETCH_FRIENDSHIP_REQUESTS =
  "@Friendship/FETCH_FRIENDSHIP_REQUESTS";
export const FETCH_FRIENDSHIP_REQUESTS_FAIL =
  "@Friendship/FETCH_FRIENDSHIP_REQUESTS_FAIL";
export const FETCH_FRIENDSHIP_REQUESTS_SUCCESS =
  "@Friendship/FETCH_FRIENDSHIP_REQUESTS_SUCCESS";

export const fetchFriendshipRequests = (userId) => ({
  type: FETCH_FRIENDSHIP_REQUESTS,
  payload: { userId },
});

export const CREATE_FRIENDSHIP_REQUEST =
  "@Friendship/CREATE_FRIENDSHIP_REQUEST";
export const CREATE_FRIENDSHIP_REQUEST_FAIL =
  "@Friendship/CREATE_FRIENDSHIP_REQUEST_FAIL";
export const CREATE_FRIENDSHIP_REQUEST_SUCCESS =
  "@Friendship/CREATE_FRIENDSHIP_REQUEST_SUCCESS";

export const createFriendshipRequest = (userId) => ({
  type: CREATE_FRIENDSHIP_REQUEST,
  payload: { userId },
});

export const ACTION_FRIENDSHIP_REQUEST =
  "@Friendship/ACTION_FRIENDSHIP_REQUEST";
export const ACTION_FRIENDSHIP_REQUEST_FAIL =
  "@Friendship/ACTION_FRIENDSHIP_REQUEST_FAIL";
export const ACTION_FRIENDSHIP_REQUEST_SUCCESS =
  "@Friendship/ACTION_FRIENDSHIP_REQUEST";

export const actionFriendshipRequest = (friendshipAction, actorId, requestId) => ({
  type: ACTION_FRIENDSHIP_REQUEST,
  payload: { friendshipAction, requestId, actorId },
});
