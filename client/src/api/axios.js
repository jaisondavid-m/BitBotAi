import axios from "axios"

export const api = axios.create({
    baseURL:"https://bitbotai.onrender.com",
});

api.interceptors.request.use((req) => {
  const token = localStorage.getItem("token");

  if (token) {
    req.headers.Authorization = token;
  }

  return req;
});
