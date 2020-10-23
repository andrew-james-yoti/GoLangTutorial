package joke

// import (
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// var _ = Describe("Jokes", func() {
// 	Context("Get Jokes", func() {
// 		When("a list of jokes is required", func() {
// 			It("should get a list of jokes", func() {
// 				jokeArr := GetJokes()
// 				Expect(len(jokeArr)).To(Equal(17))
// 			})
// 		})
// 	})
// })
import "testing"

func TestGetJokes(t *testing.T) {
	jokeArr := GetJokes()
	if len(jokeArr) != 7 {
		t.Errorf("Joke array length %d", len(jokeArr))
	}
}

func TestSetJokeLike(t *testing.T) {
	SetJokeLike(1)
	jokeArr := GetJokes()
	if jokeArr[0].Likes != 1 {
		t.Errorf("Did not set likes on joke id 1")
	}
}
