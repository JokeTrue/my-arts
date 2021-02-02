export const FETCH_PROFILE = "@Profile/FETCH_PROFILE";
export const FETCH_PROFILE_FAIL = "@Profile/FETCH_PROFILE_FAIL";
export const FETCH_PROFILE_SUCCESS = "@Profile/FETCH_PROFILE_SUCCESS";

export const fetchProfile = (id) => ({
  type: FETCH_PROFILE,
  payload: { id },
});
