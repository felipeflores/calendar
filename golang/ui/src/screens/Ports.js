import React, { useEffect, useState } from "react";
import styled from "styled-components";

import { useNavigate } from "react-router-dom";

import { getPorts } from "../service/port.service"

import Steps from "../components/Steps"
import Button from "../components/Button"

const INITAL_STATE = {
    port: ""
}
const Ports = () => {
    const history = useNavigate();
    const next = () => {
        history('/config/networks');
    }

    const [form, setForm] = useState(INITAL_STATE);
    const [ports, setPorts] = useState([]);

    useEffect(() => {
        (async () => {
            const p = await getPorts();
            setPorts(p);
            console.log(p)
        })();
    },[])


    return (
        <div>
            Portas Disponíveis
            <select className="form-select" >
                <option selected>Selecione a porta do seu dispositivo</option>
                {ports?.map((object, i) => <option key={i} value={object}>{object}</option>)}
            </select>
            <Button title="Próximo" onClick={next}/>
            <Steps/>

        </div>
    )
}
export default Ports;