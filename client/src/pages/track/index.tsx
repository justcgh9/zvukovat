import TrackItem from '@/components/TrackItem';
import styles from './page.module.scss';
import Search from '@/components/Search';
import axios, { AxiosResponse } from 'axios';
import { TrackResp, TracksResp } from '@/types/track';
import { useTypedSelector } from '@/hooks/useTypedSelector';

export default function Tracks({tracksData}:TracksResp) {
    console.log(tracksData);


    

    return <section className={styles.track_page_cont}>
        <h1 className={styles.main_title}>Tracks</h1>
        <Search/>
        <ul className={styles.tracklist}>
            {tracksData.map((track : TrackResp) => 
                <li>
                <TrackItem track={track} isFavourite={true} duration={183}/>
                </li>
            )}
        </ul>
        
    </section>;
}

export async function getServerSideProps() {
    // Fetch data from external API
    const response : AxiosResponse<TracksResp>= await axios.get("http://localhost:8080/tracks");
    const tracksData: TracksResp =  response.data;
   
    // Pass data to the page via props
    return { props: { tracksData } }
  }