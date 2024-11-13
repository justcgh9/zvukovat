export interface UserProps {
    name: string;
}

export interface User {
    _id: string;
    username: string;
	email: string;
	password: string;
	isActivated: boolean;
	activationLink: string;
	favouriteTracks: string;
}

export interface AuthResp{
    tokens: {
        accessToken: string;
        refreshToken: string;
    };
	user: User;
}

interface SetUserAction {
    type: UserActionTypes.SET_USER;
    payload: User | undefined;
}

interface ResetUserAction {
    type: UserActionTypes.RESET_USER;
}

export type UserAction = 
    SetUserAction | ResetUserAction;


export enum UserActionTypes {
    SET_USER = "SET_USER",
    RESET_USER = "RESET_USER"
}