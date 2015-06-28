package main

import (
	"time"
)

const (
	TABLE_MESSAGES = "messages"
)

type MessagesModel struct {
	Id         int64     `db:"id" json:"id"`
	Sender     int64     `db:"msgby" json:"msgby"`
	Sent       time.Time `db:"msgtime" json:"msgtime"`
	For        int64     `db:"msgfor" json:"msgfor"`
	Recipients string    `db:"msgrecip" json:"msgrecip"`
	Patient    int64     `db:"msgpatient" json:"msgpatient"`
	Person     string    `db:"msgperson" json:"msgperson"`
	Urgency    int       `db:"msgurgency" json:"msgurgency"`
	Subject    string    `db:"msgsubject" json:"msgsubject"`
	Text       string    `db:"msgtext" json:"msgtext"`
	Read       int       `db:"msgread" json:"msgread"`
	Unique     string    `db:"msgunique" json:"msgunique"`
	Tag        string    `db:"msgtag" json:"msgtag"`
	Active     string    `db:"active" json:"active"`
}

func init() {
	dbTables = append(dbTables, DbTable{TableName: TABLE_MESSAGES, Obj: MessagesModel{}})
}
