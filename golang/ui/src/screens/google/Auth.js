import React, { useEffect, useState } from "react";
import styled from "styled-components";

import { useNavigate, useSearchParams } from "react-router-dom";

import { setCode } from "../../service/google.service"

import Steps from "../../components/Steps"
import Button from "../../components/Button"

const INITAL_STATE = {
    code: "",
}


const Auth = () => {
    const [form, setForm] = useState(INITAL_STATE);

    const [searchParams, setSearchParams] = useSearchParams();

    const history = useNavigate();
    const next = async () => {
        await setCode(form)
        history('/config/ports');
    }

    useEffect(() => {
        const code = searchParams.get('code')
        console.log(code)
        setForm({
            ...form,
            code: code
        })
    }, [])


    return (
        <div>
             Auth
                    <Button title="Próximo" onClick={next}/>
                    <Steps/>
        </div>
    )
}
export default Auth;