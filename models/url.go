package models

import (
	"errors"

	"github.com/google/uuid"
)

type url struct {
	Hash        string `json:"hash"`
	OriginalUrl string `json:"orgurl"`
}

//Creating url in memory initially
//This should be saved in a database
var urlList = []url{
	{Hash: "hash1", OriginalUrl: "www.google.com"},
	{Hash: "hash2", OriginalUrl: "www.yahoo.com"},
	{Hash: "hash3", OriginalUrl: "www.facebook.com"},
}

var hashSet = map[string]bool{
	"hash1": true,
	"hash2": true,
	"hash3": true,
}

//Returns the current list of urls
//This route is not a functionalty of tinyurl
//Created for debugging purpose
func GetAllUrls() []url {
	return urlList
}

//Returns the original url for a given tinyurl/hash
func GetOriginalUrlByHash(hash string) (*url, error) {
	for _, u := range urlList {
		if u.Hash == hash {
			return &u, nil
		}
	}
	return nil, errors.New("Url not found")
}

//Create and adds the tinyurl to the list for a given original url
func CreateNewUrl(originalurl string) (*url, error) {

	for {
		hash := getHash(originalurl)
		if !hashSet[hash] {

			u := url{Hash: hash, OriginalUrl: originalurl}
			urlList = append(urlList, u)
			hashSet[hash] = true

			return &u, nil
		}
	}
}

//Get hash for a given url - very simple uuid first 8 chars
//There will be duplicates, so inerim we have a hashset to retry this function
func getHash(originalurl string) (hash string) {

	id := uuid.New()
	hash = id.String()[0:8]

	return hash
}
