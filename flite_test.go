package flite

import (
	"testing"
	"time"
)

func TestTextToSpeech(t *testing.T) {
	voice, err := VoiceSelect("awb")
	if err != nil {
		t.Fatal(err)
	}

	s := TextToSpeech("Hello World", voice, "play")
	if s == 0 {
		t.Fatalf("0 seconds of speech generated")
	}
}

func TestTextToWave(t *testing.T) {
	voice, err := VoiceSelect("awb")
	if err != nil {
		t.Fatal(err)
	}

	wav := TextToWave("Hello World", voice)
	if wav.NumChannels != 1 {
		t.Fatalf("Wave number of channels != 1")
	}

	if wav.SampleRate != 16000 {
		t.Fatalf("Wave sample rate != 1")
	}
}

func TestVoiceSelect(t *testing.T) {
	var voices = []string{"awb", "kal16", "kal", "rms", "slt"}

	for _, v := range voices {
		voice, err := VoiceSelect(v)
		if err != nil {
			t.Fatal(err)
		}

		TextToSpeech("Testing "+v, voice, "play")
		time.Sleep(1 * time.Second)

		s := TextToSpeech("Hello World", voice, "play")
		if s == 0 {
			t.Fatalf("%s: 0 seconds of speech generated", v)
		}

		time.Sleep(1 * time.Second)
	}
}
