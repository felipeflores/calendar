import React from 'react';
import styled from "styled-components";

const Steps = (props) => {

    const { step } = props;

    return (
        <ContainerSteps>
            <TabList>
                <PointerStep>
                    <Step className="current"/>
                    <Step/>
                    <Step/>
                    <Step/>
                </PointerStep>
            </TabList>
        </ContainerSteps>
    )
}

export default Steps

const ContainerSteps = styled.div`
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    bottom: -50px;
`;

const TabList = styled.ul`
    display: flex;
    padding: 0;
    margin: 0;
    list-style: none;
`;

const PointerStep = styled.li`
`;

const Step = styled.div`
    display: inline-block;
    width: 10px;
    height: 10px;
    background: #ffffff;
    border-radius: 50%;
    margin-right: 8px;
`;