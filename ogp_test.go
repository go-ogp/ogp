package ogp_test

import (
	"fmt"

	"github.com/go-ogp/ogp"
)

func ExampleWebsite() {
	result := ogp.Website().
		Title("Example").
		URL("http://example.com").
		Image(ogp.Image().URL("http://example.com/social.jpg")).
		HTML()
	fmt.Println(result)
	// Output:
	// <meta property="og:type" content="website">
	// <meta property="og:title" content="Example">
	// <meta property="og:url" content="http://example.com">
	// <meta property="og:image" content="http://example.com/social.jpg">
}

func ExampleArticle() {
	result := ogp.Article().
		Title("How to Train Your Dragons").
		URL("http://example.com/article/how-to-train-your-dragon").
		Image(ogp.Image().URL("http://example.com/image/dragon.jpg")).
		Author(ogp.Profile().URL("http://example.com/profile/dragon-master")).
		HTML()
	fmt.Println(result)
	// Output:
	// <meta property="og:type" content="article">
	// <meta property="og:title" content="How to Train Your Dragons">
	// <meta property="og:url" content="http://example.com/article/how-to-train-your-dragon">
	// <meta property="og:image" content="http://example.com/image/dragon.jpg">
	// <meta property="article:author" content="http://example.com/profile/dragon-master">
}

func ExampleBook() {
	result := ogp.Book().
		Title("Oliver Twist").
		URL("http://example.com/book/oliver-twist").
		Image(ogp.Image().URL("http://example.com/image/cover.jpg")).
		Author(ogp.Profile().URL("http://example.com/profile/charles-dickens")).
		ISBN("9780174325482").
		HTML()
	fmt.Println(result)
	// Output:
	// <meta property="og:type" content="book">
	// <meta property="og:title" content="Oliver Twist">
	// <meta property="og:url" content="http://example.com/book/oliver-twist">
	// <meta property="og:image" content="http://example.com/image/cover.jpg">
	// <meta property="book:isbn" content="9780174325482">
	// <meta property="book:author" content="http://example.com/profile/charles-dickens">
}

func ExampleProfile() {
	result := ogp.Profile().
		Title("John Smith").
		URL("http://jsmith.me").
		Image(ogp.Image().URL("http://jsmith.me/avatar.jpg")).
		FirstName("John").
		LastName("Smith").
		Username("jsmith").
		HTML()
	fmt.Println(result)
	// Output:
	// <meta property="og:type" content="profile">
	// <meta property="og:title" content="John Smith">
	// <meta property="og:url" content="http://jsmith.me">
	// <meta property="og:image" content="http://jsmith.me/avatar.jpg">
	// <meta property="profile:first_name" content="John">
	// <meta property="profile:last_name" content="Smith">
	// <meta property="profile:username" content="jsmith">
}
