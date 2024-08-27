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

        if(audioFile.value && audioFile.value.name.split('.').slice(-1)[0].toLowerCase() !== ('mp3' || 'wav' || 'ogg')){
            newErrors[newErrors.length] = "Invalid cover image format";
            console.log('invalid audiofile format');
        }

        if(cover.value && cover.value.name.split('.').slice(-1)[0].toLowerCase() !== ('jpg' || 'png' || 'jpeg')){
            newErrors[newErrors.length] = "Invalid audio file format";
            console.log('invalid image format');
        } 

        setErrors([...newErrors]);
    }

    async function handleSubmit(event: MouseEvent<HTMLButtonElement, Event>) {
        event.preventDefault();
        console.log(title.value, performer.value, lyrics.value, cover.value, audioFile.value);
        console.log("befor", errors);
        checkInput();
        console.log("after", errors);
        if (errors.length === 0){
            const formData = new FormData();
            formData.append('name', title.value);
            formData.append('artist', performer.value);
            formData.append('text', lyrics.value);
            formData.append('picture', cover.value!);
            formData.append('audio', audioFile.value!);
            axios.post('http://localhost:8080/tracks/upload', formData)
                .then(resp => router.push('/track'))
                .catch(e => console.log(e));
        }
        
    }


    return (<section id={styles.create_page_container}>
        <h2 className={styles.main_title}>Upload track</h2>
        <form className={styles.create_form}>
            <label className={styles.create_form_label}>Provide track details</label>
            <div className={styles.form_text_inputs}>
                <div className={styles.form_container_text}>
                    <TextInput label='Title' value={title.value} onChange={title.onChange}/>
                    <TextInput label='Performer' value={performer.value} onChange={performer.onChange}/>
                </div>
                <TextAreaInput label='Lyrics' value={lyrics.value} onChange={lyrics.onChange}/>
            
            </div>
            <div className={styles.form_container}>
                <div className={styles.form_file_inputs}>
                    <FileUploader label='Choose track cover' icon={ImageIcon} formats='JPG, PNG, JPEG' onChange={cover.onChange} accept='image/*'/>
                    <FileUploader label='Choose audio file' icon={AudioFileIcon} formats='MP3, WAV, OGG' onChange={audioFile.onChange} accept='audio/*'/>
                </div>
                { errors.length !== 0 && <div className={styles.errors}>
                    <h4 className={styles.errors_heading}>Please Correct the Following</h4>
                    {errors.map((error) => <p className={styles.error}>{error}</p>)}
                </div>
                }
                <button type='submit' onClick={handleSubmit} className={styles.save_btn}>Save</button>
            </div>
        </form>
    </section>);
}

// async function submitQuestion(event: FormEvent) {
//     event.preventDefault();
//     if (submitQuestionCheck() && errorMark === '' && errorYear === '') {
//       if (task !== null && task.id !== null && valueMark !== undefined) {
//         const newTask: TaskResponse = {
//           id: task?.id,
//           content: inputValueText,
//           document_id: task.document_id,
//           marks: valueMark,
//           page: task.page,
//           topic: topicIndex,
//           verified: null,
//           year: selectedYear,
//         };
//         const createResponse: AxiosResponse<TaskResponse> = await axios.patch(
//           `https://chartreuse-binghamite1373.my-vm.work/task/${task.id}`,
//           newTask,
//         );
//         if (createResponse.status === 200) {
//           console.log('submited');
//           setShowChart(false);
//           afterSave();
//         }
//       } else {
//         const newTask: TaskCreateRequest = {
//           content: inputValueText,
//           document_id: null,
//           marks: valueMark,
//           page: null,
//           topic: topicIndex,
//           verified: null,
//           year: selectedYear,
//         };
//         const createResponse: AxiosResponse<TaskResponse> = await axios.post(
//           `https://chartreuse-binghamite1373.my-vm.work/task/`,
//           newTask,
//         );
//         if (createResponse.status === 200) {
//           console.log('submited');
//           afterSave();
//         }
//       }
//     } else {
//       console.log('ne submit');
//     }
//   }


// function submitQuestionCheck() {
//     let newErrors = 0;
//     if (
//       inputValueText.trim().length === 0 ||
//       inputValueText === 'Enter the task...'
//     ) {
//       setErrorText('Task cannot be empty.');
//       newErrors++;
//     }
//     if (topic === '' || topicIndex === -1) {
//       setErrorTopic('Topic is not chosen.');
//       newErrors++;
//     }
//     if (valueMark === undefined) {
//       setErrorMark('Mark cannot be empty.');
//     }
//     if (valueMark !== undefined && valueMark > 0 && valueMark < 21) {
//       setErrorMark('');
//     }
//     return newErrors === 0;
//   }