package senders

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/mmic/mmic/config"
	"github.com/eternnoir/mmic/mmic/payloads/mattermost"
	"github.com/eternnoir/mmic/mmic/utils"
)

type MatterMostSender struct {
	WebhookUrl string
	Channel    string
	Username   string
}

func NewMatterMostSender(mmconfig config.TargetMatterMostConfig) *MatterMostSender {
	mms := &MatterMostSender{WebhookUrl: mmconfig.WebhookUrl, Channel: mmconfig.Channel, Username: mmconfig.Username}
	return mms
}

func (mms *MatterMostSender) Send(payload interface{}) error {
	mmpayload, found := payload.(*mattermost.MatterMostPayload)
	log.Debugf("MMSender get payload %#v to send.", payload)
	if !found {
		log.Debugf("MMSernder get not MMPayload payload")
		return fmt.Errorf("MatterMostSender can not parse payload. Check your payload is MatterMostPayload. %#v", payload)
	}
	return mms.sendToChannel(mmpayload)
}

func (mms *MatterMostSender) sendToChannel(payload *mattermost.MatterMostPayload) error {
	log.Infof("Send payload %#v to %s", payload, mms.WebhookUrl)
	jsonbody, err := payload.Serialize()
	if err != nil {
		return err
	}
	jsonStr := string(jsonbody)
	log.Infof("Start to fire payload to %s. Payload text %s", mms.WebhookUrl, string(jsonStr))
	_, err = utils.PostRequest(mms.WebhookUrl, jsonStr)
	if err != nil {
		log.Errorf("Send to targer: %#v payload %#v fail. Error: %s.", mms, payload, err)
	}
	return nil
}
