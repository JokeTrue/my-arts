import axios from "axios";
import Qs from "qs";

const httpClient = axios.create({
  baseURL: "/api",
  validateStatus: (status) => ( status >= 200 && status < 403 ),
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

httpClient.interceptors.response.use((response) => {
  return response
}, async function (error) {
  const originalRequest = error.config;
  if (error.response.status === 403 && !originalRequest._retry) {
    originalRequest._retry = true;

    const data = await callPost("/auth/refresh_token")
    axios.defaults.headers.common['Authorization'] = 'Bearer ' + data.token;

    return httpClient(originalRequest);
  }
  return Promise.reject(error);
});

export default httpClient;

export const callPost = (url, data) => {
  return httpClient.post(url, data);
};

export const callGet = (url, params) => {
  return httpClient.get(url, { params });
};
