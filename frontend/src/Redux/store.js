import { configureStore } from '@reduxjs/toolkit'
import  LoginSlice  from './slices/LoginSlice';

export const store = configureStore({
    reducer : {
        LoginSlice,
    }
})


window.store = store;