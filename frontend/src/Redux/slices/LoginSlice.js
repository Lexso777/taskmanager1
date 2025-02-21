import { createAsyncThunk, createSlice,  } from "@reduxjs/toolkit";
import axios from "axios"


export const fetchLogin = createAsyncThunk(
    'auth/login',
    async( userData , {rejectWithValue}) => {
        try{
            const response = await axios.post(`http://localhost:8080/login`, userData);
            console.log(response.data);

            return response.data;
        }catch(error){

            return rejectWithValue(error.response?.data?.message || "Неверный логин или пароль");
        }
    }
)

const initialState = {
    status : 'idle',
    email : null,
    error : null,
}


const LoginSlice = createSlice({
    name : 'Login',
    initialState,
    reducers : {},
    extraReducers : (buider) => {
        buider.addCase(fetchLogin.pending, (state) => {
            state.status = 'loading';
        });
        buider.addCase(fetchLogin.fulfilled, (state, action) => {
            state.status = 'resolve';
            state.email = action.payload.email;
            state.error = '';
        });
        buider.addCase(fetchLogin.rejected, (state, action) => {
            state.status = 'reject';
            state.error = action.payload;
        });
      }
})


export default LoginSlice.reducer