import Link from 'next/link';
import styles from '../styles/NavLink.module.scss';
import { NavigationLink } from '@/types/navigationLink';
import Image from 'next/image';

export default function NavLink({src, text, to} : NavigationLink){
    return (<Link href={to} className={styles.link_container}>
        <Image src={src} alt={text} className={styles.link_icon}></Image>
        <h4 className={styles.link_text}>{text}</h4>
    </Link>);
}