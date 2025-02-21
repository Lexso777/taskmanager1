import React, { useState } from 'react';
import { useSelector } from 'react-redux';
import style from './MainPage.module.css'


const MainPage = () => {

  const { email } = useSelector((state) => state.LoginSlice);

  const [task, setTask] = useState([]);
  const [newTask, setNewTask] = useState({ title: "", text: "", date: "" });


  const [toggleForm, setToggleForm] = useState(false);
  const [toggleMenu, setToggleMenu] = useState(false);
  const [toggleText, setToggleText] = useState(false);

  const toggle = () => {
    setToggleForm(!toggleForm);
  }
  const toggle1 = () => {
    setToggleMenu(!toggleMenu);
  }
  const toggle2 = () => {
    
  }




  const addTask = () => {
    if (newTask.title.trim() && newTask.text.trim()) {

      const newTaskWithDate = {
        ...newTask,
        date: new Date().toISOString(), 
      };

      setTask([...task, newTaskWithDate]);
      setNewTask({ title: "", text: "", date: "" });
      setToggleForm(false)
    }
  }


  return (
    <div className={style.body}>
      <div className={style.nav__container}>
        Гость : {email}
      </div>
      <div className={style.tasks__container}>
        <div className={style.task__container}>
          <div className={style.task__container__title}>
            <div className={style.task__container__title__name}>Coming soon</div>
            <div class={style.menu__container}>
              <button onClick={toggle1} class={style.menu__button}>...</button>
              {toggleMenu && (
                <div className={style.menu}>
                  <div href="#">Опция 1</div>
                  <div href="#">Опция 2</div>
                  <div href="#">Опция 3</div>
                </div>)}
            </div>
          </div>
          {task.map((task, index) => (
            <div key={index} className={style.task__card}>
              <div className={style.task__card__name}>{task.title}</div>
              <div className={style.task__card__description}>
                <div className={style.task__card__date}>
                  {new Date(task.date).toLocaleDateString("en-GB", {
                    day: "2-digit",
                    month: "short",
                  })}
                </div>
                <div className={style.task__card__text}>{task.text}</div>
              </div>

            </div>
          ))}
          <div className={style.button__add__task}>
            <button onClick={toggle}>ADD TASK</button>
          </div>
        </div>
        {toggleForm && (
          <div>
            <div className={style.overlay} onClick={() => setToggleForm(false)}></div>
            <form className={style.createTaskForm}>
              <div className={style.taskForm}>
                <div className={style.task__form_name}>
                  <div className={style.task__form__title}>Название создаваемой задачи </div>
                  <input
                    type="text"
                    className={style.task__form__input}
                    placeholder='Название'
                    value={newTask.title}
                    onChange={(e) => setNewTask({ ...newTask, title: e.target.value })}
                  />
                </div>
                <div className={style.task__card__text}>
                  <div className={style.task__form__title}> Описание</div>
                  <textarea
                    type="text"
                    className={style.task__form__textarea}
                    value={newTask.text}
                    onChange={(e) => setNewTask({ ...newTask, text: e.target.value })} />
                </div>
                <button onClick={addTask}>Добавить</button>
              </div>
            </form>
          </div>)}
      </div>
    </div>
  );
};

export default MainPage;