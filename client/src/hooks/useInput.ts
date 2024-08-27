import { useState } from "react"

export const useInput = (initialValue: string) => {
    const [value, setValue] = useState(initialValue);

    const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setValue(e.target.value);
    };

    return {value, onChange}
}

export const useAreaInput = (initialValue: string) => {
    const [value, setValue] = useState(initialValue);

    const onChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
        setValue(e.target.value);
    };

    return {value, onChange}
}

export const useFileInput = () => {
    const [value, setValue] = useState<File>();

    const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setValue(e.target.files![0]);
    };

    return {value, onChange}
}