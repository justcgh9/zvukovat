import Image from 'next/image';
import styles from '../styles/Search.module.scss';
import SearchIcon from '../assets/uil_search.svg';
import { ChangeEvent, ChangeEventHandler } from 'react';

interface SearchProps {
    onChange: (e: ChangeEvent<HTMLInputElement>) => Promise<void>;
    value: string; 
}

export default function Search( {onChange, value}: SearchProps){
    return (<div className={styles.search_container}>
        <Image src={SearchIcon} alt='search'/>
        <input className={styles.search_input} type='text' placeholder='Search' value={value} onChange={onChange}/>
    </div>);
}