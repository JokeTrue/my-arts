import React from "react";
import { useCurrentUser } from "../../helpers/currentUserContext";

export default function HomePage() {
  const { user } = useCurrentUser();
  return <div>IN PROGRESS...</div>;
}
