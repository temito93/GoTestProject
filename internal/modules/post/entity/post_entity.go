package entity

type PostEntity struct {
	ID     string `bson:"_id, omitempty"`
	Title  string `bson:"title"`
	UserID string `bson:"user_id"`
}
