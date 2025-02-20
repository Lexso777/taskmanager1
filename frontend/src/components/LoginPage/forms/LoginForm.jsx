import React, { useState, useRef } from 'react';
import style from './Form.module.css';
import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { fetchLogin } from '../../../Redux/slices/LoginSlice';

const LoginForm = () => {

    const [passwordVisible, setPasswordVisible] = useState(false);
    const navigate = useNavigate();
    const {error} = useSelector((state) => state.LoginSlice)
    const dispatch = useDispatch();
    const passwordInputRef = useRef(null);


    const handleSubmit = (event) => {
        event.preventDefault();

        const formData = new FormData(event.target);
        const data = {
            email: formData.get("email"),
            password: formData.get("password"),
            remember: formData.get("agree") ? true : false
        };

        dispatch(fetchLogin(data)).then((result) => {
            if(result.meta.requestStatus === 'fulfilled'){
                navigate('/main');
            }else{
                console.log("Ошибка авторизации");
            }
        });
    }   


    const togglePasswordVisibility = () => {
        setPasswordVisible(!passwordVisible);
        passwordInputRef.current.focus();
    };

    return (
        <>
            <h2 className={style.h2}>Login</h2>
            <form method="POST" action="http://localhost:8080/login" onSubmit={handleSubmit} className={style.form}>
                <input
                    className={style.input}
                    type="email"
                    name="email"
                    autoComplete="username"
                    placeholder="Логин"
                    required
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
                        type={passwordVisible ? "text" : "password"}
                        autoComplete="current-password"
                        name="password"
                        placeholder="Пароль"
                        required
                    />
                </div>
                <div className={style.error}>{error}</div>
                <label className={style.label}>
                    <input type="checkbox" name="agree" />
                    <span>Remember me</span>
                </label>
                <button className={style.button} type="submit">
                    Login
                </button>
                
            </form>
        </>
    );
};

export default LoginForm;