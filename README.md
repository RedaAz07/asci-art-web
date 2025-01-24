# Ascii-Art-Web

## Description
Ascii-Art-Web is a web-based application that allows users to generate ASCII art from text using different banner styles. The project extends the functionality of the previous ascii-art project by providing a graphical user interface (GUI) via a web server. Users can input text, select a banner style (shadow, standard, or thinkertoy), and view the ASCII art output on a webpage.

## Objectives
- Create and run an HTTP server.
- Implement a web GUI for generating ASCII art.
- Support the following banner styles:
  - `shadow`
  - `standard`
  - `thinkertoy`
- Provide the following HTTP endpoints:
  - `GET /`: Displays the main page with the input form.
  - `POST /ascii-art`: Processes user input and returns the generated ASCII art.

## HTTP Status Codes
- **200 OK**: Request was successful.
- **404 Not Found**: Templates or banners not found.
- **400 Bad Request**: Invalid requests.
- **500 Internal Server Error**: Unhandled server errors.

## Main Page Features
The main page includes:
1. **Text Input**: A field to input text for ASCII art generation.
2. **Banner Selection**: Radio buttons or a dropdown menu to select one of the banner styles.
3. **Submit Button**: A button to send the input and banner style to the server.
4. **Output Display**: Displays the generated ASCII art on the same page or on a separate results page.

## Authors
- Ahmed Baid (abaid)
- Reda Anniz (ranniz)

## Usage
### Running the Project
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd ascii-art-web
   ```
2. Build and run the server:
   ```bash
   go run main.go
   ```
3. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

### Input Format
- Enter any text in the input field.
- Select a banner style (shadow, standard, or thinkertoy).
- Click the "Generate" button to display the ASCII art.

## Implementation Details
### Algorithm
1. **Text Input Parsing**: The server receives text input and banner style via a POST request.
2. **ASCII Art Generation**:
   - The backend processes the text input using the selected banner style.
   - The ASCII art generation logic is reused from the ascii-art project.
3. **Response Handling**:
   - If successful, the ASCII art is displayed on the webpage.
   - If an error occurs (e.g., invalid input), an appropriate error message is shown.

### Project Structure
```
ASCII-ART-WEB/
├── cmd/         // entry point of the project
│   ├── main.go
├── files/           // Contains ASCII art files (e.g., shadow.txt, standard.txt, thinkertoy.txt)
│   ├── shadow.txt
│   ├── standard.txt
│   ├── thinkertoy.txt
├── functions/       // Go logic for ASCII art generation
│   ├── Ascii.go
│   ├── PrintAscii.go
├── handler/          // Handler functions (StyleHandler,ResultHandler,FormHandler)
│   ├── Handlers.go
├── styles/          // CSS styles for the application
│   ├── index.css
│   ├── result.css
│   ├── stausPage.css
├── template/        // HTML templates
│   ├── index.html
│   ├── result.html
│   ├── stausPage.html
├── go.mod           // Go module file
└── README.md        // Documentation

```

### Instructions
1. **Server Implementation**:
   - Write the HTTP server in Go.
   - Use `net/http` for routing and handling requests.
2. **Templates**:
   - Store all HTML templates in the `templates/` directory.
   - Use Go's `html/template` package to parse and render templates.
3. **Best Practices**:
   - Ensure clean and modular code.
   - Handle errors gracefully.
4. **Allowed Packages**:
   - Only standard Go packages are permitted.

## Example
1. **Input**:
   - Text: `Hello`
   - Banner: `shadow`

2. **Output**:
```                              
_|                _| _|          
_|_|_|     _|_|   _| _|   _|_|   
_|    _| _|_|_|_| _| _| _|    _| 
_|    _| _|       _| _| _|    _| 
_|    _|   _|_|_| _| _|   _|_|   
                                 
```

Visit the application at [http://localhost:8080](http://localhost:8080) to test it yourself!

