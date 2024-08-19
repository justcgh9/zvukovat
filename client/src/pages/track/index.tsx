import TrackItem from '@/components/TrackItem';
import styles from './page.module.scss';
import Search from '@/components/Search';
import axios, { AxiosResponse } from 'axios';
import { TrackResp, TracksResp } from '@/types/track';
import { useTypedSelector } from '@/hooks/useTypedSelector';
import { useState } from 'react';
import { NextThunkDispatch, wrapper } from '@/store';
import { useDispatch } from 'react-redux';
import { fetchTracks, searchTracks } from '@/store/action-creators/track';

export default function Tracks() {

    const {tracks, error} = useTypedSelector(state => state.track)
    const [query, setQuery] = useState<string>('')
    const [timer, setTimer] = useState<any>(null)
    const dispatch = useDispatch() as NextThunkDispatch

    async function search(e: React.ChangeEvent<HTMLInputElement>) {
        setQuery(e.target.value)
        if (timer) {
            clearTimeout(timer)
        }
        setTimer(
            setTimeout(async () => {
                await dispatch(await searchTracks(e.target.value))
            }, 350)
        )
        // console.log(tracks)
    }
    if (error) {
        return <h1>{error}</h1>
    }

    return <section className={styles.track_page_cont}>
        <h1 className={styles.main_title}>Tracks</h1>
        <Search onChange={search} value={query}/>
        <ul className={styles.tracklist}>
            {tracks.map((track : TrackResp) => 
                <li>
                <TrackItem track={track} isFavourite={true} duration={183}/>
                </li>
            )}
        </ul>
        
    </section>;
}

export const getServerSideProps = wrapper.getServerSideProps(store => async (context) =>{
    const dispatch = store.dispatch as NextThunkDispatch
    await dispatch(await fetchTracks())
    return {props: {id: null}}; 
  });