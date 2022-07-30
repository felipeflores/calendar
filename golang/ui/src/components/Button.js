import React from 'react';
import styled from "styled-components";

const Button = (props) => {
    const { title, ...remaining } = props;

    return (
        <Container>
            <Btn {...remaining}>
                { title }
            </Btn>
        </Container>
    )
}

export default Button
const Container = styled.div`
`;
const Btn = styled.a`
    padding: 0;
    border: none;
    display: inline-flex;
    height: 46px;
    width: 146px;
    align-items: center;
    background: #aac1f0;
    cursor: pointer;
    position: relative;
    padding-left: 33px;
    color: #fff;
    font-weight: 400;
    text-transform: uppercase;
    border-radius: 23px;
    margin-top: 10px;
    transform: perspective(1px) translateZ(0);
    transition-duration: 0.3s;
    text-decoration: none;
    ${Container}:hover & {
        background: #98add6;
        color: #fff;
    }
`;