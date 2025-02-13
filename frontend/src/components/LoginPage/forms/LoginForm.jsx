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

            <h3>Add Product</h3>
            <form method="POST" action="http://localhost:8080/create" className={style.test}>
                <div>
                    <label>Model</label>
                    <input type="text" name="model" />
                </div>
                <div>
                    <label>Company</label>
                    <input type="text" name="company" />
                </div>
                <div>
                    <label>Price</label>
                    <input type="number" name="price" />
                </div>
                <input type="submit" value="Send" />
            </form>
            {/* </form> */}
        </>
    );
};

export default LoginForm;