import React from "react";

import { Link } from "react-router-dom";
import { PageHeader } from "antd";

import "./index.css";

export default class HistoryBreadcrumbs extends React.Component {
  render() {
    const { title, subTitle, routes } = this.props;

    return (
      <PageHeader
        className="site-page-header"
        title={title}
        subTitle={subTitle}
        breadcrumb={{
          routes: routes,
          itemRender: historyBreadcrumbsRender,
        }}
      />
    );
  }
}

function historyBreadcrumbsRender(route, params, routes, paths) {
  const last = routes.indexOf(route) === routes.length - 1;
  return last ? (
    <span>{route.breadcrumbName}</span>
  ) : (
    <Link to={paths.join("/")}>{route.breadcrumbName}</Link>
  );
}
