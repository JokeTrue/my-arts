import {
  FETCH_FRIENDSHIP_REQUESTS,
  FETCH_FRIENDSHIP_REQUESTS_FAIL,
  FETCH_FRIENDSHIP_REQUESTS_SUCCESS,
} from "../actions/friendshipRequests";

const initialState = {
  requests: [],
  isLoading: false,
};

export default function (state = initialState, action) {
  const { type, payload } = action;

  switch (type) {
    case FETCH_FRIENDSHIP_REQUESTS:
      return {
        ...state,
        isLoading: true,
      };

    case FETCH_FRIENDSHIP_REQUESTS_FAIL:
      return {
        ...state,
        isLoading: false,
      };

    case FETCH_FRIENDSHIP_REQUESTS_SUCCESS:
      return {
        ...state,
        requests: payload,
        isLoading: false,
      };

    default:
      return state;
  }
}
