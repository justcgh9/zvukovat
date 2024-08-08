'use client'
import React, { RefObject, useRef } from "react";

interface FileUploaderProps {
    setFile: Function
    accept: string
    children: React.ReactNode
}

const FileUploader: React.FC<FileUploaderProps> = ({setFile, accept, children}) => {
    const ref = useRef<HTMLInputElement>()

    const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files != null) {
            setFile(e.target.files[0])
        }
    }

    return(
        <div onClick={() => ref.current?.click()}>
            <input
                type="file" 
                accept={accept} 
                style={{display: "none"}} 
                ref={ref as RefObject<HTMLInputElement>}
                onChange={onChange}    
            ></input>
            {children}
        </div>
    )
}

export default FileUploader;