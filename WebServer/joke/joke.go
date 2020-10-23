package joke

// Joke struct contains blah
type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

// GetJokes returns an array of the jokes
func GetJokes() []Joke {
	return jokes
}

// SetJokeLike increments the likes of a particular joke Item
func SetJokeLike(jokeid int) {
	for i := 0; i < len(jokes); i++ {
		if jokes[i].ID == jokeid {
			jokes[i].Likes++
		}
	}
}
