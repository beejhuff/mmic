package trello

type UpdateAction struct {
	Action Action `json="action"`
}

func (ua *UpdateAction) GetType() string {
	if _, found := ua.Action.Data["listBefore"]; found {
		return "MoveCard"
	}
	if _, found := ua.Action.Data["old"]; found {
		return "UpdateCard"
	}
	if ua.Action.Type == "createCard" {
		return "CreateCard"
	}
	return ""
}
