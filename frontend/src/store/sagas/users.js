import { put, takeEvery } from "redux-saga/effects";
import { callGet } from "../../http_client";

import openNotification from "../../helpers/openNotification";
import {
  FETCH_USERS,
  FETCH_USERS_FAIL,
  FETCH_USERS_SUCCESS,
} from "../actions/users";

function* fetchUsersSaga(action) {
  try {
    const res = yield callGet("/users/search", action.payload);
    if (res.status !== 200) {
      openNotification("error", "Fetching Users Error", res.data.message);
      yield put({ type: FETCH_USERS_FAIL });
      return;
    }
    yield put({ type: FETCH_USERS_SUCCESS, payload: res.data });
  } catch (e) {
    openNotification("error", "Fetching Users Error", e.toString());
    yield put({ type: FETCH_USERS_FAIL });
  }
}

function* watchFetchUsers() {
  yield takeEvery(FETCH_USERS, fetchUsersSaga);
}

export { watchFetchUsers };
