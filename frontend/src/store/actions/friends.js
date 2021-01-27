export const FETCH_FRIENDS = "@users/FETCH_FRIENDS";
export const FETCH_FRIENDS_FAIL = "@users/FETCH_FRIENDS_FAIL";
export const FETCH_FRIENDS_SUCCESS = "@users/FETCH_FRIENDS_SUCCESS";

export const fetchFriends = (userId) => ({
  type: FETCH_FRIENDS,
  payload: { userId },
});
