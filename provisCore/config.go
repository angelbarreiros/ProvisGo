package provisCore

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type provisConfig struct {
	AplicationKey string
	SecretKey     string
}

func NewConfig(applicationKey, secretKey string) *provisConfig {
	if strings.TrimSpace(applicationKey) == "" ||
		strings.TrimSpace(secretKey) == "" {
		log.Fatal("Invalid configuration: InstallationID, AplicationKey, and SecretKey must be provided")
	}
	return &provisConfig{
		AplicationKey: applicationKey,
		SecretKey:     secretKey}
}
func (pc *provisConfig) generateRequest(installationId string, method string, uri string, queryParams url.Values, params any) *http.Request {
	var nonce uuid.UUID = uuid.New()
	var timeStamp string = strconv.FormatInt(time.Now().Unix(), 10)
	var request = &http.Request{
		Method: method,
		URL: &url.URL{
			Scheme:   "https",
			Host:     "apibase-integraciones.provis.es",
			Path:     uri,
			RawQuery: queryParams.Encode(),
		},
		Header: make(http.Header),
	}
	var body io.Reader
	if params != nil {
		jsonBody, err := json.Marshal(params)
		if err == nil {
			body = bytes.NewReader(jsonBody)
		}
		request.Body = io.NopCloser(body)
	}

	if body != nil {
		if bodyBytes, ok := body.(*bytes.Reader); ok {
			request.ContentLength = int64(bodyBytes.Len())
		}
		request.Header.Set("Content-Type", "application/json")
	} else {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	var fullURL string = request.URL.String()

	request.Header.Set("Timestamp", time.Now().Format(time.RFC3339))
	request.Header.Set("Authorization", "hmac-256 "+installationId+":"+pc.AplicationKey+":"+pc.generateSign(fullURL, method, params, nonce, timeStamp)+":"+nonce.String()+":"+timeStamp)
	request.Header.Set("Cache-Control", "no-cache")
	return request

}
func (pc *provisConfig) generateSign(fullUrl string, method string, params any, nonce uuid.UUID, timeStamp string) string {
	var encodedUrl string = strings.ToLower(url.QueryEscape(fullUrl))
	var hashedAndEncodedParams string = ""
	if params != nil {
		var jsonBytes []byte = nil
		var err error
		jsonBytes, err = json.Marshal(params)
		if err != nil {
			log.Fatalf("Error marshalling params: %v", err)
		}
		var hash [16]byte = md5.Sum(jsonBytes)
		var encodedHash string = base64.StdEncoding.EncodeToString(hash[:])
		hashedAndEncodedParams = encodedHash

	}
	var sign = pc.AplicationKey + method + encodedUrl + timeStamp + nonce.String() + hashedAndEncodedParams
	var hmacHash = hmac.New(sha256.New, []byte(pc.SecretKey))
	var _, err = hmacHash.Write([]byte(sign))
	if err != nil {
		log.Fatalf("Error generating HMAC hash: %v", err)
	}
	var encodedHmacHash = base64.StdEncoding.EncodeToString(hmacHash.Sum(nil))
	return string(encodedHmacHash)
}
