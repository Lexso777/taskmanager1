import React, { useState } from "react";
import style from "./Form.module.css";

const RegisterForm = () => {
  const [email, setEmail] = useState("");
  const [password1, setPassword1] = useState("");
  const [password2, setPassword2] = useState("");
  const [match, setMatch] = useState(null);
  const [error, setError] = useState("");

  const checkInput = async (event) => {
    event.preventDefault();

    if (password1 !== password2) {
      setMatch(false);
      return;
    }

    setMatch(true);
    setError(""); // Очистка ошибок перед запросом

    // Отправка запроса на сервер
    try {
      const response = await fetch("http://localhost:8080/create", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password: password1 }),
      });

      const result = await response.json();

      if (response.ok) {
        alert("Регистрация успешна!");
      } else {
        setError(result.message || "Ошибка регистрации");
      }
    } catch (err) {
      setError("Ошибка соединения с сервером");
    }
  };

  return (
    <>
      <h2 className={style.h2}>Register</h2>
      <form className={style.form} onSubmit={checkInput}>
        <input
          className={style.input}
          type="email"
          name="email"
          autoComplete="email"
          placeholder="Введите адрес эл.почты"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          className={style.input}
          type="password"
          name="password"
          autoComplete="new-password"
          placeholder="Придумайте пароль"
          value={password1}
          onChange={(e) => setPassword1(e.target.value)}
          required
        />
        <input
          className={style.input}
          type="password"
          name="confirm-password"
          autoComplete="new-password"
          placeholder="Повторите пароль"
          value={password2}
          onChange={(e) => setPassword2(e.target.value)}
          required
        />
        {match === false && <div className={style.error}>Пароли не совпадают</div>}
        {error && <div className={style.error}>{error}</div>}
        <button className={style.button} type="submit">
          Register
        </button>
      </form>
    </>
  );
};

export default RegisterForm;