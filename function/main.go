package main

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
)

const (
	ADD_POST                 string = "addPost"
	ALL_POSTS                       = "allPosts"
	GET_POST                        = "getPost"
	ADD_POST_ERROR_WITH_DATA        = "addPostErrorWithData"
	RELATED_POSTS                   = "relatedPosts"
	TEST                            = "test"
)

type ReqEvent struct {
	Field     string            `json:"field"`
	Arguments map[string]string `json:"arguments"`
	Source    map[string]string `json:"source"`
}

type Post struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Url     string `json:"url"`
	Ups     int    `json:"ups"`
	Downs   int    `json:"downs"`
	// RelatePost []Post `json:"relatePost"`
}

var postsMap = map[string]Post{
	"1": {
		Id:      "1",
		Title:   "Au Bonheur des Dames",
		Author:  "Emile Zola",
		Content: "Émile Zola (1840-1902), né à Paris, est un écrivain, journaliste et homme public français, considéré comme le chef de file du naturalisme. C'est l'un des romanciers français les plus universellement populaires, l'un des plus publiés et traduits au monde, le plus adapté au cinéma et à la télévision.",
		Url:     "https://amazon.com/Bonheur-Dames-French-Emile-Zola/dp/1482759314",
		Ups:     100,
		Downs:   100,
	},
	"2": {
		Id:      "2",
		Title:   "LOTR",
		Author:  "JRRT",
		Content: "dummycontent",
		Url:     "https://www.amazon.com/J.-R.-R.-Tolkien/e/B000ARC6KA",
		Ups:     100,
		Downs:   10,
	},
	"3": {
		Id:      "3",
		Title:   "A story of Deep Dish",
		Author:  "Sir Willingham III",
		Content: "",
		Url:     "",
		Ups:     0,
		Downs:   0,
	},
	"4": {
		Id:      "4",
		Title:   "L'Insoutenable Légèreté de l'être",
		Author:  "Milan Kundera",
		Content: "n this story of irreconcilable loves and infidelities, Milan Kundera addresses himself to the nature of 20th century Being. The novel encompasses the extremes of comedy and tragedy, and embraces, it seems, all aspects of human existence.",
		Url:     "https://www.amazon.com/LInsoutenable-Legerete-lEtre-Milan-Kundera/dp/207038165X/",
		Ups:     1000,
		Downs:   0,
	},
	"5": {
		Id:      "5",
		Title:   "Fahreineit 451",
		Author:  "Ray Bradbury",
		Content: "Ray Bradbury’s internationally acclaimed novel Fahrenheit 451 is a masterwork of twentieth-century literature set in a bleak, dystopian future. Guy Montag is a fireman. In his world, where television rules and literature is on the brink of extinction, firemen start fires rather than put them out. ",
		Url:     "https://www.amazon.com/Fahrenheit-451-Ray-Bradbury/dp/1451673310",
		Ups:     50,
		Downs:   0,
	},
}

var relatedPosts = map[string][]Post{
	"1": {postsMap["4"]},
	"2": {postsMap["3"], postsMap["5"]},
	"3": {postsMap["2"], postsMap["1"]},
	"4": {postsMap["2"], postsMap["1"]},
	"5": {},
}

func handleRequest(ctx context.Context, req ReqEvent) (interface{}, error) {
	switch req.Field {
	case TEST:
		return req.Arguments, nil
	case ADD_POST:
		return req.Arguments, nil
	case ALL_POSTS:
		posts := make([]Post, 0)
		for _, p := range postsMap {
			posts = append(posts, p)
		}
		return posts, nil
	case GET_POST:
		return postsMap[req.Arguments["id"]], nil
	case ADD_POST_ERROR_WITH_DATA:
		return nil, nil
	case RELATED_POSTS:
		return relatedPosts[req.Source["id"]], nil
	default:
		return nil, errors.New("Unknown field, unable to resolve" + req.Field)
	}
}

func main() {
	lambda.Start(handleRequest)
}
