import api from "./api";

export const getPort = async () => {
    const ports = await api.get(`/v1/ports`)
    return ports?.data
}