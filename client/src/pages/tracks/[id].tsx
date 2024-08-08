import RouterButton from "@/components/RouterButton";
import { ITrack } from "@/types/track";
import React from "react"
import Grid from "@mui/material/Grid"
import Button from "@mui/material/Button"
import TextField from "@mui/material/TextField"

const ViewTrack = () => {
    const track: ITrack = {
        _id: "66b21aa988a74ad4e0727e5a",
        name: "Avada",
        artist: "Linkin Park",
        text: "Here we go for the hundredth time, hand grenade pins in every line, throw them up and let something shine, going out of my fucking mind.",
        listens: 0,
        picture: "picture/e4c95df0-c09d-4e74-a498-55b81371419f.jpg",
        audio: "audio/52500e5e-2317-4426-8d5c-ca20e7d4d8fc.mp3",
        comments: []
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
                    label="Username"
                    fullWidth
                />
                <TextField
                    label="Comment"
                    fullWidth
                    multiline
                    rows={4}
                />
                <Button>Leave Comment</Button>
            </Grid>
            <div>
                {track.comments.map(comment =>
                    <div>
                        <div>Author - {comment.username}</div>
                        <div>Comment - {comment.text}</div>
                    </div>
                )}
            </div>
        </div>
    )
}

export default ViewTrack;
