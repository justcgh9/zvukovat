import { PlayerAction, PlayerActionTypes, TracksOrder } from "@/types/player"
import { TrackResp } from "@/types/track"

export const playTrack = (): PlayerAction => {
    return {type: PlayerActionTypes.PLAY}
}

export const pauseTrack = (): PlayerAction => {
    return {type: PlayerActionTypes.PAUSE}
}

export const setDuration = (payload: number): PlayerAction => {
    return {type: PlayerActionTypes.SET_DURATION, payload}
}

export const setOrder = (payload: TracksOrder): PlayerAction => {
    return {type: PlayerActionTypes.SET_ORDER, payload}
}

export const setCurrentTime = (payload: number): PlayerAction => {
    return {type: PlayerActionTypes.SET_CURRENT_TIME, payload}
}

export const setActiveTrack = (payload: TrackResp): PlayerAction => {
    return {type: PlayerActionTypes.SET_ACTIVE, payload}
}