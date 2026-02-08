

type Food struct {
	name   string
	rating int
}

type FoodHeap []Food

func (h FoodHeap) Len() int           { return len(h) }
func (h FoodHeap) Less(i, j int) bool {
	if h[i].rating != h[j].rating {
		return h[i].rating > h[j].rating
	}
	return h[i].name < h[j].name
}
func (h FoodHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *FoodHeap) Push(x interface{}) {
	*h = append(*h, x.(Food))
}
func (h *FoodHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type FoodRatings struct {
	foodToRating   map[string]int
	foodToCuisine  map[string]string
	cuisineToFoods map[string]*FoodHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToRating:   make(map[string]int),
		foodToCuisine:  make(map[string]string),
		cuisineToFoods: make(map[string]*FoodHeap),
	}

	n := len(foods)
	for i := 0; i < n; i++ {
		food := foods[i]
		cuisine := cuisines[i]
		rating := ratings[i]

		fr.foodToRating[food] = rating
		fr.foodToCuisine[food] = cuisine

		if _, exists := fr.cuisineToFoods[cuisine]; !exists {
			fr.cuisineToFoods[cuisine] = &FoodHeap{}
			heap.Init(fr.cuisineToFoods[cuisine])
		}

		heap.Push(fr.cuisineToFoods[cuisine], Food{name: food, rating: rating})
	}

	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.foodToRating[food] = newRating
	cuisine := this.foodToCuisine[food]
	heap.Push(this.cuisineToFoods[cuisine], Food{name: food, rating: newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	h := this.cuisineToFoods[cuisine]
	for h.Len() > 0 {
		top := (*h)[0]
		if this.foodToRating[top.name] == top.rating {
			return top.name
		}
		heap.Pop(h)
	}
	return ""
}