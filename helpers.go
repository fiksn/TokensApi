/*
 * Copyright (C) 2019 Gregor Pogačnik
 */

package TokensApi

import (
	"TokensApi/entities"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

type CredentialsConfig struct {
	APIKey    string `json:"key" description:"Api key."`
	APISecret string `json:"secret" description:"Api secret."`
}

const (
	Timeout = time.Duration(5 * time.Second)
)

var (
	Credentials *CredentialsConfig
)

func Init(configPath string) {
	Credentials = parseJsonCfg(configPath)
}

func parseJsonCfg(configPath string) *CredentialsConfig {
	jsonBlob, err := ioutil.ReadFile(configPath)
	if err != nil {
		glog.Fatalf("Could not parse config: %v", err)
	}
	var cfg CredentialsConfig
	err = json.Unmarshal(jsonBlob, &cfg)
	if err != nil {
		glog.Fatalf("Unable to unmarshal json blob: %v", string(jsonBlob))
	}

	return &cfg
}

func request(url string) []byte {
	glog.V(3).Infof("request url: %v\n", url)

	client := http.Client{
		Timeout: Timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		glog.Warningf("request error %v", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Warningf("request error %v", err)
		return nil
	}

	return body
}

func calcHmac(key string, message string) string {
	sig := hmac.New(sha256.New, []byte(key))
	sig.Write([]byte(message))

	return strings.ToUpper(hex.EncodeToString(sig.Sum(nil)))
}

func requestAuthPost(url string, data url.Values) []byte {
	unixTime := strconv.Itoa(int(time.Now().UnixNano() / 1000))

	data.Add("key", Credentials.APIKey)
	data.Add("nonce", unixTime)
	signature := calcHmac(Credentials.APISecret, unixTime+Credentials.APIKey)
	data.Add("signature", signature)

	client := http.Client{
		Timeout: Timeout,
	}
	glog.V(3).Infof("requestAuth url: %v data %v\n", url, data)

	resp, err := client.PostForm(url, data)
	if err != nil {
		glog.Warningf("requestAuth error %v", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Warningf("requestAuth error %v", err)
		return nil
	}

	return body
}

func requestAuth(url string) []byte {
	unixTime := strconv.Itoa(int(time.Now().UnixNano() / 1000))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		glog.Warningf("requestAuth error %v", err)
		return nil
	}

	req.Header.Set("key", Credentials.APIKey)
	req.Header.Set("nonce", unixTime)

	signature := calcHmac(Credentials.APISecret, unixTime+Credentials.APIKey)

	req.Header.Set("signature", signature)

	client := &http.Client{Timeout: Timeout}

	glog.V(3).Infof("requestAuth url: %v\n", url)

	resp, err := client.Do(req)
	if err != nil {
		glog.Warningf("requestAuth error %v", err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Warningf("requestAuth error %v", err)
		return nil
	}

	return body
}

func deserialize(jsonBlob []byte, resp entities.Statuser) (err error) {

	var errorEntity entities.ErrorResp

	err = json.Unmarshal(jsonBlob, resp)

	if err != nil {
		glog.Warningf("Unable to unmarshal json blob: %v (%v)", string(jsonBlob), err)
		resp.SetStatus("error")
		return err
	}

	if resp.GetStatus() != "ok" {
		err = json.Unmarshal(jsonBlob, &errorEntity)
		if err != nil {
			return errors.New(errorEntity.Reason)
		}

		return errors.New(resp.GetStatus())
	}

	return nil
}
