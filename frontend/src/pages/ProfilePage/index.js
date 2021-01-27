import React from "react";

import { connect } from "react-redux";

import { fetchProfile } from "../../store/actions/profile";
import HistoryBreadcrumbs from "../../components/Breadcrumbs";

import { Descriptions, Layout, Spin } from "antd";
import "./index.css";

@connect((store) => ({
  user: store.Auth.user,
  profile: store.Profile.user,
  isLoggedIn: store.Auth.isLoggedIn,
}))
class ProfilePage extends React.Component {
  componentDidMount() {
    const { id } = this.props.match.params;
    this.props.dispatch(fetchProfile(id));
  }

  componentDidUpdate(nextProps) {
    const { id: newId } = this.props.match.params;
    const { id: oldId } = nextProps.match.params;

    if (newId !== oldId) {
      this.props.dispatch(fetchProfile(newId));
    }
  }

  render() {
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

    const { profile, user } = this.props;
    const isProfileMine = user && profile && user.id === profile.id;
    return (
      <>
        <HistoryBreadcrumbs
          routes={routes}
          title={
            isProfileMine
              ? "Profile"
              : profile &&
                `${profile.first_name} ${profile.last_name}'s Profile`
          }
          subTitle={isProfileMine ? "Your personal info" : "View personal info"}
        />

        <Layout style={{ height: "100%" }}>
          {!profile && (
            <Spin
              style={{
                display: "inline-flex",
                justifyContent: "center",
                alignItems: "center",
                height: "100%",
              }}
            />
          )}
          {profile && (
            <Descriptions
              layout="vertical"
              bordered
              style={{ marginTop: "30px" }}
            >
              <Descriptions.Item label="Email">
                {profile.email}
              </Descriptions.Item>
              <Descriptions.Item label="First Name">
                {profile.first_name}
              </Descriptions.Item>
              <Descriptions.Item label="Last Name">
                {profile.last_name}
              </Descriptions.Item>
              <Descriptions.Item label="Age">{profile.age}</Descriptions.Item>
              <Descriptions.Item label="Gender">
                {profile.gender}
              </Descriptions.Item>
              <Descriptions.Item label="Location">
                {profile.location}
              </Descriptions.Item>
              <Descriptions.Item label="Biography">
                {profile.biography}
              </Descriptions.Item>
            </Descriptions>
          )}
        </Layout>
      </>
    );
  }
}

export default ProfilePage;
