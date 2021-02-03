import { put, takeEvery } from "redux-saga/effects";
import { callGet } from "../../http_client";

import openNotification from "../../helpers/openNotification";
import {
  FETCH_FRIENDS,
  FETCH_FRIENDS_FAIL,
  FETCH_FRIENDS_SUCCESS,
} from "../actions/friends";

function* fetchFriendsSaga(action) {
  try {
    const { userId, offset, limit } = action.payload;
    const res = yield callGet(`/users/user/${userId}/friends`, {
      offset,
      limit,
    });
    if (res.status !== 200) {
      openNotification("error", "Fetching Friends Error", res.data.message);
      yield put({ type: FETCH_FRIENDS_FAIL });
      return;
    }
    yield put({ type: FETCH_FRIENDS_SUCCESS, payload: res.data });
  } catch (e) {
    openNotification("error", "Fetching Users Friends", e.toString());
    yield put({ type: FETCH_FRIENDS_FAIL });
  }
}

function* watchFetchFriends() {
  yield takeEvery(FETCH_FRIENDS, fetchFriendsSaga);
}

export { watchFetchFriends };
