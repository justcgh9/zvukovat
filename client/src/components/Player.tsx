'use client'
import React from "react";
import {Pause, PlayArrow, VolumeUp } from "@mui/icons-material";
import IconButton from '@mui/material/IconButton';
import Grid from '@mui/material/Grid';
import styles from '../styles/Player.module.scss'
import { ITrack } from "@/types/track";
import TrackProgress from "./TrackProgress";


const Player: React.FC = () => {
    const track: ITrack = {
        _id: "66b21aa988a74ad4e0727e5a",
        name: "Avada",
        artist: "Linkin Park",
        text: "Here we go for the hundredth time, hand grenade pins in every line, throw them up and let something shine, going out of my fucking mind.",
        listens: 0,
        picture: "picture/e4c95df0-c09d-4e74-a498-55b81371419f.jpg",
        audio: "audio/52500e5e-2317-4426-8d5c-ca20e7d4d8fc.mp3",
        comments: []
    }
    const active = false
    return (
        <div className={styles.player}>
            <IconButton onClick={e => e.stopPropagation()}>
                {active
                ? <Pause/>
                : <PlayArrow/>
                }
            </IconButton>
            <Grid container direction="column" style={{width:200, margin: '0 20px'}}>
                    <div>{track.name}</div>
                    <div style={{fontSize: 12, color: 'gray'}}>{track.artist}</div>
            </Grid>
            <TrackProgress left={0} right={100} onChange={() => {}}/>
            <VolumeUp style={{marginLeft: 'auto'}}/>
            <TrackProgress left={0} right={100} onChange={() => {}}/>
        </div>
    )
};

export default Player;