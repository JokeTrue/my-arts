import { put, takeEvery } from "redux-saga/effects";

import { callGet, callPost } from "../../http_client";
import openNotification from "../../helpers/openNotification";

import {
  FETCH_CURRENT_USER,
  FETCH_CURRENT_USER_FAIL,
  FETCH_CURRENT_USER_SUCCESS,
  FETCH_TOKEN,
  FETCH_TOKEN_FAIL,
  FETCH_TOKEN_SUCCESS,
  fetchCurrentUser,
  fetchToken,
  SIGN_UP,
  SIGN_UP_FAIL,
  SIGN_UP_SUCCESS,
} from "../actions/auth";

function* fetchTokenSaga(action) {
  try {
    const res = yield callPost("/login", action.payload);
    if (res.status !== 200) {
      openNotification("error", "Authentication Error", res.data.message);
      yield put({ type: FETCH_TOKEN_FAIL });
      return;
    }

    localStorage.setItem("token", res.data.token);
    yield put({ type: FETCH_TOKEN_SUCCESS });
    yield put(fetchCurrentUser());
  } catch (e) {
    openNotification("error", "Authentication Error", e.toString());
    yield put({ type: FETCH_TOKEN_FAIL });
  }
}

function* fetchCurrentUserSaga(action) {
  let token = localStorage.getItem("token");
  if (!token) {
    openNotification("error", "Authentication Error", "JWT Token is missing");
    yield put({ type: FETCH_CURRENT_USER_FAIL });
  }

  try {
    const res = yield callGet("/users/me");
    if (res.data.error) {
      openNotification("error", "Authentication Error", res.data.error);
      yield put({ type: FETCH_CURRENT_USER_FAIL });
      return;
    }
    yield put({ type: FETCH_CURRENT_USER_SUCCESS, payload: res.data });
  } catch (e) {
    openNotification("error", "Authentication Error", e.toString());
    yield put({ type: FETCH_CURRENT_USER_FAIL });
  }
}

function* signUpSaga(action) {
  debugger;
  const { payload } = action;
  const { email, password1 } = payload;

  try {
    const res = yield callPost("/sign_up", payload);
    if (res.data.error) {
      openNotification("error", "Authentication Error", res.data.error);
      yield put({ type: SIGN_UP_FAIL });
      return;
    }
    yield put({ type: SIGN_UP_SUCCESS });
    yield put(fetchToken(email, password1));
  } catch (e) {
    openNotification("error", "Authentication Error", e.toString());
    yield put({ type: SIGN_UP_FAIL });
  }
}

function* watchFetchToken() {
  yield takeEvery(FETCH_TOKEN, fetchTokenSaga);
}

function* watchFetchCurrentUser() {
  yield takeEvery(FETCH_CURRENT_USER, fetchCurrentUserSaga);
}

function* watchSignUp() {
  yield takeEvery(SIGN_UP, signUpSaga);
}

export { watchFetchToken, watchFetchCurrentUser, watchSignUp };
