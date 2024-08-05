package models

type Track struct {
    Id       string   `bson:"_id,omitempty" json:"id"`
    Name     string   `bson:"name" json:"name"`
    Artist   string   `bson:"artist" json:"artist"`
    Text     string   `bson:"text" json:"text"`
    Listens  int      `bson:"listens" json:"listens"`
    Picture  string   `bson:"picture" json:"picture"`
    Audio    string   `bson:"audio" json:"audio"`
    Comments []string `bson:"comments" json:"comments"`
}

type TrackPaginationParams struct {
    Count   int
    Offset  int
}


