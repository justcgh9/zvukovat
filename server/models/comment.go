package models

type Comment struct {
    Id          string  `bson:"_id,omitempty" json:"id"`
    Track_id    string  `bson:"track_id" json:"track_id"`
    Username    string  `bson:"username" json:"username"`
    Text        string  `bson:"text" json:"text"`

}
