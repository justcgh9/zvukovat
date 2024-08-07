import React from "react"
import Container from '@mui/material/Container'
import Stepper from '@mui/material/Stepper'
import Step from '@mui/material/Step'
import StepLabel from '@mui/material/StepLabel'
import Grid from '@mui/material/Grid'
import Card from '@mui/material/Card'

interface StepDisplayerProps{
    currentStep: number
    children: React.ReactNode
}

const steps = ['Provide track information', 'Upload thumbnail', 'Upload audio']

const StepDisplayer: React.FC<StepDisplayerProps> = ({currentStep, children}) => {
    return (
        <>
            <Container>
                <Stepper activeStep={currentStep}>
                    {steps.map((step, index) =>
                        <Step 
                            key={index}
                            completed={currentStep > index}
                        >
                            <StepLabel>{step}</StepLabel>
                        </Step>
                    )}
                </Stepper>
                <Grid container justifyContent='center' style={{margin: '70px', height: 270}}>
                    <Card style={{width: 600}}>
                        {children}
                    </Card>
                </Grid>
            </Container>
        </>
    );
};

export default StepDisplayer