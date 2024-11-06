import styles from '../styles/Player.module.scss';
import IconButton from './IconButton';
import PreviousIcon from '../assets/previous.svg';
import NextIcon from '../assets/next.svg';
import PlayIcon from '../assets/play_circle.svg';
import PauseIcon from '../assets/pause_circle.svg';
import Image from 'next/image';
import FavouriteIcon from '../assets/favourite_icon.svg';
import FavouriteChosen from '../assets/favourite_chosen.svg';
import MixIcon from '../assets/order_mix.svg';
import RepeatIcon from '../assets/order_repeat.svg';
import RepeatOneIcon from '../assets/order_repeat_one.svg';
import Pic from '../assets/image_track_example.png';
import { ChangeEvent, useEffect, useRef, useState } from 'react';
import { useTypedSelector } from '@/hooks/useTypedSelector';
import { useActions } from '@/hooks/useActions';
import { TracksOrder } from '@/types/player';
import { useDispatch } from 'react-redux';
import { resetCurrentTrack, setCurrentTrack } from '@/store/action-creators/current';
import { UnknownAction } from 'redux';


let audio: HTMLAudioElement;

export default function Player(){

    const {isPlaying, active, duration, currentTime, order} = useTypedSelector(state => state.player);
    const {pauseTrack, playTrack, setActiveTrack, setCurrentTime, setDuration, setOrder} = useActions();

    let durationInMinutes = `${Math.floor(duration/60)}:${('0' + (duration%60)).slice(-2)}`;
    const dispatch = useDispatch();
    const [rotationAngle, setRotationAngle] = useState(0);
    // const coverRef = useRef<HTMLImageElement>(null);
    // const animationRef = useRef<number | null>(null);
    

    useEffect(() => {
        if(!audio) {
            audio = new Audio();

        } else {
            if (active) {
                setAudio();
                if(isPlaying){
                    audio.play();
                    // startRotating();
                } else {
                    audio.pause();
                    // stopRotating();
                }
                
            }
        }
    }, [active]);


    useEffect (() =>{
        if(isPlaying){
            audio.play();
            // startRotating();
        } else {
            audio.pause();
            // stopRotating();
        }
    }, [isPlaying]);

    const setAudio = () => {
        if (active) {
            audio.src = 'http://localhost:8080/files/' + active.audio
            audio.volume = 1
            audio.onloadedmetadata = () => {
                setDuration(Math.ceil(audio.duration))
            }
            audio.ontimeupdate = () => {
                setCurrentTime(Math.ceil(audio.currentTime))
                
            }
        }
    }

    const play = () => {
        if (!isPlaying) {
            playTrack();
            if(active){
                dispatch(setCurrentTrack(active.id) as UnknownAction);
            }
        } else {
            pauseTrack();
            dispatch(resetCurrentTrack() as UnknownAction);
        }
    }
   
    // const startRotating = () => {
    //     const startTime = Date.now();
    //     const initialAngle = rotationAngle;
      
    //     const rotate = () => {
    //       const elapsedTime = Date.now() - startTime;
    //       const newAngle = initialAngle + (elapsedTime * 0.06); // 0.06 degrees per ms
    //       setRotationAngle(newAngle % 360);
    //       if (coverRef.current) {
    //         coverRef.current.style.transform = `rotate(${newAngle % 360}deg)`;
    //       }
    //       animationRef.current = requestAnimationFrame(rotate);
    //     };
      
    //     animationRef.current = requestAnimationFrame(rotate);
    //   };
      
    //   const stopRotating = () => {
    //     if (animationRef.current) {
    //       cancelAnimationFrame(animationRef.current);
    //     }
    //   };
      

    const changeCurrentTime = (e: React.ChangeEvent<HTMLInputElement>) => {
        audio.currentTime = Number(e.target.value);
        setCurrentTime(Number(e.target.value));
    }

    return ( active && <div className={styles.player_container}>
        <div className={styles.player_left_container}>
            <div className={styles.track_control_container}>
                <IconButton icon={PreviousIcon} onClick={ () => {}} alt='previous'/>
                <IconButton icon={isPlaying? PauseIcon : PlayIcon} onClick={play} alt='play'/>
                <IconButton icon={NextIcon} onClick={ () => {}} alt='next'/>
            </div>
            <Image className={styles.player_track_cover} src={'http://localhost:8080/files/' + active.picture} alt='track cover' width={64} height={64} /*ref={coverRef}*/
        style={{ transform: `rotate(${rotationAngle}deg)` }}/>
            <div className={styles.track_info}>
                <h4 className={styles.track_title}>{active.name}</h4>
                <p className={styles.track_singer}>{active.artist}</p>
            </div>
        </div>
        <div className={styles.player_right_container}>
            <div className={styles.player_time_container}>
            <p className={styles.track_time}>{`${Math.floor(currentTime/60)}:${('0' + (currentTime%60)).slice(-2)}`}</p>
            <input className={styles.player_time_range} type='range' value={currentTime} min={0} max={duration} onChange={changeCurrentTime}/>
            <p className={styles.track_time}>{durationInMinutes}</p>
            </div>
            <IconButton icon={order === TracksOrder.MIX ? MixIcon : (order === TracksOrder.REPEAT ? RepeatIcon : RepeatOneIcon)} alt={'mix'} onClick={() => {}}/>
            <IconButton icon={FavouriteIcon} alt={'favourite'} onClick={() => {}}/>
            
            {/* <IconButton icon={!isFavourite ? FavouriteIcon : FavouriteChosen} alt={!isFavourite ? 'not favourite' : 'favourite'} onClick={() => isFavourite = !isFavourite}/> */}
        </div>

    </div>);
}