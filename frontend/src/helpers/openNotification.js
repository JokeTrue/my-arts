import React from "react";

import { notification } from "antd";

const openNotification = (type, title, description) => {
  notification[type]({
    duration: 30,
    message: title,
    description: description,
  });
};

export default openNotification;
