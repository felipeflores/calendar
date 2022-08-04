import api from "./api";

export const setMqtt = async (mqtt) => {
    const response = await api.post(`/v1/mqtt`, mqtt)
    return response?.data
}