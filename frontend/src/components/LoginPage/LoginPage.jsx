import { useEffect, useState } from 'react';
import style from './LoginPage.module.css'
import LoginForm from './forms/LoginForm';
import RegisterForm from './forms/RegisterForm';



const LoginPage = () => {

    const [activeTab, setActiveTab] = useState('login'); 
    const [message, setMessage] = useState('');
    const [error, setError] = useState('');

    useEffect(() => {
        fetch('http://localhost:8080/')
            .then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                return response.text();
            })
            .then(data => {
                setMessage(data);
                setError(''); // Сбрасываем ошибку, если запрос успешен
            })
            .catch(error => {
                console.error('Ошибка:', error);
                setError('Не удалось загрузить данные с сервера.');
            });
    }, []);

    return (
        <div className={style.loginPage__container}>
            <div>{message}</div>
            <div className={style.login__module__window}>
                <div className={style.module__window__buttons}>
                    <button 
                    onClick={()=> setActiveTab('login')}
                    className={activeTab === 'login' ? style.buttonAvtive : style.buttons}
                    >
                        Login
                    </button>
                    <button
                     onClick={()=> setActiveTab('register')}
                    className={activeTab === 'register' ? style.buttonAvtive : style.buttons}
                    >
                        Registration
                    </button>
                </div>
                <div className={style.login}>
                    {activeTab === 'login' && <LoginForm />}
                    {activeTab === 'register' && <RegisterForm/>}
                </div>
            </div>
        </div>
    );
};

export default LoginPage;