import axios from "axios"

const instance = axios.create({
  baseURL: 'http://127.0.0.1:8000',
  withCredentials: true
});

instance.interceptors.request.use(function (config) {
    console.log(config);
  const token = localStorage.getItem("token")
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
    return config;
  }, function (error) {
    console.log(error);
    return Promise.reject(error);
  });
instance.interceptors.response.use(function (response) {
    console.log(response);
    
    return response;
  }, function (error) {
    console.log(error);
    
    return Promise.reject(error);
});
  
export default instance