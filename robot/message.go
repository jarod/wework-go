package robot

import (
	"bytes"
	"encoding/json"
	"io"
)

// Msg robot message
type Msg struct {
	MsgType  string       `json:"msgtype"`
	Text     *TextMsg     `json:"text,omitempty"`
	Markdown *MarkdownMsg `json:"markdown,omitempty"`
}

// TextMsg text message
type TextMsg struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

// MarkdownMsg markdown message
type MarkdownMsg struct {
	Content string `json:"content"`
}

// Text create text message
func Text(content string) *Msg {
	return &Msg{
		MsgType: "text",
		Text: &TextMsg{
			Content: content,
		},
	}
}

// Mentioned set mentioned list. Text message only
func (m *Msg) Mentioned(l ...string) *Msg {
	if m.Text != nil {
		m.Text.MentionedList = l
	}
	return m
}

// MentionedMobile set mentioned mobile list. Text message only
func (m *Msg) MentionedMobile(l ...string) *Msg {
	if m.Text != nil {
		m.Text.MentionedMobileList = l
	}
	return m
}

// Markdown create markdown message
func Markdown(content string) *Msg {
	return &Msg{
		MsgType: "markdown",
		Markdown: &MarkdownMsg{
			Content: content,
		},
	}
}

// Reader reader of marshaled data
func (m *Msg) Reader() (io.Reader, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
