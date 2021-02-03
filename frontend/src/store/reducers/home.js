import {
  FETCH_USERS_COUNT,
  FETCH_USERS_COUNT_FAIL,
  FETCH_USERS_COUNT_SUCCESS,
} from "../actions/home";

const initialState = {
  usersCount: null,
  isLoading: false,
};

export default function (state = initialState, action) {
  const { type, payload } = action;

  switch (type) {
    case FETCH_USERS_COUNT:
      return {
        ...state,
        isLoading: true,
      };

    case FETCH_USERS_COUNT_FAIL:
      return {
        ...state,
        isLoading: false,
      };

    case FETCH_USERS_COUNT_SUCCESS:
      return {
        ...state,
        isLoading: false,
        usersCount: payload,
      };

    default:
      return state;
  }
}
