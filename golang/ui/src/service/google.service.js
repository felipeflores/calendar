import api from "./api";

export const setCredentials = async (credentials) => {
    const signup = await api.post(`/v1/google/signup`, credentials)
    return signup?.data
}

export const setCode = async (code) => {
    const response = await api.post(`/v1/google/code`, code)
    return response?.data
}