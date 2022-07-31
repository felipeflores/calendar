import React, { useEffect, useState } from "react";
import styled from "styled-components";

import { useNavigate } from "react-router-dom";


import { getNetworks, setNetwork } from "../service/esp.service"

import Steps from "../components/Steps"
import Button from "../components/Button"

const INITAL_STATE = {
    network: "",
    password: "",
}

const INITAL_STATE_NETWORKS = {
    networks: [],
}
const Networks = () => {
    const history = useNavigate();
    const next = async () => {
        await setNetwork(form)
        history('/config/networks');
    }

    const [form, setForm] = useState(INITAL_STATE);
    const [networks, setNetworks] = useState(INITAL_STATE_NETWORKS);

    useEffect(() => {
        (async () => {
            const n = await getNetworks();
            setNetworks(n);
            console.log(n)
        })();
    },[])

    const selectNetwork = (e) => {
        setForm({
            network: e
        })
    }
    const change = (id, e) => {
        console.log(id, e)
        setForm({
            ...form,
            [id]: e.target.value
        })
    }

    return (
        <div>
            Portas Disponíveis
            <Table>
                <Thead>
                    <Tr>
                        <Th style={{ width: "70%" }}>Rede</Th>
                        <Th style={{ width: "20%" }}>Sinal</Th>
                        <Th style={{ width: "10%" }}>Segura</Th>
                    </Tr>
                </Thead>
                <Tbody>
                    {networks?.networks?.map((network, i) => (
                        <Tr key={i} onClick={() => selectNetwork(network.ssi)}>
                            <Td style={{ width: "70%" }}>{network.ssi}</Td>
                            <Td style={{ width: "20%" }}>{network.rssi}</Td>
                            <Td style={{ width: "10%" }}>{network.encripted}</Td>
                        </Tr>)
                    )}
                </Tbody>
            </Table>
            {
                form.network &&
                    <div class="mb-3">
                        <label for="password" class="form-label">Password</label>
                        <input type="password" class="form-control" 
                            id="password" onChange={(e) => change("password", e)}/>
                    </div>
            }
            <Button title="Próximo" onClick={next}/>
            <Steps/>

        </div>
    )
}
export default Networks;

const Table = styled.table`
    width:80%;
`;

const Tr = styled.tr`
`;

const Th = styled.th`
`;

const Thead = styled.thead`
    display: block;
`;
const Tbody = styled.tbody`
    display: block;
    height: 100px;       /* Just for the demo          */
    overflow-y: auto;    /* Trigger vertical scroll    */
    overflow-x: hidden;  /* Hide the horizontal scroll */
`;

const Td = styled.td`
    cursor: pointer;
`;