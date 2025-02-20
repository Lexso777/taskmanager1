import { useState } from 'react';
import style from './LoginPage.module.css'
import LoginForm from './forms/LoginForm';
import RegisterForm from './forms/RegisterForm';



const LoginPage = () => {

    const [activeTab, setActiveTab] = useState('login'); 

    return (
        <div className={style.loginPage__container}>
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