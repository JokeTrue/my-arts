import React from "react";

export const CurrentUserContext = React.createContext({ user: {} });

export const useCurrentUser = () => React.useContext(CurrentUserContext)