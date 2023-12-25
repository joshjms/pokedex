package models

type Pokemon struct {
	ID          int       `json:"id"`
	Name        Name      `json:"name"`
	Type        []string  `json:"type"`
	Base        Base      `json:"base"`
	Species     string    `json:"species"`
	Description string    `json:"description"`
	Evolution   Evolution `json:"evolution"`
	Profile     Profile   `json:"profile"`
	Image       Image     `json:"image"`
}

type Name struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Chinese  string `json:"chinese"`
	French   string `json:"french"`
}

type Base struct {
	HP        int `json:"HP"`
	Attack    int `json:"Attack"`
	Defense   int `json:"Defense"`
	SpAttack  int `json:"Sp. Attack"`
	SpDefense int `json:"Sp. Defense"`
	Speed     int `json:"Speed"`
}

type Evolution struct {
	Prev []interface{}   `json:"prev"`
	Next [][]interface{} `json:"next"`
}

type Profile struct {
	Height  string          `json:"height"`
	Weight  string          `json:"weight"`
	Egg     []string        `json:"egg"`
	Ability [][]interface{} `json:"ability"`
	Gender  string          `json:"gender"`
}

type Image struct {
	Sprite    string `json:"sprite"`
	Thumbnail string `json:"thumbnail"`
	Hires     string `json:"hires"`
}
