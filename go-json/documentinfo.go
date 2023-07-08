package main

type DocumentInfo struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description []string `json:"description"`
}
