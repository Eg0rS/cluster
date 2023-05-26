package model

import "mime/multipart"

type RadioTest struct {
	Title       string
	Description string
	Questions   []Question
	TestType    string
}

type Question struct {
	Title   string
	Answers []Answer
}

type Answer struct {
	Text    string
	IsRight bool
}

type Request struct {
	Title          string
	Description    string
	TestId         int
	UserId         int
	OrganizationId int
}

type GetRequests struct {
	UserId string
}

type CreateTextTest struct {
	Title       string
	Description string
	File        multipart.File
	Header      *multipart.FileHeader
}

type TextTest struct {
	Title       string
	Description string
	FileName    string
	TestType    string
}

type AllTests struct {
	RadioTests []RadioTest
	TextTests  []TextTest
}
