import React from "react";
import styled from "styled-components";

import { Outlet } from "react-router-dom";

const Config = () => {

    return (
        <div class="container-sm container-fluid">
            <Wrapper>
                <Wizzard>
                    <Outlet />
                </Wizzard>
            </Wrapper>
        </div>
    );
};

export default Config;

const Wrapper = styled.div`
    max-width: 1400px;
    margin: auto;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
`;
const Wizzard = styled.div`
    position: relative;
    background: #fff;
    height: 400px;
    width: 738px;
    padding: 71px 93px 0;
    border-radius: 10px;
    box-shadow: 0px 2px 7px 0px rgb(0 0 0 / 10%);
`;
