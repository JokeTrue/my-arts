import Qs from "qs";
import axios from "axios";

const httpClient = axios.create({
  baseURL: "/api",
  validateStatus: (status) => status >= 200 && status < 403 && status !== 401,
  paramsSerializer: (params) => Qs.stringify(params, { arrayFormat: "repeat" }),
});

httpClient.interceptors.request.use(
  async (config) => {
    config.headers = { Accept: "application/json" };
    let token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    Promise.reject(error);
  }
);

httpClient.interceptors.response.use(
  (response) => {
    return response;
  },
  async function (error) {
    const originalRequest = error.config;
    if (error.response.status === 403 || error.response.status === 401) {
      localStorage.removeItem("token");
      window.location.reload();
      return httpClient(originalRequest);
    }
    return Promise.reject(error);
  }
);

export default httpClient;

export const callPost = (url, data) => {
  return httpClient.post(url, data);
};

export const callGet = (url, params) => {
  return httpClient.get(url, { params });
};
