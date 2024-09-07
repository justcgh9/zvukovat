import styles from '../styles/User.module.scss';
import Image from 'next/image';
import UserIcon from '../assets/user.svg';
import { UserProps } from '../types/user';
import SignUpIcon from '../assets/uil_sign-in-alt.svg';
import Link from 'next/link';
import { useTypedSelector } from '@/hooks/useTypedSelector';

export default function User({name}: UserProps){
    const user = useTypedSelector(state => state.user);
    let username;
    (Object.keys(user).length === 0) ? username = "": username = user.username;
    console.log(username)
    return (<div className={styles.user_container}>
        {username && <>
            <Image src={UserIcon} alt='user'/>
            <h4 className={styles.username}>{username}</h4>
        </>}
        {!username && <>
            <Link href="/signin" className={styles.link_container}>
                <button className={styles.filled_button}>Sign in</button>
            </Link>
            <Link href="/signup" className={styles.link_container}>
                <button className={styles.button_with_icon}>
                    Sign up
                    <Image src={SignUpIcon} alt='sign up'/>
                </button>
            </Link>
        </>}
    </div>);
}