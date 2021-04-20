package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type RecognitionConfig struct {
	Encoding string
}

//"encoding": enum,
//"sampleRateHertz": integer,
//"audioChannelCount": integer,
//"enableSeparateRecognitionPerChannel": boolean,
//"languageCode": string,
//"maxAlternatives": integer,
//"profanityFilter": boolean,
//"speechContexts": [
//{
//object (SpeechContext)
//}
//],
//"enableWordTimeOffsets": boolean,
//"enableAutomaticPunctuation": boolean,
//"diarizationConfig": {
//object (SpeakerDiarizationConfig)
//},
//"metadata": {
//object (RecognitionMetadata)
//},
//"model": string,
//"useEnhanced": boolean

type RecognitionAudio struct {
	Content string `json:"content"` // base64 encoded string
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
	bytes, err := ioutil.ReadAll(audio)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(bytes)
}
