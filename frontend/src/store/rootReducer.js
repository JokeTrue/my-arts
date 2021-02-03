import { combineReducers } from "redux";

import homeReducer from "./reducers/home";
import authReducer from "./reducers/auth";
import usersReducer from "./reducers/users";
import profileReducer from "./reducers/profile";
import friendsReducer from "./reducers/friends";
import friendshipRequestsReducer from "./reducers/friendshipRequests";

const rootReducer = combineReducers({
  Home: homeReducer,
  Auth: authReducer,
  Users: usersReducer,
  Profile: profileReducer,
  Friends: friendsReducer,
  FriendshipRequests: friendshipRequestsReducer,
});

export default rootReducer;
