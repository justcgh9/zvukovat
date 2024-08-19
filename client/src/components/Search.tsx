import Image from 'next/image';
import styles from '../styles/Search.module.scss';
import SearchIcon from '../assets/uil_search.svg';

export default function Search(){
    return (<div className={styles.search_container}>
        <Image src={SearchIcon} alt='search'/>
        <input className={styles.search_input} type='text' placeholder='Search' value={""}/>
    </div>);
}