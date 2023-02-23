package domain

type Pizza struct {
	ID       string `json:"id" bson:"_id"`
	ImageUrl string `json:"imageUrl" bson:"imageUrl"`
	Title    string `json:"title" bson:"title"`
	Types    []int  `json:"types" bson:"types"`
	Sizes    []int  `json:"sizes" bson:"sizes"`
	Price    int    `json:"price" bson:"price"`
	Rating   int    `json:"rating" bson:"rating"`
}
