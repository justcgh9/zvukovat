export interface TextInputProps {
    label: string;
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
    value: string;
}

export interface TextAreaInputProps {
    label: string;
    onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
    value: string;
}

export interface FileUploaderProps{
    label: string;
    icon: string;
    formats: string;
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
    accept: string;
}