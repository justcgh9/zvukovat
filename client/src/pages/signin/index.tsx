import Link from 'next/link';
import styles from '../../styles/SignUpPage.module.scss';
import TextInput from '@/components/TextInput';
import PasswordInput from '@/components/PasswordInput';
import { useInput } from '@/hooks/useInput';
import { useState } from 'react';
import { useDispatch } from 'react-redux';
import { NextThunkDispatch } from '@/store';
import { loginUser } from '@/store/action-creators/user';
import { useTypedSelector } from '@/hooks/useTypedSelector';
import { UnknownAction } from 'redux';
import router from 'next/router';

export default function SignIn(){
    const email = useInput('');
    const password = useInput('');
    const [errors, setErrors] = useState<string[]>([]);
    const dispatch = useDispatch() as NextThunkDispatch;

    
    function checkInput(){
        let newErrors = [];

        if(email.value.trim() === ""){
            newErrors[newErrors.length] = "Email is required";
            console.log('empty email', errors);
        }

        if (email.value.trim() !== "" && !(email.value.trim().match(/\w+@\w+\.\w+/))){
            newErrors[newErrors.length] = "Invalid email format";
            console.log('invalid email', errors);
        }

        if(password.value.trim() === ""){
            newErrors[newErrors.length] = "Password is required";
            console.log('empty password', errors);
        }

        setErrors([...newErrors]);
        return newErrors.length === 0;
    }

    async function handleSubmit(){
        if(checkInput()){
            dispatch(await loginUser(email.value, password.value));
            // router.push('/track');
        }
    }

    
    return (<section id={styles.sign_up}>
        <div className={styles.create_account_cont}>
            <div className={styles.heading_cont}>
                <h2 className={styles.heading}>Welcome back</h2>
                <p className={styles.paragraph}>Don't have an account yet?<Link className={styles.link} href='/signup'>Sign up</Link></p>
            </div>
            <form className={styles.form}>
                <TextInput label='Email' value={email.value} onChange={email.onChange} background={true} isStretch={true}></TextInput>
                <PasswordInput label='Password' value={password.value} onChange={password.onChange} background={true} isStretch={true}/>
                <button type='button' className={styles.submit_button} onClick={handleSubmit}>Login</button>
                { errors.length !== 0 &&
                    <p className={styles.error}>{errors.toLocaleString().replaceAll(',', ', ')}</p>
                }
            </form>
        </div>
    </section>);
}