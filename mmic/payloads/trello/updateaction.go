package trello

type UpdateAction struct {
	Action Action `json="action"`
}

func (ua UpdateAction) GetType() string {
	if _, found := ua.Action.Data["listBefore"]; found {
		return "MoveCard"
	}
	if _, found := ua.Action.Data["old"]; found {
		if _, title := ua.Action.Data["old"].(map[string]interface{})["name"]; title {

			return "UpdateCard"
		}
	}
	if ua.Action.Type == "createCard" {
		return "CreateCard"
	}
	return ""
}
