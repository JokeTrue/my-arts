export const FETCH_PROFILE = "@profile/FETCH_PROFILE";
export const FETCH_PROFILE_FAIL = "@profile/FETCH_PROFILE_FAIL";
export const FETCH_PROFILE_SUCCESS = "@profile/FETCH_PROFILE_SUCCESS";

export const fetchProfile = (id) => ({
  type: FETCH_PROFILE,
  payload: { id },
});
