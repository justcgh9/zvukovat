import api from "@/http"
import { TrackAction, TrackActionTypes } from "@/types/track"
import axios from "axios"
import { Dispatch } from "react"

export const fetchTracks = async () => {
    return async (dispatch: Dispatch<TrackAction>) => {
        try {
            const response = await api.get("http://localhost:8080/tracks")
            dispatch({type: TrackActionTypes.FETCH_TRACKS, payload: response.data})
        } catch (e) {
            dispatch({type: TrackActionTypes.FETCH_TRACKS_ERROR, payload: 'Error fetching tracks'})
        }
    }
}

export const searchTracks = async (query: string) => {
    return async (dispatch: Dispatch<TrackAction>) => {
        try {
            const response = await api.get("http://localhost:8080/tracks/search?name=" + query)
            console.log(response.data)
            dispatch({type: TrackActionTypes.FETCH_TRACKS, payload: response.data})
        } catch (e) {
            dispatch({type: TrackActionTypes.FETCH_TRACKS_ERROR, payload: 'Error fetching tracks'})
        }
    }
}