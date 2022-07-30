import api from "./api";

export const getPorts = async () => {
    const ports = await api.get(`/v1/ports`)
    return ports?.data
}