## Project Structure

```console
reflex/
├── cmd/                  # Entry points for the application
│   └── reflex/           # Specific executable directory
│       └── reflex.go     # Main application logic
├── utils/                # Utilities and helpers
│   ├── parser.go         # For parsing and validating URLs
│   ├── injector.go       # For modifying URLs and injecting identifiers
├── clients/              # Packages for making requests
│   ├── headless.go       # For headless browser requests
│   ├── request.go        # For non-browser HTTP requests
├── services/             # Core logic for processing and checking
│   ├── requesthandler.go # Logic for sending requests and analyzing responses
│   ├── checker.go        # Logic for validating identifier reflection in responses
├── config/               # Configuration and constants
│   ├── settings.go       # App settings, e.g., timeouts, headers, etc.
```