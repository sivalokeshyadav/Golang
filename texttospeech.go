package main

import "github.com/hegedustibor/htgo-tts"

func TTS() {
    speech := htgotts.Speech{Folder: "audio", Language: "en"}
    speech.Speak("Flying to the moon")
}
   