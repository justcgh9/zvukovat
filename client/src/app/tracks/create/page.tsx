'use client'
import StepDisplayer from "@/components/StepDisplayer";
import Grid from '@mui/material/Grid'
import Button from '@mui/material/Button'
import { useState } from "react";


export default function CreateTrack() {
    const [currentStep, setCurrentStep] = useState(0)

    const next = () => {
        setCurrentStep(prev => prev + 1)
    }

    const back = () => {
        setCurrentStep(prev => prev - 1)
    }

    return (

        <div>
            <StepDisplayer currentStep={currentStep}>
                <h1>Uploading track</h1>
            </StepDisplayer>
            <Grid container justifyContent="space-between">
                <Button onClick={back} disabled={currentStep <= 0}>Back</Button>
                <Button onClick={next} disabled={currentStep >= 3}>Next</Button>
            </Grid>
        </div>
    )
}