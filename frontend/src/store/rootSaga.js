import { all, fork } from "redux-saga/effects";

import {
  watchFetchCurrentUser,
  watchFetchToken,
  watchSignUp,
} from "./sagas/auth";

import { watchFetchUsers } from "./sagas/users";
import { watchFetchProfile } from "./sagas/profile";
import { watchFetchFriends } from "./sagas/friends";
import {
  watchActionFriendshipRequests,
  watchCreateFriendshipRequests,
  watchFetchFriendshipRequests,
} from "./sagas/friendshipRequests";

export default function* rootSaga() {
  yield all([
    fork(watchSignUp),
    fork(watchFetchToken),
    fork(watchFetchCurrentUser),

    fork(watchFetchUsers),
    fork(watchFetchProfile),

    fork(watchFetchFriends),
    fork(watchFetchFriendshipRequests),
    fork(watchCreateFriendshipRequests),
    fork(watchActionFriendshipRequests),
  ]);
}
