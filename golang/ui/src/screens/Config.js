import React, { useEffect } from "react";

import { getPort } from "../service/teste.service"

const Config = () => {

    useEffect(() => {
        console.log(getPort())
    },[])


    return (
        <div>teste</div>
    );
};

export default Config;