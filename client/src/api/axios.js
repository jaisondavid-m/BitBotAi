import axios from "axios"

export const api = axios.create({
    baseURL:"http://localhost:8000",
});

api.interceptors.request.use((req) => {
  const token = localStorage.getItem("token");

  if (token) {
    req.headers.Authorization = `Bearer ${token}`;
  }

  return req;
});
