import { put, takeEvery } from "redux-saga/effects";
import { callGet } from "../../http_client";

import openNotification from "../../helpers/openNotification";
import {
  FETCH_PROFILE,
  FETCH_PROFILE_FAIL,
  FETCH_PROFILE_SUCCESS,
} from "../actions/profile";

function* fetchProfileSaga(action) {
  const { id } = action.payload;
  try {
    const res = yield callGet("/users/user/" + id);
    if (res.status !== 200) {
      openNotification("error", "Fetching Profile Error", res.data.message);
      yield put({ type: FETCH_PROFILE_FAIL });
      return;
    }
    yield put({ type: FETCH_PROFILE_SUCCESS, payload: res.data });
  } catch (e) {
    openNotification("error", "Fetching Profile Error", e.toString());
    yield put({ type: FETCH_PROFILE_FAIL });
  }
}

function* watchFetchProfile() {
  yield takeEvery(FETCH_PROFILE, fetchProfileSaga);
}

export { watchFetchProfile };
