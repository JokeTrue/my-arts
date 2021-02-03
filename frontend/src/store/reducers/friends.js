import {
  FETCH_FRIENDS,
  FETCH_FRIENDS_FAIL,
  FETCH_FRIENDS_SUCCESS,
} from "../actions/friends";

const initialState = {
  users: [],
  isLoading: false,
  offset: 0,
  limit: 100,
  hasMore: true,
};

export default function (state = initialState, action) {
  const { type, payload } = action;

  switch (type) {
    case FETCH_FRIENDS:
      return {
        ...state,
        hasMore: true,
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
        isLoading: false,
        users: [...state.users, ...payload],
        offset: state.offset + payload.length,
        hasMore: payload.length === state.limit,
      };

    default:
      return state;
  }
}
