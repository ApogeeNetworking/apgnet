package thinobot

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ApogeeNetworking/apgnet/apgreq"
)

// Service ...
type Service struct {
	baseURL string
	http    *apgreq.Service
}

// NewService ...
func NewService(uri string, req *apgreq.Service) *Service {
	return &Service{
		baseURL: uri,
		http:    req,
	}
}

// GetBots ...
func (s *Service) GetBots() ([]Bot, error) {
	var bots []Bot
	uri := fmt.Sprintf("%s/bot", s.baseURL)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return bots, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return bots, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&bots); err != nil {
		return bots, err
	}
	return bots, nil
}

// GetBot ...
func (s *Service) GetBot(id string) (Bot, error) {
	var bot Bot
	uri := fmt.Sprintf("%s/bot/%s", s.baseURL, id)
	req, err := s.http.GenerateRequest(uri, "GET", nil)
	if err != nil {
		return bot, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return bot, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&bot); err != nil {
		return bot, err
	}
	return bot, nil
}

// CreateBot ...
func (s *Service) CreateBot(bot Bot) (Bot, error) {
	var newBot Bot
	data, err := json.Marshal(&bot)
	if err != nil {
		return bot, err
	}
	sr := bytes.NewReader(data)
	uri := fmt.Sprintf("%s/bot", s.baseURL)
	req, err := s.http.GenerateRequest(uri, "POST", sr)
	if err != nil {
		return newBot, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return newBot, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&newBot); err != nil {
		return newBot, err
	}
	return newBot, nil
}

// UpdateBot ...
func (s *Service) UpdateBot(bot Bot) (Bot, error) {
	var updatedBot Bot
	data, err := json.Marshal(&bot)
	if err != nil {
		return updatedBot, err
	}
	sr := bytes.NewReader(data)
	uri := fmt.Sprintf("%s/bot/%s", s.baseURL, bot.ID)
	req, err := s.http.GenerateRequest(uri, "PUT", sr)
	if err != nil {
		return updatedBot, err
	}
	res, err := s.http.MakeRequest(req)
	if err != nil {
		return updatedBot, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&updatedBot); err != nil {
		return updatedBot, err
	}
	return updatedBot, nil
}

func (s *Service) DeleteBot() {}

// ToggleBotWatcher ...
func (s *Service) ToggleBotWatcher(id string, watch bool) error {
	action := "watch"
	if !watch {
		action = "unwatch"
	}
	uri := fmt.Sprintf("%s/bot/%s/%s", s.baseURL, id, action)
	req, err := s.http.GenerateRequest(uri, "PUT", nil)
	if err != nil {
		return err
	}
	_, err = s.http.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) NotifyBot() {}
