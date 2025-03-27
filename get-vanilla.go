package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const versionManifestURL = "https://launchermeta.mojang.com/mc/game/version_manifest.json"

type VersionManifest struct {
	Latest struct {
		Release string `json:"release"`
	} `json:"latest"`
	Versions []struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"versions"`
}

type VersionData struct {
	Downloads struct {
		Server struct {
			URL string `json:"url"`
		} `json:"server"`
	} `json:"downloads"`
}

func main() {
	// Fetch the version manifest
	resp, err := http.Get(versionManifestURL)
	if err != nil {
		fmt.Println("Error fetching version manifest:", err)
		return
	}
	defer resp.Body.Close()

	var manifest VersionManifest
	if err := json.NewDecoder(resp.Body).Decode(&manifest); err != nil {
		fmt.Println("Error decoding version manifest:", err)
		return
	}

	latestVersion := manifest.Latest.Release
	var versionURL string
	for _, v := range manifest.Versions {
		if v.ID == latestVersion {
			versionURL = v.URL
			break
		}
	}

	// Fetch the specific version details
	resp, err = http.Get(versionURL)
	if err != nil {
		fmt.Println("Error fetching version details:", err)
		return
	}
	defer resp.Body.Close()

	var versionData VersionData
	if err := json.NewDecoder(resp.Body).Decode(&versionData); err != nil {
		fmt.Println("Error decoding version data:", err)
		return
	}

	serverURL := versionData.Downloads.Server.URL
	fileName := fmt.Sprintf("minecraft_server.%s.jar", latestVersion)

	// Download the server JAR
	resp, err = http.Get(serverURL)
	if err != nil {
		fmt.Println("Error downloading server JAR:", err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Printf("Downloaded Minecraft server version %s as %s\n", latestVersion, fileName)
}
