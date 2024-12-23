import styles from '../styles/User.module.scss';
import Image from 'next/image';
import UserIcon from '../assets/user.svg';
import SignUpIcon from '../assets/uil_sign-in-alt.svg';
import Link from 'next/link';
import { useTypedSelector } from '@/hooks/useTypedSelector';
import IconButton from './IconButton';
import Logout from '../assets/logout.svg';
import { logoutUser } from '@/store/action-creators/user';
import { useDispatch } from 'react-redux';
import { NextThunkDispatch } from '@/store';
import { useEffect } from 'react';

export default function User(){
    const user = useTypedSelector(state => state.user);
    let username;
    (Object.keys(user).length === 0) ? username = "": username = user.username;
    const dispatch = useDispatch() as NextThunkDispatch;

    useEffect(()=>{console.log(user);},[user])

    async function handleLogout() {
        await dispatch(await logoutUser());
    }
    return (<div className={styles.user_container}>
        {username && <>
            <div className={styles.user_inner}>
                <Image src={UserIcon} alt='user'/>
                <h4 className={styles.username}>{username}</h4>
            </div>
            <IconButton icon={Logout} alt='logout' onClick={handleLogout} title='Logout'/>
        </>}
        {!username && 
            <div className={styles.user_inner}>
            <Link href="/signin" className={styles.link_container}>
                <button className={styles.filled_button}>Sign in</button>
            </Link>
            <Link href="/signup" className={styles.link_container}>
                <button className={styles.button_with_icon}>
                    Sign up
                    <Image src={SignUpIcon} alt='sign up'/>
                </button>
            </Link>
            </div>}
    </div>);
}