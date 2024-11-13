import { UserAction, UserActionTypes, User } from "../../types/user";


const initialState : User = {} as User;

export const userReducer = (state = initialState, action: UserAction) : User => {
    switch (action.type) {
        case UserActionTypes.SET_USER:
            if (action.payload) {
                return {
                    ...state,
                    _id: action.payload._id,
                    username: action.payload.username,
                    email: action.payload.email,
                    password: action.payload.password,
                    isActivated: action.payload.isActivated,
                    activationLink: action.payload.activationLink,
                    favouriteTracks: action.payload.favouriteTracks
                };
            }
            return state;
        case UserActionTypes.RESET_USER:
            return {...state, ...initialState};
        default:
            return state;
    }
}
