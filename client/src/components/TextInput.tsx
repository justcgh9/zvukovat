import { TextInputProps } from '@/types/create';
import styles from '../styles/TextInput.module.scss';


export default function TextInput({label, onChange, value, isStretch=false, background=false, type='text'}: TextInputProps){
    return (
    <div className={`${styles.input_container} ${(isStretch ? styles.full_width : '')}`}>
        <input
          type={type}
          id={`textinput-${label}`}
          className={styles.input}
          name={`textinput-${label}`}
          value={value}
          onChange={onChange}
          aria-labelledby="label-textinput"
          autoComplete='off'
        />
        <label className={styles.label} htmlFor={`textinput-${label}`} id="label-textinput">
          <div className={`${styles.text} ${(background ? styles.dark : '')}`}>{label}</div>
        </label>
      </div>
    );
}