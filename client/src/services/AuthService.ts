import api from "@/http";
import { AuthResp } from "@/types/user";
import { AxiosResponse } from "axios";


export default class AuthService{
    static async login(email: string, password: string): Promise<AxiosResponse<AuthResp>>{
        return api.post<AuthResp>('/login', {"email": email, "password": password}, {withCredentials: true});
    }

    static async registration(username: string, email: string, password: string): Promise<AxiosResponse<AuthResp>>{
        let response = api.post<AuthResp>('/registration', {"username": username, "email": email, "password": password}, {withCredentials: true});
        return response;
    }

    static async logout(): Promise<void>{
        return api.post('/logout', {withCredentials: true});
    }
}