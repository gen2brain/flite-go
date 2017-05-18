## flite-go [![GoDoc](https://godoc.org/github.com/gen2brain/flite-go?status.svg)](https://godoc.org/github.com/gen2brain/flite-go)

Golang bindings for [Flite](http://www.speech.cs.cmu.edu/flite/index.html) (festival-lite)

### Requirements

* [Flite](http://www.speech.cs.cmu.edu/flite/index.html)

##### Ubuntu

    apt-get install flite-dev

##### Fedora

    dnf install flite-devel

### Installation

    go get -v github.com/gen2brain/flite-go

### Example

```go
package main

import "github.com/gen2brain/flite-go"

func main() {
	// The valid names are "awb", "kal16", "kal", "rms" and "slt"
	voice, err := flite.VoiceSelect("kal")
	if err != nil {
		panic(err)
	}

    	// Use "play" for output and it will be sent to the audio device
	flite.TextToSpeech("Hello World", voice, "output.wav")
}
```
