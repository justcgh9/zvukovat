package models

type Album struct {
    Id      string   `bson:"_id,omitempty" json:"id"`
    Tracks  []string `bson:"tracks" json:"tracks"`
    Artist  string   `bson:"artist" json:"artist"`
    Picture string   `bson:"picture" json:"picture"`
}
