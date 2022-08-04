import React, { useEffect, useState } from "react";
import styled from "styled-components";

import { useNavigate, useSearchParams  } from "react-router-dom";

import { setCredentials } from "../../service/google.service"

import Steps from "../../components/Steps"
import Button from "../../components/Button"

const INITAL_STATE = {
    credentials: "",
}

const Google = () => {
    const [searchParams, setSearchParams] = useSearchParams();

    const history = useNavigate();
    const next = async () => {
        const signup = await setCredentials(form.credentials)
        
        window.location.replace(signup.url);

        // history('/config/ports');
    }
    const [form, setForm] = useState(INITAL_STATE);
   
    const change = (id, e) => {
        setForm({
            ...form,
            [id]: e.target.value,
        })
    }

    return (
        <div>
             O nome do seu dispositivo é: 

             <textarea id="credentials" onChange={(e) => change("credentials", e)}>
                
             </textarea>

             
                    <Button title="Próximo" onClick={next}/>
                    <Steps/>
        </div>
    )
}
export default Google;