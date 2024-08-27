import { TextInputProps } from '@/types/create';
import styles from '../styles/TextInput.module.scss';


export default function TextInput({label, onChange, value}: TextInputProps){
    return (
    <div className={styles.input_container}>
        <input
          type="text"
          id="textinput"
          className={styles.input}
          name="textinput"
          value={value}
          onChange={onChange}
          aria-labelledby="label-textinput"
        />
        <label className={styles.label} htmlFor="textinput" id="label-textinput">
          <div className={styles.text}>{label}</div>
        </label>
      </div>
    );
}