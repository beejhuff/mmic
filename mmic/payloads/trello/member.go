package trello

/**
Example json payload

      "id": "518f2cf9a64e0e0612000e08",
      "avatarHash": "745e8c711185ddccb3ca099e2d2f079b",
      "fullName": "Josh Holbrook",
      "initials": "JFH",
      "username": "jesusabdullah"
**/

type Member struct {
	Id         string `json:"id"`
	AvatarHash string `json:"avatarHash"`
	FullName   string `json:"fullName"`
	Initials   string `json:"initials"`
	Username   string `json:"username"`
}
