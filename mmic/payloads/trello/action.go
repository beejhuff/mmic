package trello

import (
	"time"
)

/**
Json payload example.

{
    "id": "56e3e87819415768324a7d0d",
    "idMemberCreator": "51482f9e7b37e943390046c6",
    "data": {
      "list": {
        "name": "HI",
        "id": "56e3e2405958f895130c56fb"
      },
      "board": {
        "shortLink": "5LSKmLTd",
        "name": "webhooktest",
        "id": "56e3e23a807617aa236c9bcb"
      },
      "card": {
        "shortLink": "uJqoB3Sl",
        "idShort": 4,
        "name": "a a a",
        "id": "56e3e7f8c66458cc426ae2e2",
        "pos": 147455
      },
      "old": {
        "pos": 73727
      }
    },
    "type": "updateCard",
    "date": "2016-03-12T09:59:20.078Z",
    "memberCreator": {
      "id": "51482f9e7b37e943390046c6",
      "avatarHash": "8d1b88b7f5fbc54ff6c376b63dd2d05d",
      "fullName": "Frank Wang",
      "initials": "FW",
      "username": "frankwang2"
    }
  }

**/

type Action struct {
	Id              string                 `json:"id"`
	IdMemberCreator string                 `json:"idMemberCreator"`
	Type            string                 `json:"type"`
	Date            time.Time              `json:"date"`
	Data            map[string]interface{} `json:"data"`
	MemberCreator   *Member                `json"memberCreator"`
}
