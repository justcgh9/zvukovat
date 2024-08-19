import RouterButton from "@/components/RouterButton";
import { IComment, ITrack } from "@/types/track";
import React, { useState } from "react"
import Grid from "@mui/material/Grid"
import Button from "@mui/material/Button"
import TextField from "@mui/material/TextField"
import { GetServerSideProps } from "next";
import axios from "axios";
import { useInput } from "@/hooks/useInput";



const ViewTrack = ({serverTrack}: {serverTrack: ITrack}) => {
    const [track, setTrack] = useState(serverTrack)
    const [comments, setComments] = useState<IComment[]>([])
    const username = useInput('')
    const text = useInput('')

    const addComment = async () => {
        try {
            const response = await axios.post('http://localhost:8080/tracks/' + track.id + '/comment', {
                username: username.value,
                text: text.value
            })
            setTrack({...track, comments: [...track.comments, response.data]})
            fetchComments()
        } catch (e) {
            console.log(e)
        }


    }

    const fetchComments = async () => {
        try {
            const response = await axios.get("http://localhost:8080/tracks/" + track.id + "/comment")
            setComments(response.data)
        } catch (e) {
            console.log(e)
        }
    }



    return (
        <div>
            <RouterButton to="/tracks">
                Back to the list
            </RouterButton>
            <Grid container style={{margin: '20px 0'}}>
                <img src={'http://localhost:8080/files/' + track.picture} alt={track.name + ' picture'} width={200} height={200}/>
                <div style={{margin: '20px 0'}}>
                    <h1> Track name - {track.name} </h1>
                    <h1> Track author - {track.artist} </h1>
                    <h1> Number of plays - {track.listens} </h1>
                </div>
            </Grid>
            <h1>Lyrics</h1>
            <p>{track.text}</p>
            <h1>Comments</h1>
            <Grid container>
                <TextField
                    {...username}
                    label="Username"
                    fullWidth
                />
                <TextField
                    {...text}
                    label="Comment"
                    fullWidth
                    multiline
                    rows={4}
                />
                <Button onClick={addComment}>Leave Comment</Button>
            </Grid>
            <div>
                {comments?.map(comment =>
                    <div key={comment.id}>
                        <div>Автор - {comment.username}</div>
                        <div>Комментарий - {comment.text}</div>
                    </div>
                )}
            </div>
        </div>
    )
}

export default ViewTrack;

export const getServerSideProps: GetServerSideProps = async ({params}) => {
    
    const response = await axios.get('http://localhost:8080/tracks/' + params?.id)
    
    return {
        props: {
            serverTrack: response.data
        }
    }
}
