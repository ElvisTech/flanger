package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gordonklaus/portaudio"
)

const (
	sampleRate    = 144000 // High-fidelity sample rate
	bufferSize    = 256    // Good balance between latency and CPU load
	modFreq       = 0.25   // Low modulation frequency for smooth flanger sweep
	delayBufferMs = 50     // Delay buffer size to accommodate the delay effect
)

func main() {
	// Parse depth parameter from command line (in milliseconds)
	depth := 0.008 // Default depth in seconds, subtle for high fidelity
	if len(os.Args) > 1 {
		// Convert depth argument to float
		userDepth, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil || userDepth <= 0 {
			log.Fatalf("Invalid depth value. Depth must be a positive number.")
		}
		depth = userDepth
	} else {
		fmt.Print("enter the depth in secons")
	}

	// Initialize PortAudio
	err := portaudio.Initialize()
	if err != nil {
		log.Fatalf("Error initializing PortAudio: %v", err)
	}
	defer portaudio.Terminate()

	// Create delay buffer
	delayBufferSize := int(delayBufferMs * sampleRate / 1000)
	delayBuffer := make([]float32, delayBufferSize)
	bufferIndex := 0
	phase := 0.0
	modulationRate := 2 * math.Pi * modFreq / sampleRate

	// Open PortAudio stream with both input and output
	stream, err := portaudio.OpenDefaultStream(1, 1, sampleRate, bufferSize, func(in, out []float32) {
		for i := 0; i < len(in); i++ {
			// Calculate modulated delay time
			modulatedDelay := depth * (1 + math.Sin(phase))
			phase += modulationRate
			if phase > 2*math.Pi {
				phase -= 2 * math.Pi
			}
			delaySamples := int(modulatedDelay * sampleRate)

			// Calculate read index with wrapping
			readIndex := (bufferIndex - delaySamples + len(delayBuffer)) % len(delayBuffer)

			// Apply flanger effect by mixing original input with delayed signal
			out[i] = 0.7*in[i] + 0.5*delayBuffer[readIndex] // Increased mix for a stronger effect

			// Store the current sample in the delay buffer
			delayBuffer[bufferIndex] = in[i]
			bufferIndex = (bufferIndex + 1) % len(delayBuffer)
		}
	})
	if err != nil {
		log.Fatalf("Error opening PortAudio stream: %v", err)
	}
	defer stream.Close()

	// Start the audio stream
	err = stream.Start()
	if err != nil {
		log.Fatalf("Error starting PortAudio stream: %v", err)
	}
	log.Printf("Microphone pass-through with flanger effect started with depth %v seconds. Press Ctrl+C to stop.", depth)

	// Wait for CTRL+C to terminate the program
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // Block until a signal is received

	// Stop the stream
	err = stream.Stop()
	if err != nil {
		log.Fatalf("Error stopping PortAudio stream: %v", err)
	}
}
