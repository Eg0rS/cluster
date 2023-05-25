package personnel

import "mime/multipart"

type RadioTest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

type Question struct {
	Title   string   `json:"title"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Text    string `json:"text"`
	IsRight bool   `json:"is_right"`
}

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	TestId      int    `json:"test_id"`
	UserId      int    `json:"user_id"`
}

type TextTest struct {
	Title       string
	Description string
	File        multipart.File
	Header      *multipart.FileHeader
}
