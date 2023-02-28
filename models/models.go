package models

import (
    "database/sql"
    "os"
    _ "github.com/mattn/go-sqlite3"
    pusher "github.com/pusher/pusher-http-go"
)

var client = pusher.Client{
    AppID: os.Getenv("AppID"),
    Key: os.Getenv("Key"),
    Secret: os.Getenv("Secret"),
    Cluster: os.Getenv("Cluster"),
    Secure: true,
    }

    type Poll struct {
        ID          int     `json:"id"`
        Name        string  `json:"name"`
        Topic       string  `json:"topic"`
        Src         string  `json:"src"`
        Upvotes     int     `json:"uptvotes"`
        Downvotes   int     `json:"downvotes"`
    }

type PollCollection struct {
    Polls []Poll `json:"items"`
}

func GetPolls(db *sql.DB) PollCollection  {
    sql := "SELECT * FROM polls"

    rows, err :=db.Query(sql)

    if err != nil {
        panic(err)
    }

    defer rows.Close()

    result := PollCollection{}

    for rows.Next(){
        poll := Poll{}

        err2 := rows.Scan(&poll.ID, &poll.Name, &poll.Topic, &poll.Src, &poll.Upvotes, &poll.Downvotes)

        if err2 != nil{
            panic(err2)
        }

        result.Polls = append(result.Polls, poll)
    }

    return result
}

func UpdatePoll(db *sql.DB, index int, name string, upvotes int, downvotes int) (int64, error) {
    sql := "UPDATE polls SET (upvotes, downvotes) = (?, ?) WHERE id = ?"

    stmt, err := db.Prepare(sql)

    if err != nil {
        panic(err)
    }

    defer stmt.Close()

    result, err2 := stmt.Exec(upvotes, downvotes, index)

    if err2 != nil {
        panic(err2)
    }

    pollUpdate := Poll{
        ID:        index,
        Name:      name,
        Upvotes:   upvotes,
        Downvotes: downvotes,
        }

        client.Trigger("poll-channel", "poll-update", pollUpdate)
    return result.RowsAffected()
}