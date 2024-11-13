import Link from 'next/link';
import styles from '../../styles/SignUpPage.module.scss';
import TextInput from '@/components/TextInput';
import PasswordInput from '@/components/PasswordInput';
import { useInput } from '@/hooks/useInput';
import { useState } from 'react';
import { useDispatch } from 'react-redux';
import { NextThunkDispatch } from '@/store';
import { registerUser } from '@/store/action-creators/user';
import router from 'next/router';

export default function SignUp(){
    const firstName = useInput('');
    const lastName = useInput('');
    const email = useInput('');
    const password = useInput('');
    const [errors, setErrors] = useState<string[]>([]);
    const dispatch = useDispatch() as NextThunkDispatch;
    
    async function handleSubmit(){
        if(checkInput()){
            console.log(`${firstName.value} ${lastName.value}`, email.value, password.value)
            dispatch(await registerUser(`${firstName.value} ${lastName.value}`, email.value, password.value)).then((a) => {a ? setErrors(a): router.push("/track")});
        }
    }

    function checkInput(){
        let newErrors = [];
        if(firstName.value.trim() === ""){
            newErrors[newErrors.length] = "First name is required";
            console.log('empty name');
        }

        if(lastName.value.trim() === ""){
            newErrors[newErrors.length] = "Last name is required";
            console.log('empty last name');
        }

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

    

    return (<section id={styles.sign_up}>
        <div className={styles.create_account_cont}>
            <div className={styles.heading_cont}>
                <h2 className={styles.heading}>Create account</h2>
                <p className={styles.paragraph}>Aready have an account?<Link className={styles.link} href='/signin'>Sign in</Link></p>
            </div>
            <form className={styles.form}>
                <div className={styles.text_inputs}>
                    <TextInput label='First name' value={firstName.value} onChange={firstName.onChange} background={true}></TextInput>
                    <TextInput label='Last name' value={lastName.value} onChange={lastName.onChange} background={true}></TextInput>
                </div>
                <TextInput label='Email' value={email.value} onChange={email.onChange} background={true} isStretch={true} type='email'></TextInput>
                <PasswordInput label='Password' value={password.value} onChange={password.onChange} background={true} isStretch={true}/>
                <button type='button' className={styles.submit_button} onClick={handleSubmit}>Create account</button>
                { errors.length !== 0 &&
                    <p className={styles.error}>{errors.toLocaleString().replaceAll(',', ', ')}</p>
                }
            </form>
        </div>
    </section>);
}