# Flanger Audio Effect

A real-time audio flanger effect processor written in Go. This application captures audio from your microphone, applies a flanger effect, and outputs the processed audio to your speakers or headphones.

## What is a Flanger?

A flanger is an audio effect that mixes the original signal with a slightly delayed copy of itself, where the delay time is modulated by a low-frequency oscillator. This creates a sweeping, "whooshing" sound similar to a jet plane flying overhead - the effect that gave flanging its name.

## Features

- Real-time audio processing with low latency
- High-fidelity audio (144kHz sample rate)
- Adjustable flanger depth via command-line parameter
- Clean signal handling for graceful shutdown

## Requirements

- Go 1.23.2 or higher
- PortAudio library installed on your system

## Installation

### 1. Install PortAudio

#### macOS
```bash
brew install portaudio
```

#### Linux
```bash
sudo apt-get install portaudio19-dev
```

#### Windows
Download and install PortAudio from [the official website](http://www.portaudio.com/download.html)

### 2. Clone the repository
```bash
git clone https://github.com/yourusername/flanger.git
cd flanger
```

### 3. Build the application
```bash
go build -o flanger
```

## Usage

Run the application with an optional depth parameter (in seconds):

```bash
./flanger [depth]
```

Example:
```bash
./flanger 0.01  # Run with a depth of 0.01 seconds
```

If no depth parameter is provided, the application will prompt you to enter one.

To stop the application, press `Ctrl+C`.

## How It Works

The application:

1. Captures audio input from your default microphone
2. Stores recent audio samples in a delay buffer
3. Calculates a continuously varying delay time using sine wave modulation
4. Mixes the original signal with the delayed signal
5. Outputs the processed audio to your default output device

## Parameters

- **Sample Rate**: 144kHz for high-fidelity audio processing
- **Buffer Size**: 256 samples for a good balance between latency and CPU load
- **Modulation Frequency**: 0.25Hz for a smooth flanger sweep
- **Delay Buffer**: 50ms to accommodate the delay effect
- **Depth**: User-configurable parameter that controls the intensity of the effect

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgements

- [PortAudio](http://www.portaudio.com/) for the cross-platform audio I/O library
- [gordonklaus/portaudio](https://github.com/gordonklaus/portaudio) for the Go bindings to PortAudio

