import { AuthResp } from '@/types/user';
import axios, { AxiosResponse, InternalAxiosRequestConfig } from 'axios';

export const API_URL = "http://localhost:8080/";

const api = axios.create({
    withCredentials: true,
    baseURL: API_URL,
});


api.interceptors.request.use(function (config: InternalAxiosRequestConfig){
    const token = localStorage.getItem('token');
    config.headers.Authorization =  token ? `Bearer ${token}` : '';
    return config;
});

api.interceptors.response.use((config: AxiosResponse) => {
    return config;
}, async (error) => {
    const originalRequest = error.config;
    if (error.response.status == 401 && error.config && !error.config._retry) {
        originalRequest._retry = true;
        try {
            const response = await axios.get<AuthResp>(`${API_URL}/refresh`, {withCredentials: true})
            localStorage.setItem('token', response.data.tokens.accessToken);
            return api.request(originalRequest);
        } catch (e) {
            console.log('not auth')
        }
    }
    throw error;
});
export default api;