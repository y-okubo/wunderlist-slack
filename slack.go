package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const Username = "My Bot"

var (
	// IncomingURL is secret
	IncomingURL = "incoming_url"
)

// Field is struct
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// Attachment is struct
type Attachment struct {
	Fallback   string  `json:"fallback"`
	Color      string  `json:"color"`
	Pretext    string  `json:"pretext"`
	AuthorName string  `json:"author_name"`
	AuthorLink string  `json:"author_link"`
	AuthorIcon string  `json:"author_icon"`
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Text       string  `json:"text"`
	Fields     []Field `json:"fields"`
	ImageURL   string  `json:"image_url"`
	ThumbURL   string  `json:"thumb_url"`
	Footer     string  `json:"footer"`
	FooterIcon string  `json:"footer_icon"`
	Ts         int     `json:"ts"`
}

// Slack is message struct
type Slack struct {
	Text        string       `json:"text"`
	Username    string       `json:"username"`
	IconEmoji   string       `json:"icon_emoji"`
	IconURL     string       `json:"icon_url"`
	Channel     string       `json:"channel"`
	Attachments []Attachment `json:"attachments"`
}

func slack(orders []uint, todos Todos) {
	var attachments []Attachment

	for _, ID := range orders {
		_, ok := todos[ID]
		if ok {
			f := []Field{
				{
				// Title: todos[ID].Title,
				// Value: underway(todos[ID].Starred),
				// Short: false,
				},
			}

			a := Attachment{
				Title:  todos[ID].Title,
				Text:   underway(todos[ID].Starred),
				Color:  color(todos[ID].Starred),
				Fields: f,
			}

			attachments = append(attachments, a)
		}
	}

	params, _ := json.Marshal(Slack{
		"やることリスト",
		Username,
		"",
		"",
		"#general",
		attachments})

	resp, _ := http.PostForm(
		IncomingURL,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}

func underway(flag bool) string {
	if flag {
		return "着手"
	}

	return "未着手"
}

func color(flag bool) string {
	if flag {
		return "good"
	}

	return "danger"
}
