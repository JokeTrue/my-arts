export const FETCH_FRIENDS = "@Friendship/FETCH_FRIENDS";
export const FETCH_FRIENDS_FAIL = "@Friendship/FETCH_FRIENDS_FAIL";
export const FETCH_FRIENDS_SUCCESS = "@Friendship/FETCH_FRIENDS_SUCCESS";

export const fetchFriends = (userId, offset, limit) => ({
  type: FETCH_FRIENDS,
  payload: { userId, offset, limit },
});
