import { ITrack } from "@/types/track"
import Grid from "@mui/material/Grid"
import Box from "@mui/material/Box"
import TrackItem from "./TrackItem"

interface TrackListProps {
    tracks: ITrack[];
}

const TrackList: React.FC<TrackListProps> = ({tracks}: TrackListProps) => {
    return(
        <Grid container direction="column">
            <Box p={2}>
                {tracks.map(track =>
                    <TrackItem
                        key={track._id}
                        track={track}
                    />
                )}
            </Box>
        </Grid>
    )
}

export default TrackList