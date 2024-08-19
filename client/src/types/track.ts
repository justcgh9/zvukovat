import { StaticImageData } from "next/image";

export interface TrackProps{
    track: TrackResp;
    isFavourite: boolean;
    duration: number; // will be past from back
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









export interface IComment {
    id: string;
    track_id: string;
    username: string;
    text: string;
}

export interface ITrack {
    id: string;
    name: string;
    artist: string;
    text: string;
    listens: number;
    picture: string;
    audio: string;
    comments: IComment[]
}

export interface TrackState {
    tracks: ITrack[]
    error: string
}

export enum TrackActionTypes {
    FETCH_TRACKS = 'FETCH_TRACKS',
    FETCH_TRACKS_ERROR = 'FETCH_TRACKS_ERROR',
}

interface FetchTracksAction {
    type: TrackActionTypes.FETCH_TRACKS
    payload: ITrack[]
}

interface FetchTracksErrorAction {
    type: TrackActionTypes.FETCH_TRACKS_ERROR
    payload: string
}

export type TrackAction = 
    FetchTracksAction
    |   FetchTracksErrorAction