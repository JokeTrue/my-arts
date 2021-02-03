import { put, takeEvery } from "redux-saga/effects";
import { callGet } from "../../http_client";

import openNotification from "../../helpers/openNotification";
import {
  FETCH_USERS_COUNT,
  FETCH_USERS_COUNT_FAIL,
  FETCH_USERS_COUNT_SUCCESS,
} from "../actions/home";

function* fetchUsersCount(action) {
  try {
    const res = yield callGet("/users/total_count");
    if (res.status !== 200) {
      openNotification("error", "Fetching Users Count Error", res.data.message);
      yield put({ type: FETCH_USERS_COUNT_FAIL });
      return;
    }
    yield put({ type: FETCH_USERS_COUNT_SUCCESS, payload: res.data });
  } catch (e) {
    openNotification("error", "Fetching Users Count Error", e.toString());
    yield put({ type: FETCH_USERS_COUNT_FAIL });
  }
}

function* watchFetchUsersCount() {
  yield takeEvery(FETCH_USERS_COUNT, fetchUsersCount);
}

export { watchFetchUsersCount };
