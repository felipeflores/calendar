import React, { useEffect } from "react";
import styled from "styled-components";

import { useNavigate } from "react-router-dom";


import Steps from "../components/Steps"
import Button from "../components/Button"

const Start = () => {
    const history = useNavigate();
    const next = () => {
        history('/config/google/credentials');
    }



    return (
        <div>
             Conecte o device e clique em próximo
                    <Button title="Próximo" onClick={next}/>
                    <Steps/>
        </div>
    )
}
export default Start;