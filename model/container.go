package model

import (
	"time"
)

// This struct is the pure data structure for the application,
// You are free to add any tag you want or leave it emtpy
// if you leave it emtpy your http handlers can have a copy of the model with the tags needed
// for example you could omit certain fields to an user that is not an admin while showing everything to an admin
// or you could decide to define them here and use them in the packages.
// You choose.
type (
	AppPost struct {
		GithubRepoUrl string `json:"github-repo"`
		GithubBranch  string `json:"github-branch"`
		Language      string `json:"language"`
		Port          string `json:"port"`
		Description   string `json:"description,omitempty"`
		Envs          []Env  `json:"envs,omitempty"`
	}

	Env struct {
		Key   string `json:"key"`
		Value string `bson:"value" json:"value"`
	}

	Application struct {
		ID        string    `bson:"_id" json:"-"`
		CreatedAt time.Time `bson:"createdAt" json:"createdAt,omitempty"`
		StudentID int       `bson:"studentID" json:"studentID,omitempty"`

		ContainerID  string `bson:"containerID" json:"containerID,omitempty"`
		Status       string `bson:"status" json:"status,omitempty"`
		Type         string `bson:"type" json:"type,omitempty"`
		Port         string `bson:"port" json:"port,omitempty"`
		ExternalPort string `bson:"externalPort" json:"externalPort,omitempty"`
		Lang         string `bson:"lang" json:"lang,omitempty"`
		Img          string `bson:"img,omitempty" json:"img,omitempty"`
		Envs         []Env  `bson:"envs,omitempty" json:"envs,omitempty"`

		Name           string   `bson:"name" json:"name,omitempty"`
		HostName       string   `bson:"hostName" json:"hostName,omitempty"`
		Description    string   `bson:"description" json:"description,omitempty"`
		GithubRepo     string   `bson:"githubRepo,omitemtpy" json:"githubRepo,omitempty"`
		GithubBranch   string   `bson:"githubBranch,omitemtpy" json:"githubBranch,omitempty"`
		LastCommitHash string   `bson:"lastCommitHash,omitemtpy" json:"lastCommitHash,omitempty"`
		IsPublic       bool     `bson:"isPublic" json:"isPublic"`
		IsUpdatable    bool     `bson:"isUpdatable,omitempty" json:"isUpdatable"`
		Tags           []string `bson:"tags,omitempty" json:"tags,omitempty"`
		Stars          []string `bson:"stars,omitempty" json:"stars,omitempty"`
	}
)
