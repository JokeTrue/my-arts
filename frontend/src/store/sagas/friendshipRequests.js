import { put, takeEvery, takeLeading } from "redux-saga/effects";

import { callGet, callPost } from "../../http_client";
import openNotification from "../../helpers/openNotification";

import {
  ACTION_FRIENDSHIP_REQUEST,
  ACTION_FRIENDSHIP_REQUEST_FAIL,
  ACTION_FRIENDSHIP_REQUEST_SUCCESS,
  CREATE_FRIENDSHIP_REQUEST,
  CREATE_FRIENDSHIP_REQUEST_FAIL,
  CREATE_FRIENDSHIP_REQUEST_SUCCESS,
  FETCH_FRIENDSHIP_REQUESTS,
  FETCH_FRIENDSHIP_REQUESTS_FAIL,
  FETCH_FRIENDSHIP_REQUESTS_SUCCESS,
} from "../actions/friendshipRequests";

function* fetchFriendshipRequestsSaga(action) {
  try {
    const { userId } = action.payload;
    const res = yield callGet("/friendship/list/" + userId);
    if (res.status !== 200) {
      openNotification(
        "error",
        "Fetching Friendship Requests Error",
        res.data.message
      );
      yield put({ type: FETCH_FRIENDSHIP_REQUESTS_FAIL });
      return;
    }
    yield put({ type: FETCH_FRIENDSHIP_REQUESTS_SUCCESS, payload: res.data });
  } catch (e) {
    openNotification(
      "error",
      "Fetching Friendship Requests Error",
      e.toString()
    );
    yield put({ type: FETCH_FRIENDSHIP_REQUESTS_FAIL });
  }
}

function* createFriendshipRequestSaga(action) {
  try {
    const { userId } = action.payload;
    const res = yield callPost("/friendship/create/" + userId);
    if (res.status !== 201) {
      openNotification(
        "error",
        "Creating Friendship Request Error",
        res.data.error
      );
      yield put({ type: CREATE_FRIENDSHIP_REQUEST_FAIL });
      return;
    }
    yield put({ type: CREATE_FRIENDSHIP_REQUEST_SUCCESS });
    openNotification("success", "Friendship Request was created!");
  } catch (e) {
    openNotification(
      "error",
      "Creating Friendship Request Error",
      e.toString()
    );
    yield put({ type: CREATE_FRIENDSHIP_REQUEST_FAIL });
  }
}

function* actionFriendshipRequestSaga(action) {
  try {
    const { friendshipAction, requestId, actorId } = action.payload;
    const res = yield callPost(
      `/friendship/action/${friendshipAction}/${requestId}`
    );
    if (res.status !== 200) {
      openNotification(
        "error",
        "Action on Friendship Request Error",
        res.data.error
      );
      yield put({ type: ACTION_FRIENDSHIP_REQUEST_FAIL });
      return;
    }

    yield put({ type: ACTION_FRIENDSHIP_REQUEST_SUCCESS });
    openNotification(
      "success",
      `Friendship Request was successfully ${friendshipAction}${
        friendshipAction === "accept" ? "ed" : "d"
      }`
    );
    yield put({
      type: FETCH_FRIENDSHIP_REQUESTS,
      payload: { userId: actorId },
    });
  } catch (e) {
    openNotification(
      "error",
      "Action on Friendship Request Error",
      e.toString()
    );
    yield put({ type: ACTION_FRIENDSHIP_REQUEST_FAIL });
  }
}

function* watchFetchFriendshipRequests() {
  yield takeEvery(FETCH_FRIENDSHIP_REQUESTS, fetchFriendshipRequestsSaga);
}

function* watchCreateFriendshipRequests() {
  yield takeEvery(CREATE_FRIENDSHIP_REQUEST, createFriendshipRequestSaga);
}

function* watchActionFriendshipRequests() {
  yield takeLeading(ACTION_FRIENDSHIP_REQUEST, actionFriendshipRequestSaga);
}

export {
  watchFetchFriendshipRequests,
  watchCreateFriendshipRequests,
  watchActionFriendshipRequests,
};
