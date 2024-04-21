import axios from 'axios';

// Create an Axios instance
const axiosInstance = axios.create({
  baseURL: 'http://localhost:8080/', // Base URL is taken from environment variables
  headers: {
    'Content-Type': 'application/json'
  }
});

console.log(axiosInstance.defaults.baseURL);

// Set up request interceptors
axiosInstance.interceptors.request.use(config => {
  // You can modify the request configuration before the request is sent
  // For example, attaching an authentication token:
  const token = sessionStorage.getItem('authToken');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
}, error => {
  // Do something with request error
  return Promise.reject(error);
});

// Set up response interceptors
axiosInstance.interceptors.response.use(response => {
  // Any status code that lie within the range of 2xx cause this function to trigger
  return response;
}, error => {
  // Any status codes that falls outside the range of 2xx cause this function to trigger
  // Do something with response error
  console.error('Something went wrong:', error);
  return Promise.reject(error);
});

export default axiosInstance;
