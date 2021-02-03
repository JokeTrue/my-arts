import React, { useEffect } from "react";

import { useDispatch, useSelector } from "react-redux";

import { fetchProfile } from "../../store/actions/profile";
import HistoryBreadcrumbs from "../../components/Breadcrumbs";
import { useCurrentUser } from "../../helpers/currentUserContext";

import { Descriptions, Layout, Spin } from "antd";
import "./index.css";

export default function ProfilePage(props) {
  const routes = [
    {
      path: "home",
      breadcrumbName: "Home",
    },
    {
      path: "profile",
      breadcrumbName: "Profile",
    },
  ];

  const dispatch = useDispatch();
  const { user: currentUser } = useCurrentUser();

  const { user: profileUser, isLoading } = useSelector(
    (state) => state.Profile
  );

  useEffect(() => {
    dispatch(fetchProfile(props.match.params.id));
  }, [dispatch, props.match.params.id]);

  const isProfileMine =
    currentUser && profileUser && currentUser.id === profileUser.id;
  return (
    <>
      <HistoryBreadcrumbs
        routes={routes}
        title={
          isProfileMine
            ? "Profile"
            : profileUser &&
            `${profileUser.first_name} ${profileUser.last_name}'s Profile`
        }
        subTitle={isProfileMine ? "Your personal info" : "View personal info"}
      />

      <Layout style={{ height: "100%" }}>
        {isLoading && (
          <Spin
            style={{
              display: "inline-flex",
              justifyContent: "center",
              alignItems: "center",
              height: "100%",
            }}
          />
        )}

        {profileUser && (
          <Descriptions layout="vertical" bordered>
            <Descriptions.Item label="Email">
              {profileUser.email}
            </Descriptions.Item>
            <Descriptions.Item label="First Name">
              {profileUser.first_name}
            </Descriptions.Item>
            <Descriptions.Item label="Last Name">
              {profileUser.last_name}
            </Descriptions.Item>
            <Descriptions.Item label="Age">{profileUser.age}</Descriptions.Item>
            <Descriptions.Item label="Gender">
              {profileUser.gender}
            </Descriptions.Item>
            <Descriptions.Item label="Location">
              {profileUser.location}
            </Descriptions.Item>
            <Descriptions.Item label="Biography">
              {profileUser.biography}
            </Descriptions.Item>
          </Descriptions>
        )}
      </Layout>
    </>
  );
}
