import { TrackResp } from "./track";


export enum TracksOrder {
    MIX = "MIX",
    REPEAT = "REPEAT",
    REPEAT_ONE  = "REPEAT_ONE"
}

export interface PlayerState {
    active: null | TrackResp;
    order: TracksOrder;
    duration: number;
    currentTime: number;
    isPlaying: boolean;
}

export enum PlayerActionTypes {
    PLAY = "PLAY",
    PAUSE = "PAUSE",
    SET_ACTIVE = "SET_ACTIVE",
    SET_DURATION = "SET_DURATION",
    SET_CURRENT_TIME = "SET_CURRENT_TIME",
    SET_ORDER = "SET_ORDER"
}

interface PlayAction {
    type: PlayerActionTypes.PLAY;
}

interface PauseAction {
    type: PlayerActionTypes.PAUSE;
}

interface SetActiveAction {
    type: PlayerActionTypes.SET_ACTIVE;
    payload: TrackResp;
}

interface SetDurationAction {
    type: PlayerActionTypes.SET_DURATION;
    payload: number;
}

interface SetCurrentTimeAction {
    type: PlayerActionTypes.SET_CURRENT_TIME;
    payload: number;
}

interface SetOrderAction {
    type: PlayerActionTypes.SET_ORDER;
    payload: TracksOrder;
}

export type PlayerAction = 
    PlayAction
    |   PauseAction
    |   SetActiveAction
    |   SetCurrentTimeAction
    |   SetDurationAction
    |   SetOrderAction



