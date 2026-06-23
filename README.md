# Ascii-Art-Web

A lightweight web application built in Go that transforms standard text expressions into graphical ASCII art representations using distinct banner styles.

---

## Description

**Ascii-Art-Web** brings the classic command-line ASCII art generator to the web browser. The application sets up a fully functioning native Go HTTP server that serves an intuitive graphical user interface (GUI). Users can type text, select their preferred font style, and instantly generate ASCII art on the page.

### Key Features
* **Interactive Web GUI:** Simple and clean user interface built with HTML templates.
* **Multiple Banners:** Supports three iconic banner fonts: `standard`, `shadow`, and `thinkertoy`.
* **Robust Error Handling:** Returns strict, appropriate HTTP status codes based on API health and user inputs.
* **Zero External Dependencies:** Built entirely using Go's built-in standard library packages.

---

## Usage

### Prerequisites
* [Go (Golang)](https://go.dev/) installed on your machine (version 1.16 or higher recommended).

### Running the Server Local Environment

1. **Clone the repository:**
   ```bash
   git clone https://learn.zone01kisumu.ke
   cd ascii-art-web
   ```

2. **Start the Go web server:**
   ```bash
   go run main.go
   ```

3. **Open your web browser:**
   Navigate to the local network address provided by your terminal execution output:
   ```text
   http://localhost:8080
   ```

### Using the Web Interface
1. **Enter Text:** Type your desired text or phrase into the large text input field.
2. **Choose Banner:** Use the radio buttons or dropdown menu to select between `standard`, `shadow`, or `thinkertoy` styles.
3. **Generate:** Click the submission button to trigger the server rendering and display the output on your dashboard screen.

---

## Implementation Details

### Architecture & Endpoints
The backend utilizes Go's standard `net/http` package to route inbound traffic, parse form metadata, and handle response payloads. The application implements the following core HTTP endpoints:

* **`GET /`**
  Renders the primary dashboard page via Go's internal `html/template` processing framework. If a template file is missing or broken, the server halts gracefully and protects user interaction.
* **`POST /ascii-art`**
  Intercepts the client form submission payload containing the user's raw text and the target banner selection. It passes this data to the generation engine and returns the completed text output block back onto the webpage UI template.

### Expected HTTP Status Codes
The backend continuously evaluates request validity and yields appropriate response statuses:
* **`200 OK`**: The request succeeded, and the ASCII art was successfully generated.
* **`400 Bad Request`**: Malformed payloads or invalid text parameters were received.
* **`404 Not Found`**: The requested web endpoint, banner file, or template layout resource does not exist.
* **`500 Internal Server Error`**: The server hit an unhandled exception or file processing conflict.

### The Core Algorithm
1. **Form Extraction:** The server extracts data from the `POST` form request. It checks both the raw user string and the chosen banner format.
2. **File Processing:** The backend looks for the chosen banner file inside the root repository directory (e.g., `standard.txt`). It maps the character lines step by step.
3. **Array Splitting:** The input string is broken down into text lines by splitting on newline characters (`\n`).
4. **Graphical Assembly:** Every unique character from the input text is matched to its corresponding location in the banner file. Since each ASCII character is exactly 8 lines high, the engine prints line-by-line across all input characters to construct the visual graphics accurately.
5. **UI Update:** The raw text block is escaped safely and injected into a `<pre>` HTML container tag within the template block for clean viewing.

---

## Authors

This project was built by a collaborative development group at **Zone01 Kisumu**:

* **Taheera Abdallah** ([@tabdalla](https://learn.zone01kisumu.ke)) — *Group Captain*
* **Mercy Moraa** ([@memoraa](https://learn.zone01kisumu.ke))
* **Valentine Omondi** ([@valeomondi](https://learn.zone01kisumu.ke))
