# OutGuess // N17 Forensics Lab

A browser-based steganography workbench for puzzles and digital forensics. It combines an OutGuess-compatible JPEG payload reader, an N17 RGB-LSB PNG reader, and a quick file scanner.

## Features
- **OutGuess extraction (JPEG):** Go/WebAssembly retrieval routine running in the browser. A blank key uses the default OutGuess key; custom keys are supported.
- **N17 payload recovery (PNG):** checks RGB least-significant bits for this project's N17S marker, recovers payload, and decodes ROT13 when present.
- **Quick forensic scan:** reports image signatures and PNG chunks, searches printable strings for puzzle-related terms, and looks for data appended after JPEG EOI / PNG IEND markers.
- **Local-first:** selected images are read in the browser and are not sent to a server. No analytics or external JavaScript dependencies.

## Use it without installing anything
After GitHub Pages has been enabled for this repository, open:

**https://arsam92.github.io/-out-and-guss/**

One-time setup: repository Settings → Pages → build and deployment source: **GitHub Actions**. The workflow builds the Go/WASM extractor and deploys the static page on pushes to main.

## How to use
1. Open the hosted page and choose or drag in an original image.
2. For OutGuess, select a JPEG, optionally enter a key, and press **Extract OutGuess payload**.
3. For a PNG generated with the N17 LSB writer, press **Recover N17 PNG payload**.
4. Use **Run forensic scan** to inspect signatures, text metadata, strings, and appended bytes.

## Limitations
- This browser app extracts OutGuess payloads; embedding is not implemented.
- OutGuess reads JPEG coefficient data. PNG files use the separate N17 RGB-LSB reader, not the OutGuess algorithm.
- The N17 reader recognizes this project's N17S marker and bit layout; it is not a universal PNG steganography detector.
- Quick scan is heuristic. A clean report cannot prove an image contains no hidden information.
- Re-encoding, resizing, screenshots, or recompression can destroy hidden data. Keep original files.

## Development
Go 1.23+ is only needed to build/test the source; visitors do not need Go, Python, Node.js, or extensions. Run go test ./... to run Go checks. Build the browser extractor with GOOS=js GOARCH=wasm go build -o outguess.wasm ./cmd/wasm. The deployment workflow copies wasm_exec.js from the Go toolchain and publishes the static site.

## Credits
The Go OutGuess reader is derived from henkman/outguess (MIT; see LICENSE-OUTGUESS). Parts of the JPEG codec are derived from Go's image/jpeg implementation and the MPEG Software Simulation Group IDCT code; notices are preserved in source headers and THIRD_PARTY_NOTICES.md. The original OutGuess project was developed by Niels Provos. This is an independent browser workbench, not the official OutGuess project.