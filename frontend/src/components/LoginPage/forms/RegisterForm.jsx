import React from 'react';
import style from './Form.module.css'

const RegisterForm = () => {
    return (
        <>
            <h2 className={style.h2}>Register</h2>
            <form className={style.form}>
                <input
                    className={style.input}
                    type="email"
                    autoComplete="username"
                    placeholder='Введите адрес эл.почты'
                />
                <input
                    className={style.input}
                    type="password"
                    autoComplete="current-password"
                    placeholder='Придумайте пароль'
                    
                />
                <input
                    className={style.input}
                    type="password"
                    autoComplete="current-password"
                    placeholder='Повторите пароль'
                />
                <button
                    className={style.button}
                // onClick={() => counted()}
                >
                    Register
                </button>
            </form>
        </>
    );
};

export default RegisterForm;