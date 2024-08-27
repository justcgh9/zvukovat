import { StaticImageData } from "next/image";

export interface TrackProps{
    track: TrackResp;
    isFavourite: boolean;
    // duration: number; // will be past from back
}

export interface TrackResp {
    artist: string;
    audio: string;
    comments: any; //will be deleted
    id: string;
    listens: number;
    name: string;
    picture: string;
    text: string;
    //duration: number;
}

export interface TracksResp {
    tracksData: TrackResp[];
}

export interface TrackState {
    tracks: TrackResp[]
    error: string
}

export enum TrackActionTypes {
    FETCH_TRACKS = 'FETCH_TRACKS',
    FETCH_TRACKS_ERROR = 'FETCH_TRACKS_ERROR',
}

interface FetchTracksAction {
    type: TrackActionTypes.FETCH_TRACKS
    payload: TrackResp[]
}

interface FetchTracksErrorAction {
    type: TrackActionTypes.FETCH_TRACKS_ERROR
    payload: string
}

export type TrackAction = 
    FetchTracksAction
    |   FetchTracksErrorAction