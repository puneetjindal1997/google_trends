package main

import (
	"context"
	"log"
	"reflect"

	"github.com/groovili/gogtrends"
	"github.com/pkg/errors"
)

const (
	locUS  = "IN"
	catAll = "all"
	langEn = "EN"
)

func main() {
	ctx := context.Background()

	log.Println("Daily trending searches:")
	explore, err := gogtrends.Explore(ctx, &gogtrends.ExploreRequest{
		ComparisonItems: []*gogtrends.ComparisonItem{
			{
				Keyword: "GO",
				Geo:     locUS,
				Time:    "today 12-m",
			},
		},
		Category: 31,
		Property: "",
	}, langEn)
	// dailySearches, err := gogtrends.Daily(ctx, langEn, locUS)
	handleError(err, "Failed to get daily searches")

	ref := reflect.ValueOf(explore)
	if ref.Kind() != reflect.Slice {
		log.Fatal("Not able to processed")
	}
	for i := 0; i < ref.Len(); i++ {
		log.Println(ref.Index(i).Interface())
	}
	// for _, val := range dailySearches {
	// 	fmt.Println(val.Title, "Traffic:- ", val.FormattedTraffic)
	// 	for _, articles := range val.Articles {
	// 		fmt.Println("Articles", articles)
	// 	}
	// }
	// printItems(dailySearches)
}

func handleError(err error, errMsg string) {
	if err != nil {
		log.Fatal(errors.Wrap(err, errMsg))
	}
}
