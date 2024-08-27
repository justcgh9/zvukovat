import { TextAreaInputProps } from '@/types/create';
import styles from '../styles/TextArea.module.scss';


export default function TextAreaInput({label, onChange, value}: TextAreaInputProps){
    return (
    <div className={styles.input_textarea_container}>
        <textarea
          id="textarea"
          className={`${styles.input_textarea} ${value==="" ? "" : styles.has_value}`}
          name="textarea"
          value={value}
          onChange={onChange}
          aria-labelledby="label-textarea"
        />
        <label className={styles.label_textarea} htmlFor="textarea" id="label-textarea">
          <div className={styles.text_textarea}>{label}</div>
        </label>
      </div>
    );
}