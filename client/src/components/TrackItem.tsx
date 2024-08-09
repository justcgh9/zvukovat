import { ITrack } from "@/types/track";
import Card from "@mui/material/Card"
import Grid from "@mui/material/Grid"
import styles from "../styles/TrackItem.module.scss"
import { Delete, Pause, PlayArrow } from "@mui/icons-material";
import IconButton from '@mui/material/IconButton'
import { useRouter } from "next/navigation";
import React, { useActionState } from "react";
import { useActions } from "@/hooks/useActions";

interface TrackItemProps {
    track: ITrack
    active?: boolean
}

const TrackItem: React.FC<TrackItemProps> = ({track, active=false} : TrackItemProps) => {
    const router = useRouter()
    const {playTrack, pauseTrack, setActiveTrack} = useActions()

    const play = (e: React.MouseEvent) => {
        e.stopPropagation()
        setActiveTrack(track)
        pauseTrack()
    }

    return (
        <div>
            <Card className={styles.track} onClick={() => router.push('/tracks/' + track.id)}>
                <IconButton onClick={play}>
                    {active
                    ? <Pause/>
                    : <PlayArrow/>
                    }
                </IconButton>
                <img width={70} height={70} src={'http://localhost:8080/files/' + track.picture} alt={track.name + 'picture'}/>
                <Grid container direction="column" style={{width:200, margin: '0 20px'}}>
                    <div>{track.name}</div>
                    <div style={{fontSize: 12, color: 'gray'}}>{track.artist}</div>
                </Grid>
                {active && <div>02:42 / 3:22</div>}
                <IconButton onClick={e => e.stopPropagation()} style={{marginLeft: 'auto'}}>
                    <Delete/>
                </IconButton>                    
            </Card>
        </div>
    )
} 

export default TrackItem;