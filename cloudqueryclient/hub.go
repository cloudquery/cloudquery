package cloudqueryclient

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

const cloudQueryPublicKey = `-----BEGIN PGP PUBLIC KEY BLOCK-----

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
NdQkdLbveQ+US4vVAzRFJjRAvGVq14lRxiTreQ==
=9Zuc
-----END PGP PUBLIC KEY BLOCK-----`

func sha256File(filePath string) (string, error){
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), err
}

func validateChecksumProvider(providerPath string, checksumPath string) error {
	sha256sum, err := sha256File(providerPath)
	if err != nil {
		return err
	}
	f, err := os.Open(checksumPath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "  ")
		if len(split) != 2 {
			return fmt.Errorf("checksum file in incorrect format")
		}
		if strings.Contains(split[1], runtime.GOOS) &&  split[1] == sha256sum {
			return nil
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func validateFile(targetPath string, signaturePath string) error {
	keyring, err := openpgp.ReadArmoredKeyRing(strings.NewReader(cloudQueryPublicKey))
	if err != nil {
		return err
	}

	target, err := os.Open(targetPath)
	if err != nil {
		return err
	}
	defer target.Close()

	signature, err := os.Open(signaturePath)
	if err != nil {
		return err
	}
	defer signature.Close()

	log.Println("validating provider signature...")
	_, err = openpgp.CheckDetachedSignature(keyring, target, signature)
	if err != nil {
		log.Println("provider signature validation failed")
		return err
	}

	return nil
}

func getProviderPath(name string, version string) (string, error) {
	org := "cloudquery"
	split := strings.Split(name, "/")
	if len(split) == 2 {
		org = split[0]
		name = split[1]
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.cq/providers/%s/%s/%s", workingDir, org, name, version), nil
}

func downloadProvider(name string, version string) error {
	org := "cloudquery"
	split := strings.Split(name, "/")
	if len(split) == 2 {
		org = split[0]
		name = split[1]
	}

	isRegistered, err := isProviderRegistered(org, name)
	if err != nil {
		return err
	}
	if !isRegistered {
		return fmt.Errorf("provider %s is not registered at https://hub.cloudquery.io", name)
	}

	providerURL := ""
	checksumsURL := ""
	if version == "" || version == "latest" {
		version = "latest"
		providerURL = fmt.Sprintf("https://github.com/%s/cq-provider-%s/releases/latest/download/cq-provider-%s_%s_%s",
			org, name, name, runtime.GOOS, runtime.GOARCH)
		checksumsURL = fmt.Sprintf("https://github.com/%s/cq-provider-%s/releases/latest/download/checksums.txt",
			org, name)
	} else {
		providerURL = fmt.Sprintf("https://github.com/%s/cq-provider-%s/releases/download/%s/cq-provider-%s_%s_%s",
			org, name, version, name, runtime.GOOS, runtime.GOARCH)
		checksumsURL = fmt.Sprintf("https://github.com/%s/cq-provider-%s/releases/download/%s/checksums.txt",
			org, name, version)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	pluginDir := fmt.Sprintf("%s/.cq/providers/%s/%s", workingDir, org, name)
	err = os.MkdirAll(pluginDir, os.ModePerm)
	if err != nil {
		return err
	}
	pluginPath := fmt.Sprintf("%s/%s-%s-%s", pluginDir, version, runtime.GOOS, runtime.GOARCH)
	checksumsPath := pluginDir + "/" + version + ".checksums.txt"

	log.Printf("downloading provider plugin %s/%s...\n", org, name)
	err = downloadFile(pluginPath, providerURL)
	if err != nil {
		return err
	}

	err = os.Chmod(pluginPath, 0754)
	if err != nil {
		return err
	}

	// download checksums
	err = downloadFile(checksumsPath, checksumsURL)
	if err != nil {
		return err
	}

	// download checksums signature
	err = downloadFile(checksumsPath + ".sig", checksumsURL + ".sig")
	if err != nil {
		return err
	}

	err = validateFile(checksumsPath, checksumsPath + ".sig")
	if err != nil {
		return err
	}

	err = validateChecksumProvider(pluginPath, checksumsPath)
	if err != nil {
		return err
	}

	return nil
}


func isProviderRegistered(org string, name string) (bool, error) {
	url := fmt.Sprintf(
		"https://firestore.googleapis.com/v1/projects/hub-cloudquery/databases/(default)/documents/orgs/%s/providers/%s",
		org, name)
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		switch res.StatusCode {
		case http.StatusNotFound:
			return false, nil
		default:
			return false, fmt.Errorf(res.Status)
		}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}


	return true, nil
}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusNotFound:
			return fmt.Errorf("file doesn't exist at %s", url)
		default:
			return fmt.Errorf(resp.Status)
		}
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

