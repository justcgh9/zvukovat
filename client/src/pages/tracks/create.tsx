import StepDisplayer from "@/components/StepDisplayer";
import Grid from '@mui/material/Grid'
import Button from '@mui/material/Button'
import { useState } from "react";
import  TextField  from "@mui/material/TextField";
import FileUploader from "@/components/FileUploader";
import { useInput } from "@/hooks/useInput";
import axios from "axios";
import { useRouter } from "next/navigation";


export default function CreateTrack() {
    const router = useRouter()
    const [currentStep, setCurrentStep] = useState(0)
    const [thumbnail, setThumbnail] = useState(null)
    const [audio, setAudio] = useState(null)

    const name = useInput('')
    const artist = useInput('')
    const text = useInput('')

    const next = () => {
        if (currentStep !== 2) {
            setCurrentStep(prev => prev + 1)
            return
        }

        if (thumbnail == null || audio == null) {
            return
        }

        const formData = new FormData()
        formData.append('name', name.value)
        formData.append('artist', artist.value)
        formData.append('text', text.value)
        formData.append('picture', thumbnail)
        formData.append('audio', audio)
        axios.post('http://localhost:8080/tracks/upload', formData)
            .then(resp => router.push('/tracks'))
            .catch(e => console.log(e))
    }

    const back = () => {
        setCurrentStep(prev => prev - 1)
    }

    return (

        <div>
            <StepDisplayer currentStep={currentStep}>
                {currentStep === 0 &&
                <Grid container direction={'column'} style={{padding: 20}}>
                    <TextField
                        {...name}
                        style={{marginTop: 15}}
                        label={"Track name"}
                    />
                    <TextField
                        {...artist}
                        style={{marginTop: 15}}
                        label={"Author"}
                    />
                    <TextField
                        {...text}
                        style={{marginTop: 15}}
                        label={"Lyrics"}
                        multiline
                        rows={3}
                    />
                </Grid>
                }
                {currentStep === 1 &&
                    <FileUploader setFile={setThumbnail} accept="image/*">
                        <Button>Upload Thumbnail</Button>
                    </FileUploader>
                }
                {currentStep === 2 && 
                    <FileUploader setFile={setAudio} accept="audio/*">
                        <Button>Upload Audio</Button>
                    </FileUploader>

                }
            </StepDisplayer>
            <Grid container justifyContent="space-between">
                <Button onClick={back} disabled={currentStep <= 0}>Back</Button>
                <Button onClick={next} disabled={currentStep >= 3}>Next</Button>
            </Grid>
        </div>
    )
}