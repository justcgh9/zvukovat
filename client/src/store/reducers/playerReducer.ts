import { PlayerAction, PlayerActionTypes, PlayerState, TracksOrder } from "@/types/player"

const initialState: PlayerState = {
    currentTime: 0,
    duration: 200,
    active: null,
    order: TracksOrder.MIX,
    isPlaying: false
}

export const playerReducer = (state = initialState, action: PlayerAction): PlayerState => {
    switch (action.type) {
        case PlayerActionTypes.PLAY:
            return{...state, isPlaying: true}
        case PlayerActionTypes.PAUSE:
            return{...state, isPlaying: false}
        case PlayerActionTypes.SET_ACTIVE:
            return{...state, active: action.payload, duration: 200, currentTime: 0}
        case PlayerActionTypes.SET_CURRENT_TIME:
            return{...state, currentTime: action.payload}
        case PlayerActionTypes.SET_DURATION:
            return{...state, duration: action.payload}
        case PlayerActionTypes.SET_ORDER:
            return{...state, order: action.payload}
        default:
            return state

    }
}