import styles from '../styles/TrackItem.module.scss';
import IconButton from './IconButton';
import PlayIcon from '../assets/play.svg';
import PauseIcon from '../assets/pause.svg';
import Image from 'next/image';
import FavouriteIcon from '../assets/favourite_icon.svg';
import FavouriteChosen from '../assets/favourite_chosen.svg';
import { TrackProps } from '@/types/track';
import { MouseEvent, useEffect } from 'react';
import { useActions } from '@/hooks/useActions';
import {resetCurrentTrack, setCurrentTrack} from "../store/action-creators/current";
import { useTypedSelector } from '@/hooks/useTypedSelector';
import { useDispatch } from 'react-redux';
import { UnknownAction } from 'redux';

export default function TrackItem({track, isFavourite = false, duration}: TrackProps){
    let durationInMinutes = `${Math.floor(duration/60)}:${('0' + (duration%60)).slice(-2)}`;
    const {playTrack, pauseTrack, setActiveTrack} = useActions();
    const { currentId } = useTypedSelector(state => state.current);
    const dispatch = useDispatch();
    console.log(currentId);

    const play = (e: React.MouseEvent) => {
        if (currentId === track.id){
            dispatch(resetCurrentTrack() as UnknownAction);
        } else {
            dispatch(setCurrentTrack(track.id) as UnknownAction);
        }

        e.stopPropagation()
        setActiveTrack(track)
        pauseTrack()
    }



    return (<div className={styles.track_container}>
        <div className={styles.track_left_cont}>
            <IconButton className={styles.playing_btn} icon={(currentId !== track.id) ? PlayIcon : PauseIcon} onClick={play} alt={(currentId !== track.id) ? "Play" : "Pause"}/>
            <Image className={styles.cover_img} width={64} height={64} src={'http://localhost:8080/files/' + track.picture} alt={track.name}/>
            <div className={styles.track_info}>
                <h4 className={styles.track_title}>{track.name}</h4>
                <p className={styles.track_singer}>{track.artist}</p>
            </div>
        </div>
        <div className={styles.track_right_cont}>
            <p className={styles.track_time}>{durationInMinutes}</p>
            <IconButton icon={!isFavourite ? FavouriteIcon : FavouriteChosen} alt={!isFavourite ? 'not favourite' : 'favourite'} onClick={() => isFavourite = !isFavourite}/>
        </div>
   </div>);
}