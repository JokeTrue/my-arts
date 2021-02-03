import {
  FETCH_USERS,
  FETCH_USERS_FAIL,
  FETCH_USERS_SUCCESS,
} from "../actions/users";

const initialState = {
  users: [],
  isLoading: false,
  query: "",
  offset: 0,
  limit: 40,
  hasMore: true,
};

export default function (state = initialState, action) {
  const { type, payload } = action;

  switch (type) {
    case FETCH_USERS:
      const queryChanged = payload.query !== state.query;
      return {
        ...state,
        hasMore: true,
        isLoading: true,
        query: payload.query,
        offset: queryChanged ? 0 : state.offset,
        users: queryChanged ? [] : state.users,
      };

    case FETCH_USERS_FAIL:
      return {
        ...state,
        isLoading: false,
      };

    case FETCH_USERS_SUCCESS:
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
