package models

type Playlist struct {
    Id      string   `bson:"_id,omitempty" json:"id"`
    Tracks  []string `bson:"tracks" json:"tracks"`
    Name    string   `bson:"name"   json:"name"`
    Owner  string   `bson:"owner" json:"owner"`
    Picture string   `bson:"picture" json:"picture"`
    IsPrivate bool  `bson:"isPrivate" json:"isPrivate"`
}

type AddToPlaylistDTO struct {
    TrackId string  `json:"track_id"`
}
