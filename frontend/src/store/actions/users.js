export const FETCH_USERS = "@Users/FETCH_USERS";
export const FETCH_USERS_FAIL = "@Users/FETCH_USERS_FAIL";
export const FETCH_USERS_SUCCESS = "@Users/FETCH_USERS_SUCCESS";

export const fetchUsers = (query, offset, limit) => ({
  type: FETCH_USERS,
  payload: { query, offset, limit },
});
