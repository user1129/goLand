package domain

type Pizza struct {
	ID       string `json:"id" bson:"_id"`
	ImageUrl string `json:"imageUrl" bson:"imageUrl"`
	Title    string `json:"title" bson:"title"`
	Types    []int  `json:"types" bson:"types"`
	Sizes    []int  `json:"sizes" bson:"sizes"`
	Price    int    `json:"price" bson:"price"`
	Category int    `json:"category,omitempty" bson:"category"`
	Rating   int    `json:"rating" bson:"rating"`
}

type PizzaFitlerReq struct {
	SortBy   *string `query:"sortBy"`
	OrderBy  *string `query:"orderBy"`
	Category *int    `query:"category"`
}
