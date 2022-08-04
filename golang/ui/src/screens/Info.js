import React, { useEffect, useState } from "react";
import styled from "styled-components";

import { useNavigate } from "react-router-dom";

import { getInfo } from "../service/esp.service"
import { setMqtt } from "../service/mqtt.service"

import Steps from "../components/Steps"
import Button from "../components/Button"

const INITAL_STATE = {
    info: {},
    name: "",
    broker: "",
    port: "1883",
    eventsCalendar: "",
}

const Start = () => {
    const history = useNavigate();
    const next = async () => {
        await setMqtt(form)


        history('/config/ports');
    }
    const [form, setForm] = useState(INITAL_STATE);
    

    useEffect(() => {
        (async () => {
            const i = await getInfo();
            setForm({
                ...form,
                info: i
            })
        })();
    },[])

    const change = (id, e) => {
        setForm({
            ...form,
            [id]: e.target.value,
        })
    }


    return (
        <div>
            <p>
                O nome do seu dispositivo é: {form.info.chip_id}
            </p>
            <div class="row">
                <div class="col-6">
                    <label>Qual o nome que você quer usar:</label>
                </div>
                <div class="col-6">
                    <input class="form-control form-control-lg" id="name" onChange={(e) => change("name", e)} value={form.name}/>
                </div>
            </div>
            <div class="row">
                <div class="col-6">
                    <label>Broker</label>
                </div>
                <div class="col-6">
                    <input class="form-control form-control-lg" id="broker" onChange={(e) => change("broker", e)} value={form.broker}/>
                </div>
            </div>
            <div class="row">
                <div class="col-6">
                    <label>Porta</label>
                </div>
                <div class="col-6">
                    <input class="form-control form-control-lg" id="port" onChange={(e) => change("port", e)} value={form.port}/>
                </div>
            </div>
            <div class="row">
                <div class="col-6">
                    <label>Evento Calendario</label> 
                </div>
                <div class="col-6">
                    <input class="form-control form-control-lg" id="eventsCalendar" onChange={(e) => change("eventsCalendar", e)} value={form.eventsCalendar}/>
                </div>
            </div>
            
            
                    <Button title="Próximo" onClick={next}/>
                    <Steps/>
        </div>
    )
}
export default Start;