import React, { useEffect, useState } from "react";
import styled from "styled-components";

import { useNavigate } from "react-router-dom";

import { getPorts } from "../service/port.service"
import { start, doReset } from "../service/esp.service"

import Steps from "../components/Steps"
import Button from "../components/Button"

const INITAL_STATE = {
    port: ""
}
const Ports = () => {
    const history = useNavigate();
    const next = async () => {
        await start(form.port)
        await doReset()
        history('/config/info');
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

    const change = (e) => {
        setForm({
            [e.target.name]: e.target.value
        })
    }

    return (
        <div>
            Portas Disponíveis
            <select className="form-select" 
                id="port" name="port"
                onChange={change} >
                <option selected>Selecione a porta do seu dispositivo</option>
                {ports?.map((object, i) => <option key={i} value={object}>{object}</option>)}
            </select>
            <Button title="Próximo" onClick={next}/>
            <Steps/>

        </div>
    )
}
export default Ports;