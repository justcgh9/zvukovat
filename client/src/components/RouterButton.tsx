'use client'
import Button from '@mui/material/Button'
import { useRouter } from 'next/navigation';
import React from 'react';

interface RouterButtonProps {
    children: React.ReactNode;
    to: string;
}

const RouterButton : React.FC<RouterButtonProps> 
    = ({
        children,
        to
    }: RouterButtonProps) => {

    const router = useRouter()
    return (
        <>
            <Button onClick={() => router.push(to)}>
                {children}
            </Button>
        </>
    )

}

export default RouterButton;