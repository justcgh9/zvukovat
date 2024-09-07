import { TextInputProps } from '@/types/create';
import styles from '../styles/PasswordInput.module.scss';
import { useState } from 'react';


export default function PasswordInput({label, onChange, value, isStretch=false, background=false}: TextInputProps){
    const [isHidden, setIsHidden] = useState<boolean>(true);

    return (
        <>
        <div className={`${styles.input_container} ${(isStretch ? styles.full_width : '')}`}>
            <input
            type={isHidden? "password" : "text"}
            id="textinput"
            className={styles.input}
            name="textinput"
            value={value}
            onChange={onChange}
            aria-labelledby="label-textinput"
            autoComplete='off'
            data-1p-ignore data-lpignore="true"
            />
            <label className={styles.label} htmlFor="textinput" id="label-textinput">
            <div className={`${styles.text} ${(background ? styles.dark : '')}`}>{label}</div>
            </label>
        </div>
        <p className={styles.hide_text} onClick={()=>{setIsHidden((prev)=>!prev)}}>{isHidden ? "Show password": "Hide password"}</p>
        </>
    );
}