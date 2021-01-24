import { combineReducers } from "redux";

import authReducer from "./reducers/auth";

const rootReducer = combineReducers({
  Auth: authReducer,
});

export default rootReducer;
