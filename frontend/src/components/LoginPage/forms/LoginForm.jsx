import React, { useRef, useState } from 'react';
import style from './Form.module.css'

const LoginForm = () => {


    const [passwordVisible, setpasswordVisible] = useState(false);
    const passwordInputRef = useRef(null);


    const [password, setPassword] = useState('');


    const togglePasswordVisibility = () => {
        setpasswordVisible(!passwordVisible);

        if (passwordInputRef.current) {
            passwordInputRef.current.focus();
        }
    }

    return (
        <>
            <h2 className={style.h2}>Login</h2>
            {/* <form method="POST" action="http://localhost:8080/create" className={style.test}>
                <div>
                    <label>Model</label>
                    <input type="text" name="model" />
                </div>
<<<<<<< HEAD
                <label className={style.label}>
                    <input type="checkbox" name="agree"></input>
                    <span>Remember me</span>
                </label>
                <button
                    className={style.button}
                // onClick={() => counted()}
                >
                    Login
                </button> */}
            <form method='POST' action="http://localhost:8080/create" className={style.form}>
                <input
                    className={style.input}
                    type="email"
                    autoComplete="username"
                    placeholder='Логин'
                />
                <div className={style.password__container}>
                    <svg
                        onClick={togglePasswordVisibility}
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        style={{ cursor: 'pointer' }}
                    >
                        <path d="M1 12s3-7 11-7 11 7 11 7-3 7-11 7-11-7-11-7z"></path>
                        <circle cx="12" cy="12" r="3"></circle>
                    </svg>
                    <input
                        ref={passwordInputRef}
                        className={style.input}
                        type={passwordVisible ? 'text' : 'password'}
                        autoComplete="current-password"
                        placeholder='Пароль'
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </div>
                <label className={style.label}>
                    <input type="checkbox" name="agree"></input>
                    <span>Remember me</span>
                </label>
                <button
                    className={style.button}
                // onClick={() => counted()}
                >
                    Login
                </button>
            </form>
        </>
    );
};

export default LoginForm;