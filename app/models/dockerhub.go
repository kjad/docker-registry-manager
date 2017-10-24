package manager

import (
	"github.com/sirupsen/logrus"
	manifestV2 "github.com/docker/distribution/manifest/schema2"
	client "github.com/heroku/docker-registry-client/registry"
)

const (
	HubURL = "https://registry.hub.docker.com"
)

// HubUser and HubPassword set to empty so basic auth is not used
var HubUser = ""
var HubPassword = ""

// HubGetManifest retrieves the manifest for the given repo and tag name
func HubGetManifest(repoName string, tagName string) (*manifestV2.DeserializedManifest, error) {
	hub, err := client.New(HubURL, HubUser, HubPassword)
	if err != nil {
		return nil, err
	}
	// add and retrieve oauth token for connections
	hub.Client.Transport = client.WrapTransport(hub.Client.Transport, HubURL, HubUser, HubPassword)
	manifest, err := hub.ManifestV2("library/"+repoName, tagName)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error":           err.Error(),
			"Repository Name": repoName,
			"Tag Name":        tagName,
		}).Info("Failed to retrieve manifest information")
	}

	return manifest, err
}
