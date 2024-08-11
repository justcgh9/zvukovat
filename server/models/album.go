package models

type Album struct {
    Id      string   `bson:"_id,omitempty" json:"id"`
    Tracks  []string `bson:"tracks" json:"tracks"`
    Name    string   `bson:"name"   json:"name"`
    Artist  string   `bson:"artist" json:"artist"`
    Picture string   `bson:"picture" json:"picture"`
}

type AddToAlbumDTO struct {
    TrackId string  `json:"track_id"`
}
