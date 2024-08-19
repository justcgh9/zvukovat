import styles from '../styles/User.module.scss';
import Image from 'next/image';
import UserIcon from '../assets/user.svg';
import { UserProps } from '../types/user';
import SignUpIcon from '../assets/uil_sign-in-alt.svg';

export default function User({name}: UserProps){
    return (<div className={styles.user_container}>
        {name && <>
            <Image src={UserIcon} alt='user'/>
            <h4 className={styles.username}>{name}</h4>
        </>}
        {!name && <>
            <button className={styles.filled_button}>Sign in</button>
            <button className={styles.button_with_icon}>
                Sign up
                <Image src={SignUpIcon} alt='sign up'/>
            </button>
        </>}
    </div>);
}