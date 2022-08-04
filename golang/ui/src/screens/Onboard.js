import React, { useEffect, useState } from "react";
import styled from "styled-components";

import { useNavigate } from "react-router-dom";

import { getInfo } from "../service/esp.service"

import Steps from "../components/Steps"
import Button from "../components/Button"

const INITAL_STATE = {
    info: {},
    
}

const Onboard = () => {
    const history = useNavigate();
    const next = () => {
        history('/config/ports');
    }
    const [form, setForm] = useState(INITAL_STATE);
    

    useEffect(() => {
        (async () => {
            const i = await getInfo();
            setForm({
                info: i
            })
        })();
    },[])


    return (
        <div>
             Onboard
        </div>
    )
}
export default Onboard;