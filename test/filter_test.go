package main

import (
	"testing"

	"bytes"
	"io/ioutil"

	"time"

	"github.com/bootjp/go_twitter_bot_for_nicopedia/item"

	"os"

	"log"

	"github.com/mmcdole/gofeed"
)

func TestFilterSuccess(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	f, err := ioutil.ReadFile("../testdata/bootjp.me/feed.xml")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Get actual value
	fp := gofeed.NewParser()
	feed, err := fp.Parse(bytes.NewReader(f))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if len(item.FilterDate(feed.Items, time.Now())) != 0 {
		t.Fatalf("count miss match.")
		os.Exit(1)
	}

	f, err = ioutil.ReadFile("../testdata/bootjp.me/feed.xml")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	feed, err = fp.Parse(bytes.NewReader(f))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", "2018-04-01 09:33:06", loc)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if len(item.FilterDate(feed.Items, tm)) != 0 {
		t.Fatalf("item count miss match. expect 0 got %d.", len(item.FilterDate(feed.Items, tm)))
		os.Exit(1)
	}

	f, err = ioutil.ReadFile("../testdata/bootjp.me/feed.xml")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	feed, err = fp.Parse(bytes.NewReader(f))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	tm, err = time.ParseInLocation("2006-01-02 15:04:05", "2018-04-01 09:33:05", loc)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if len(item.FilterDate(feed.Items, tm)) != 1 {
		t.Fatalf("item count miss match. expect 0 got %d.", len(item.FilterDate(feed.Items, tm)))
		os.Exit(1)
	}

	f, err = ioutil.ReadFile("../testdata/bootjp.me/feed.xml")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	feed, err = fp.Parse(bytes.NewReader(f))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	tm, err = time.ParseInLocation("2006-01-02 15:04:05", "2018-04-01 00:33:04", loc)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if len(item.FilterDate(feed.Items, tm)) != 1 {
		t.Fatalf("item count miss match. expect 0 got %d.", len(item.FilterDate(feed.Items, tm)))
		os.Exit(1)
	}
}
