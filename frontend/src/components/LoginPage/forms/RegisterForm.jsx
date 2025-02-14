import React, { useState } from 'react';
import style from './Form.module.css'

const RegisterForm = () => {


    const [input1, setInput1] = useState('');
    const [input2, setInput2] = useState('');
    const [match, setMatch] = useState(null);

    const checkInput = (event) => {
        event.preventDefault();


        if (input1 === input2) {
            setMatch(true)

        event.target.submit();
        } else {
            setMatch(false)
        }
    }


    return (
        <>
            <h2 className={style.h2}>Register</h2>
            <form  method='POST' action="http://localhost:8080/create" onSubmit={checkInput}  className={style.form}>
            <input
                className={style.input}
                type="email"
                name="email"
                autoComplete="email"
                placeholder='Введите адрес эл.почты'
            />
            <input
                className={style.input}
                type="password"
                name="password"
                autoComplete="current-password"
                placeholder='Придумайте пароль'
                onChange={(e) => setInput1(e.target.value)}
            />
            <input
                className={style.input}
                type="password"
                name="current-password"
                autoComplete="current-password"
                placeholder='Повторите пароль'
                onChange={(e) => setInput2(e.target.value)}
            />
            {match !== null && (
                <div>
                    {match ? <div></div> : <div>Пароли не совпадают</div>}
                </div>
            )}
            <button
                className={style.button}
                type='submit'
            >
                Register
            </button>
            </form>
        </>
    );
};

export default RegisterForm;