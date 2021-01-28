import React from "react";

import { Link } from "react-router-dom";
import { PageHeader } from "antd";

import "./index.css";

const HistoryBreadcrumbs = (props) => {
  const { title, subTitle, routes } = props;

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
};

export default HistoryBreadcrumbs;

const historyBreadcrumbsRender = (route, params, routes, paths) => {
  const last = routes.indexOf(route) === routes.length - 1;
  return last ? (
    <span>{route.breadcrumbName}</span>
  ) : (
    <Link to={paths.join("/")}>{route.breadcrumbName}</Link>
  );
};
