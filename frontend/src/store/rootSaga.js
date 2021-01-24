import { all, fork } from 'redux-saga/effects';

import { watchFetchToken, watchFetchCurrentUser } from "./sagas/auth";


export default function* rootSaga() {
  yield all([
    fork(watchFetchToken),
    fork(watchFetchCurrentUser),
  ]);
}
