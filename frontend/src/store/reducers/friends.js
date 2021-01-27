import {
  FETCH_FRIENDS,
  FETCH_FRIENDS_FAIL,
  FETCH_FRIENDS_SUCCESS,
} from "../actions/friends";

const initialState = {
  users: [],
  isLoading: false,
};

export default function (state = initialState, action) {
  const { type, payload } = action;

  switch (type) {
    case FETCH_FRIENDS:
      return {
        ...state,
        isLoading: true,
      };

    case FETCH_FRIENDS_FAIL:
      return {
        ...state,
        isLoading: false,
      };

    case FETCH_FRIENDS_SUCCESS:
      return {
        ...state,
        users: payload,
        isLoading: false,
      };

    default:
      return state;
  }
}
