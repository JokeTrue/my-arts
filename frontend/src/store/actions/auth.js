export const FETCH_TOKEN = "@Auth/FETCH_TOKEN";
export const FETCH_TOKEN_FAIL = "@Auth/FETCH_TOKEN_FAIL";
export const FETCH_TOKEN_SUCCESS = "@Auth/FETCH_TOKEN_SUCCESS";

export const FETCH_CURRENT_USER = "@Auth/FETCH_CURRENT_USER";
export const FETCH_CURRENT_USER_FAIL = "@Auth/FETCH_CURRENT_USER_FAIL";
export const FETCH_CURRENT_USER_SUCCESS = "@Auth/FETCH_CURRENT_USER_SUCCESS";

export const SIGN_UP = "@Auth/SIGN_UP";
export const SIGN_UP_FAIL = "@Auth/SIGN_UP_FAIL";
export const SIGN_UP_SUCCESS = "@Auth/SIGN_UP_SUCCESS";

export const LOGOUT = "@Auth/LOGOUT";

export const fetchToken = (username, password) => ({
  type: FETCH_TOKEN,
  payload: { username, password },
});

export const fetchCurrentUser = () => ({
  type: FETCH_CURRENT_USER,
});

export const signUp = (username, password) => ({
  type: SIGN_UP,
  payload: { username, password },
});

export const logout = () => (dispatch) => {
  localStorage.removeItem("token");
  dispatch({ type: LOGOUT });
};
