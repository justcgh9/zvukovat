import TextInput from '@/components/TextInput';
import styles from '../../styles/Create.module.scss';
import { useInput, useAreaInput, useFileInput } from '@/hooks/useInput';
import TextAreaInput from '@/components/TextAreaInput';
import FileUploader from '@/components/FileUploader';
import ImageIcon from '../../assets/image_icon.svg';
import AudioFileIcon from '../../assets/audio_file_icon.svg';
import { MouseEvent, useState } from 'react';
import axios from 'axios';
import router from 'next/router';
import api from '@/http';

export default function CreateTrack(){
    const title = useInput('');
    const performer = useInput('');
    const lyrics = useAreaInput('');
    const cover = useFileInput();
    const audioFile = useFileInput();

    const [errors, setErrors] = useState<string[]>([]);
    
    function checkInput(){
        let newErrors = [];
        if(title.value.trim() === ""){
            newErrors[newErrors.length] = "Track title is required";
            console.log('empty title');
        }

        if(performer.value.trim() === ""){
            newErrors[newErrors.length] = "Performer name is required";
            console.log('empty performer', errors);
        }

        if(audioFile.value === undefined){
            newErrors[newErrors.length] = "Audio file is required to complete the song";
            console.log('empty audio', errors);
        }

        if(audioFile.value && !(audioFile.value.name.split('.').slice(-1)[0].toLowerCase() === 'mp3' || audioFile.value.name.split('.').slice(-1)[0].toLowerCase() === 'wav' || audioFile.value.name.split('.').slice(-1)[0].toLowerCase() === 'ogg')){
            newErrors[newErrors.length] = "Invalid audio file format";
            console.log('invalid audiofile format ->',  audioFile.value.name.split('.').slice(-1)[0].toLowerCase());
        }

        if(cover.value && !(cover.value.name.split('.').slice(-1)[0].toLowerCase() === 'jpg' || cover.value.name.split('.').slice(-1)[0].toLowerCase() === 'png' || cover.value.name.split('.').slice(-1)[0].toLowerCase() === 'jpeg')){
            newErrors[newErrors.length] = "Invalid cover image format";
            console.log('invalid image format ->', cover.value.name);
        } 

        setErrors([...newErrors]);
        return newErrors.length === 0;
    }

    async function handleSubmit(event: MouseEvent<HTMLButtonElement, Event>) {
        event.preventDefault();
        if (checkInput()){
            const formData = new FormData();
            formData.append('name', title.value);
            formData.append('artist', performer.value);
            formData.append('text', lyrics.value);
            formData.append('picture', cover.value!);
            formData.append('audio', audioFile.value!);
            api.post('/tracks', formData)
                .then(resp => router.push('/track'))
                .catch(e => console.log(e));
        }
        
    }


    return (<section id={styles.create_page_container}>
        <h2 className={styles.main_title}>Upload track</h2>
        <form className={styles.create_form}>
            <label className={styles.create_form_label}>Provide track details</label>
            
            <div className={styles.form_container}>
                <div className={styles.form_text_inputs}>
                    <div className={styles.form_container_text}>
                        <TextInput label='Title' value={title.value} onChange={title.onChange}/>
                        <TextInput label='Performer' value={performer.value} onChange={performer.onChange}/>
                    </div>
                    <TextAreaInput label='Lyrics' value={lyrics.value} onChange={lyrics.onChange}/>
                
                </div>
                <div className={styles.form_file_inputs}>
                    <FileUploader label='Choose track cover' icon={ImageIcon} formats='JPG, PNG, JPEG' onChange={cover.onChange} accept='image/*'/>
                    <FileUploader label='Choose audio file' icon={AudioFileIcon} formats='MP3, WAV, OGG' onChange={audioFile.onChange} accept='audio/*'/>
                </div>
            </div>
            { errors.length !== 0 && <div className={styles.errors}>
                <h4 className={styles.errors_heading}>Please Correct the Following</h4>
                {errors.map((error) => <p key={error} className={styles.error}>{error}</p>)}
            </div>
            }
            <button type='submit' onClick={handleSubmit} className={styles.save_btn}>Save</button>
        </form>
    </section>);
}
