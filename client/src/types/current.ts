export interface CurrentTrack {
    currentId: string;
}

interface SetCurrentAction {
    type: CurrentActionTypes.SET_CURRENT;
    payload: string;
}

interface ResetCurrentAction {
    type: CurrentActionTypes.RESET_CURRENT;
}

export type CurrentAction = 
    SetCurrentAction | ResetCurrentAction;


export enum CurrentActionTypes {
    SET_CURRENT = "SET_CURRENT",
    RESET_CURRENT = "RESET_CURRENT"
}