import {
  FETCH_USERS,
  FETCH_USERS_FAIL,
  FETCH_USERS_SUCCESS,
} from "../actions/users";

const initialState = {
  users: [],
  isLoading: false,
};

export default function (state = initialState, action) {
  const { type, payload } = action;

  switch (type) {
    case FETCH_USERS:
      return {
        ...state,
        isLoading: true,
      };

    case FETCH_USERS_FAIL:
      return {
        ...state,
        isLoading: false,
      };

    case FETCH_USERS_SUCCESS:
      return {
        ...state,
        users: payload,
        isLoading: false,
      };

    default:
      return state;
  }
}
