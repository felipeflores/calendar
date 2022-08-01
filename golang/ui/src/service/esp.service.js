import api from "./api";

export const start = async (port) => {
    const ports = await api.post(`/v1/esp/${port}/start`)
    return ports?.data
}

export const doReset = async () => {
    const ports = await api.post(`/v1/esp/reset`)
    return ports?.data
}

export const getInfo = async () => {
    const networks = await api.get(`/v1/esp/info`)
    return networks?.data
}

export const getNetworks = async () => {
    const networks = await api.get(`/v1/esp/networks`)
    return networks?.data
}

export const setNetwork = async (network) => {
    const networks = await api.post(`/v1/esp/networks`, network)
    return networks?.data
}

