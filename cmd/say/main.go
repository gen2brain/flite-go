package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gen2brain/flite-go"
)

func main() {
	v := flag.String("v", "slt", "voice to use, valid values are awb, kal16, kal, rms, slt")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("usage:\n\t%s \"text to speak\"\n", os.Args[0])
		os.Exit(1)
	}

	voice, err := flite.VoiceSelect(*v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	flite.TextToSpeech(strings.Join(flag.Args(), " "), voice, "play")
}
