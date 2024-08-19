import { CurrentAction, CurrentActionTypes, CurrentTrack } from "../../types/current";


const initialState : CurrentTrack = {
    currentId: "",
}

export const currentReducer = (state = initialState, action: CurrentAction) : CurrentTrack => {
    switch (action.type) {
        case CurrentActionTypes.SET_CURRENT:
            return {...state, currentId: action.payload};
        case CurrentActionTypes.RESET_CURRENT:
            return {...state, currentId: ""};
        default:
            return state;
    }
}
