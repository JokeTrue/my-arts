import { all, fork } from "redux-saga/effects";

import {
  watchFetchCurrentUser,
  watchFetchToken,
  watchSignUp,
} from "./sagas/auth";

export default function* rootSaga() {
  yield all([
    fork(watchSignUp),
    fork(watchFetchToken),
    fork(watchFetchCurrentUser),
  ]);
}
