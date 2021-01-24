import { notification } from "antd";
import React from "react";

const openNotification = (type, title, description) => {
  notification[type]({
    duration: 30,
    message: title,
    description: description,
  });
};

export default openNotification;
