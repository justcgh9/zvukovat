import { CurrentAction, CurrentActionTypes } from "../../types/current";

export const setCurrentTrack = (payload: string) : CurrentAction => {
    return {type: CurrentActionTypes.SET_CURRENT, payload};
}

export const resetCurrentTrack = () : CurrentAction => {
    return {type: CurrentActionTypes.RESET_CURRENT};
}
