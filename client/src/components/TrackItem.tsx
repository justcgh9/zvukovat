import styles from '../styles/TrackItem.module.scss';
import IconButton from './IconButton';
import PlayIcon from '../assets/play.svg';
import PauseIcon from '../assets/pause.svg';
import Image from 'next/image';
import FavouriteIcon from '../assets/favourite_icon.svg';
import FavouriteChosen from '../assets/favourite_chosen.svg';
import { TrackProps, TrackResp } from '@/types/track';
import { MouseEvent, useEffect, useState } from 'react';
import { useActions } from '@/hooks/useActions';
import {resetCurrentTrack, setCurrentTrack} from "../store/action-creators/current";
import { useTypedSelector } from '@/hooks/useTypedSelector';
import { useDispatch } from 'react-redux';
import { UnknownAction } from 'redux';
import DefaultCover from '../assets/default_cover.svg';



export default function TrackItem({track, isFavourite = false}: TrackProps){

    const [audio, setAudio] = useState<HTMLAudioElement>();
    const [duration, setDuration] = useState<number>(0);


    useEffect(()=>{
        let audio = new Audio('http://localhost:8080/files/' + track.audio);
        audio.onloadedmetadata = () => {
            setDuration(Math.ceil(audio.duration));
        }

    }, []);
    
    const {playTrack, pauseTrack, setActiveTrack} = useActions();
    const { currentId } = useTypedSelector(state => state.current);
    const dispatch = useDispatch();

    const play = (e: React.MouseEvent) => {
        e.stopPropagation()
        if (currentId === track.id){
            dispatch(resetCurrentTrack() as UnknownAction);
            dispatch(pauseTrack()  as UnknownAction);
        } else {
            dispatch(playTrack() as UnknownAction);
            dispatch(setCurrentTrack(track.id) as UnknownAction);
            dispatch(setActiveTrack(track) as UnknownAction);
            
        }
    }



    return (<div className={styles.track_container}>
        <div className={styles.track_left_cont}>
            <IconButton className={styles.playing_btn} icon={(currentId !== track.id) ? PlayIcon : PauseIcon} onClick={play} alt={(currentId !== track.id) ? "Play" : "Pause"}/>
            <Image className={styles.cover_img} width={64} height={64} src={track.picture ? 'http://localhost:8080/files/' + track.picture : DefaultCover} alt={track.name}/>
            <div className={styles.track_info}>
                <h4 className={styles.track_title}>{track.name}</h4>
                <p className={styles.track_singer}>{track.artist}</p>
            </div>
        </div>
        <div className={styles.track_right_cont}>
            <p className={styles.track_time}>{`${Math.floor(duration/60)}:${('0' + (duration%60)).slice(-2)}`}</p>
            <IconButton icon={!isFavourite ? FavouriteIcon : FavouriteChosen} alt={!isFavourite ? 'not favourite' : 'favourite'} onClick={() => isFavourite = !isFavourite}/>
        </div>
   </div>);
}