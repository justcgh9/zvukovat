import Grid from '@mui/material/Grid'
import Card from '@mui/material/Card'
import Box from '@mui/material/Box'
import TextField from '@mui/material/TextField'
import RouterButton from '@/components/RouterButton'
import { ITrack } from '@/types/track'
import TrackList from '@/components/TrackList'
import { useTypedSelector } from '@/hooks/useTypedSelector'
import { NextThunkDispatch, wrapper } from '@/store'
import { fetchTracks, searchTracks } from '@/store/action-creators/track'
import { GetServerSideProps } from 'next/types'
import { useState } from 'react'
import { useDispatch } from 'react-redux'

export default function Tracks() {
    const {tracks, error} = useTypedSelector(state => state.track)
    const [query, setQuery] = useState<string>('')
    const [timer, setTimer] = useState<any>(null)
    const dispatch = useDispatch() as NextThunkDispatch

    const search = async (e: React.ChangeEvent<HTMLInputElement>) => {
        setQuery(e.target.value)
        if (timer) {
            clearTimeout(timer)
        }
        setTimer(
            setTimeout(async () => {
                await dispatch(await searchTracks(e.target.value))
            }, 500)
        )
        // console.log(tracks)
    }
    if (error) {
        return <h1>{error}</h1>
    }
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
                    <TextField
                        fullWidth
                        value={query}
                        label='Search for songs'
                        onChange={search}
                    />
                    <TrackList tracks={tracks}/>
                </Card>
            </Grid>
        </>
    )
}


export const getServerSideProps = wrapper.getServerSideProps(store => async (context) =>{
    const dispatch = store.dispatch as NextThunkDispatch
    await dispatch(await fetchTracks())
    return {props: {id: null}}; 
  });