import styles from '../styles/Navigation.module.scss';
import NavLink from './NavLink';
import TracksIcon from '../assets/track_icon.svg';
import AlbumsIcon from '../assets/album_icon.svg';
import PlaylistsIcon from '../assets/playlist_icon.svg';
import PerformersIcon from '../assets/performers_icon.svg';
import FavouriteIcon from '../assets/favourite_icon.svg';

export default function Navigation(){
    return (<div className={styles.nav_container}>
        <h3 className={styles.nav_section_name}>My music</h3>
        <nav className={styles.nav}>
            <ul>
                <li>
                    <NavLink src={TracksIcon} text='Tracks' to='/track' />
                </li>
                <li>
                    <NavLink src={AlbumsIcon} text='Albums' to='/albums' />
                </li>
                <li>
                    <NavLink src={PlaylistsIcon} text='Playlists' to='/playlists' />
                </li>
                <li>
                    <NavLink src={PerformersIcon} text='Performers' to='/performers' />
                </li>
                <li>
                    <NavLink src={FavouriteIcon} text='Favourites' to='/favourites' />
                </li>
            </ul>
        
        </nav>
        </div>);
}