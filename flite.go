// Package flite implements bindings for Flite (festival-lite)
package flite

/*
#cgo CFLAGS: -I/usr/include/flite
#cgo LDFLAGS: -lflite_cmu_us_awb -lflite_cmu_us_kal16 -lflite_cmu_us_kal -lflite_cmu_us_rms -lflite_cmu_us_slt -lflite_usenglish -lflite_cmulex -lflite -lm

#include "flite.h"
#include <stdlib.h>

cst_voice* register_cmu_us_awb();
cst_voice* register_cmu_us_kal16();
cst_voice* register_cmu_us_kal();
cst_voice* register_cmu_us_rms();
cst_voice* register_cmu_us_slt();
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Voice type
type Voice struct {
	name       *int8
	features   *features
	ffunctions *features
	init       *utterance
}

func (v *Voice) cptr() *C.cst_voice {
	return (*C.cst_voice)(unsafe.Pointer(v))
}

// newVoiceFromPointer returns new Voice from pointer
func newVoiceFromPointer(ptr unsafe.Pointer) *Voice {
	return (*Voice)(ptr)
}

// Wave type
type Wave struct {
	Type        *int8
	SampleRate  int32
	NumSamples  int32
	NumChannels int32
	Padding     [4]byte
	Samples     *int16
}

func (w *Wave) cptr() *C.cst_wave {
	return (*C.cst_wave)(unsafe.Pointer(w))
}

// newWaveFromPointer returns new Wave from pointer
func newWaveFromPointer(ptr unsafe.Pointer) *Wave {
	return (*Wave)(ptr)
}

// features type
type features struct {
	head *featValPair
	ctx  unsafe.Pointer
}

// utterance type
type utterance struct {
	features   *features
	ffunctions *features
	relations  *features
	ctx        unsafe.Pointer
}

// featValPair type
type featValPair struct {
	name string
	val  *val
	next *featValPair
}

// val type
type val struct {
	C [16]byte
}

// VoiceSelect returns a Voice for the voice name.
// The valid names are "awb", "kal16", "kal", "rms", "slt".
// It will return error if there is no match.
func VoiceSelect(name string) (*Voice, error) {
	var ret *C.struct_cst_voice_struct
	switch name {
	case "awb":
		ret = C.register_cmu_us_awb()
	case "kal16":
		ret = C.register_cmu_us_kal16()
	case "kal":
		ret = C.register_cmu_us_kal()
	case "rms":
		ret = C.register_cmu_us_rms()
	case "slt":
		ret = C.register_cmu_us_slt()
	default:
		return nil, fmt.Errorf("flite: no match for voice: %s", name)
	}

	if ret == nil {
		return nil, fmt.Errorf("flite: no voice")
	}

	v := newVoiceFromPointer(unsafe.Pointer(ret))
	return v, nil
}

// TextToSpeech synthesizes the text with the given voice.
// outtype may be a filename where the generated wav is written to, or "play" and it will be sent to the audio device, or "none" and it will be discarded.
// The return value is the number of seconds of speech generated.
func TextToSpeech(text string, voice *Voice, outtype string) float64 {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	cvoice := voice.cptr()

	couttype := C.CString(outtype)
	defer C.free(unsafe.Pointer(couttype))

	ret := C.flite_text_to_speech(ctext, cvoice, couttype)
	v := (float64)(ret)
	return v
}

// TextToWave returns a waveform synthesized from the text with the given voice.
func TextToWave(text string, voice *Voice) *Wave {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	cvoice := voice.cptr()

	ret := C.flite_text_to_wave(ctext, cvoice)
	v := newWaveFromPointer(unsafe.Pointer(ret))
	return v
}
