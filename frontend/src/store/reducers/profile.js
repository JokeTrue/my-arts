import {
  FETCH_PROFILE,
  FETCH_PROFILE_FAIL,
  FETCH_PROFILE_SUCCESS,
} from "../actions/profile";

const initialState = {
  user: undefined,
  isLoading: false,
};

export default function (state = initialState, action) {
  const { type, payload } = action;

  switch (type) {
    case FETCH_PROFILE:
      return {
        ...state,
        isLoading: true,
      };

    case FETCH_PROFILE_FAIL:
      return {
        ...state,
        isLoading: false,
      };

    case FETCH_PROFILE_SUCCESS:
      return {
        ...state,
        user: payload,
        isLoading: true,
      };

    default:
      return state;
  }
}
