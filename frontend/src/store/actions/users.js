export const FETCH_USERS = "@users/FETCH_USERS";
export const FETCH_USERS_FAIL = "@users/FETCH_USERS_FAIL";
export const FETCH_USERS_SUCCESS = "@users/FETCH_USERS_SUCCESS";

export const fetchUsers = (query) => ({
  type: FETCH_USERS,
  payload: { query },
});
