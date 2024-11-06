import { combineReducers } from "redux";
import { playerReducer } from "./playerReducer";
import {HYDRATE} from 'next-redux-wrapper';
import { trackReducer } from "./trackReducer";
import {currentReducer} from "./currentReducer";
import { current } from "@reduxjs/toolkit";
import { userReducer } from "./userReducer";

const rootReducer = combineReducers({
    player: playerReducer,
    track: trackReducer,
    current: currentReducer,
    user: userReducer
})

export const reducer = (state: any, action: any) => {
    if (action.type === HYDRATE) {
        const nextState = {
            ...state, 
            ...action.payload, 
        }
        if (state.count) nextState.count = state.count 
        return nextState
    } else {
        return rootReducer(state, action)
    }
}


export type RootState = ReturnType<typeof rootReducer>