import {
  FETCH_CURRENT_USER_FAIL,
  FETCH_CURRENT_USER_SUCCESS,
  FETCH_TOKEN_FAIL,
  FETCH_TOKEN_SUCCESS,
  LOGOUT,
  SIGN_UP_FAIL,
  SIGN_UP_SUCCESS,
} from "../actions/auth";

const initialState = {
  user: undefined,
  isLoggedIn: false,
};

export default function (state = initialState, action) {
  const { type, payload } = action;

  switch (type) {
    case FETCH_TOKEN_SUCCESS:
      return {
        ...state,
        isLoggedIn: true,
      };
    case FETCH_TOKEN_FAIL:
      return {
        ...state,
        isLoggedIn: false,
      };

    case FETCH_CURRENT_USER_SUCCESS:
      return {
        ...state,
        user: payload,
        isLoggedIn: true,
      };
    case FETCH_CURRENT_USER_FAIL:
      return {
        ...state,
        user: undefined,
        isLoggedIn: false,
      };

    case SIGN_UP_SUCCESS:
      return {
        ...state,
        isLoggedIn: true,
      };
    case SIGN_UP_FAIL:
      return {
        ...state,
        isLoggedIn: false,
      };

    case LOGOUT:
      return {
        ...state,
        isLoggedIn: false,
      };

    default:
      return state;
  }
}
