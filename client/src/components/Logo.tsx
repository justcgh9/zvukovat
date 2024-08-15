import styles from '../styles/Logo.module.scss';
import Icon from '../assets/icon.svg';
import Image from 'next/image';

export default function Logo(){
    return (<div className={styles.logo_container}>
        <Image id={styles.icon_logo} src={Icon} alt='icon'></Image>
        <h1 id={styles.logo_text}>Zvukovat</h1>
    </div>);
}