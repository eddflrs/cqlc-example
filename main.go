package main

import (
    "github.com/relops/cqlc/cqlc"
    "log"
    "tux21b.org/v1/gocql"
    "twitter"
)

var TWEETS = twitter.TweetsTableDef()

func main() {

    host := "127.0.0.1"
    keyspace := "twitter_example"

    cluster := gocql.NewCluster(host)
    cluster.Keyspace = keyspace
    session, err := cluster.CreateSession()

    if err != nil {
        log.Fatalf("Could not create CQL session: %s", err)
    }

    ctx := cqlc.NewContext()

    err = ctx.Upsert(TWEETS).
        SetString(TWEETS.NAME, "tweeter").
        Where(TWEETS.ID.Eq(1)).
        Exec(session)

    if err != nil {
        log.Fatalf("Could not execute CQL upsert: %s", err)
    }

    iter, err := ctx.Select().
        From(TWEETS).
        Where(TWEETS.ID.Eq(1)).
        Fetch(session)

    if err != nil {
        log.Fatalf("Could not execute CQL select: %s", err)
    }

    tweets := twitter.BindTweets(iter)

    err = iter.Close()
    if err != nil {
        log.Fatalf("Could not bind tweets: %s", err)
    }

    log.Printf("Got tweets: %+v\n", tweets)
}
