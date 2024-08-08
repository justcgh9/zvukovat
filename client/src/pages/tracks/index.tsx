import Grid from '@mui/material/Grid'
import Card from '@mui/material/Card'
import Box from '@mui/material/Box'
import RouterButton from '@/components/RouterButton'
import { ITrack } from '@/types/track'
import TrackList from '@/components/TrackList'

export default function Tracks() {
    const tracks: ITrack[] = [
        {
            _id: "66b21aa988a74ad4e0727e5a",
            name: "Avada",
            artist: "Linkin Park",
            text: "Here we go for the hundredth time, hand grenade pins in every line, throw them up and let something shine, going out of my fucking mind.",
            listens: 0,
            picture: "picture/e4c95df0-c09d-4e74-a498-55b81371419f.jpg",
            audio: "http://localhost:8080/files/audio/52500e5e-2317-4426-8d5c-ca20e7d4d8fc.mp3",
            comments: []
        }
    ]
    return (
        <>
            <Grid container justifyContent='center'>
                <Card style={{width: 900}}>
                    <Box p={3}>
                        <Grid container justifyContent='space-between'>
                            <h1> Tracks List </h1>
                            <RouterButton to='/tracks/create'> Upload Track </RouterButton>
                        </Grid>
                    </Box>
                    <TrackList tracks={tracks}/>
                </Card>
            </Grid>
        </>
    )
}