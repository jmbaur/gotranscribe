package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RecognitionConfig struct {
	// Required fields
	Encoding        string `json:"encoding"`
	SampleRateHertz int    `json:"sampleRateHertz"`
	LanguageCode    string `json:"languageCode"`
}

type RecognitionAudio struct {
	Content string `json:"content"` // base64 encoded string
}

type SpeechToTextRequest struct {
	Config RecognitionConfig `json:"config"`
	Audio  RecognitionAudio  `json:"audio"`
}

type SpeechRecognitionResult struct {
	Results []struct {
		Alternatives []struct {
			Confidence float64 `json:"confidence"`
			Transcript string  `json:"transcript"`
		} `json:"alternatives"`
	} `json:"results"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("need to supply a file as an argument")
	}

	file := os.Args[1]
	audio, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	fileData, err := ioutil.ReadAll(audio)
	if err != nil {
		log.Fatalf("failed to read file: %v\n", err)
	}
	audioData := base64.StdEncoding.EncodeToString(fileData)

	data, err := json.Marshal(SpeechToTextRequest{
		Config: RecognitionConfig{
			Encoding:        "FLAC",
			SampleRateHertz: 48000,
			LanguageCode:    "en-US",
		},
		Audio: RecognitionAudio{
			Content: audioData,
		},
	})

	res, err := http.Post(fmt.Sprintf("https://speech.googleapis.com/v1/speech:recognize?key=%s",
		os.Getenv("API_KEY")),
		"application/json",
		bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("request to Google Speech-to-Text API failed: %v\n", err)
	}
	defer res.Body.Close()

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response from Google: %v\n", err)
	}

	if res.StatusCode != 200 {
		log.Fatalln(string(responseData))
	}

	var result SpeechRecognitionResult
	if err := json.Unmarshal(responseData, &result); err != nil {
		log.Fatalf("failed to read data from Google: %v\n", err)
	}

	fmt.Printf("\nTranscript:\n")
	for _, t := range result.Results {
		fmt.Printf("%s", t.Alternatives[0].Transcript)
	}
}
