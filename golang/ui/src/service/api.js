import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080",
  headers: {
    "Content-Type": "application/json",
  },
});

// api.interceptors.request.use(
//   async (req) => {
//     const authToken = localStorage.getItem("access_token");
//     if (authToken) req.headers.Authorization = authToken;
//     return req;
//   },
//   (err) => {
//     return Promise.reject(err);
//   }
// );

export default api;