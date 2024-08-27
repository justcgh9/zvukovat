import { FileUploaderProps } from '@/types/create';
import styles from '../styles/FileUploader.module.scss';
import Image from 'next/image';


export default function FileUploader({label, icon, formats, onChange, accept} : FileUploaderProps){
    return (<div className={styles.file_uploader}>
        <div className={styles.heading}>
            
            <label className={styles.label} htmlFor='fileinput'>{label}</label>
            <Image src={icon} alt='label'/>
        </div>
        <div className={styles.inner_container}>
            
            <div className={styles.input_container}>
                <input type='file' accept={accept} id='fileinput' className={styles.input} onChange={onChange}></input>
                <p className={styles.formats}>Formats: {formats}</p>
            </div>
        </div>
    </div>);
}