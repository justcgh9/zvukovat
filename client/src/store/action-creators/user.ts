import AuthService from "@/services/AuthService";
import { AuthResp, User, UserAction, UserActionTypes } from "../../types/user";
import axios, { AxiosError } from "axios";
import { Dispatch } from "redux";
import { API_URL } from "@/http";



export const loginUser = async (email: string, password:string) => {
    return async (dispatch: Dispatch<UserAction>) => {
        try {
            const response = await AuthService.login(email, password);
            if(response.status === 200){
                localStorage.setItem('token', response.data.accessToken);
                dispatch(setUser(response.data.user));
                console.log("user-> ", response.data.user);
                return null;
            } 
        } catch (e: unknown) {
            if (e instanceof AxiosError){
                return e.response?.data.split("\n")[0];
            }
        } 
    }
}

export const registerUser = async (username: string, email: string, password:string) => {
    return async (dispatch: Dispatch<UserAction>) => {
        try {
            const response = await AuthService.registration(username, email, password);
            if (response.status === 200){
                localStorage.setItem('token', response.data.accessToken);
                setUser(response.data.user);
            }
        }  catch (e: unknown) {
            if (e instanceof AxiosError){
                return e.response?.data;
            }
        } 
    }
}

export const logoutUser = async () => {
    return async (dispatch: Dispatch<UserAction>) => {
        try {
            const response = await AuthService.logout();
            localStorage.removeItem('token');
            resetUser();
        }  catch (e: unknown) {
            if (e instanceof AxiosError){
                console.log(e.response?.data?.message);
            }
        }
    }
}

export const checkAuth = async () => {
    return async (dispatch: Dispatch<UserAction>) => {
        try {
            const response = await axios.post<AuthResp>(`${API_URL}refresh`, {withCredentials: true});
            localStorage.setItem('token', response.data.accessToken);
            setUser(response.data.user);
            console.log(response.data.user);
        } catch (e: unknown) {
            if (e instanceof AxiosError){
                console.log(e.response?.data?.message);
            }
        }
    }
}



const setUser = (payload: User) : UserAction => {
    console.log("userr", payload);
    return {type: UserActionTypes.SET_USER, payload};
}

const resetUser = () : UserAction => {
    return {type: UserActionTypes.RESET_USER};
}

