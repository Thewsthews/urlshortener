# URL Shortener 〰️

This is a simple URL shortener built with Go. It provides an API to shorten long URLs and retrieve the original URL when accessed with the shortened code. The implementation uses an in-memory store (a map), but it can be extended to use a database like Redis or PostgreSQL for persistence.

## Features 🎭
- Shorten long URLs to short, unique codes.
- Redirect users from the shortened URL to the original URL.
- Uses an in-memory map for storage.

## Installation & Setup
### Prerequisites
- Go 1.18 or later installed

### Steps to Run 🍀
1. Clone the repository:
   ```sh
   git clone https://github.com/Thewsthews/url-shortener-go.git
   cd url-shortener-go
   ```
2. Run the application:
   ```sh
   go run main.go
   ```
3. The server will start on port `8080`.

## API Endpoints

### 1. Shorten a URL
**Endpoint:** `POST /shorten`

**Request Body:**
```json
{
  "long_url": "https://example.com"
}
```

**Response:**
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### 2. Redirect to Original URL
**Endpoint:** `GET /{shortURL}`

**Example:**
Visiting `http://localhost:8080/abc123` in a browser redirects to `https://example.com`.

## Notes
- Shortened URLs are stored in memory and will be lost when the server restarts.
- To make the URLs persistent, you can integrate a database like Redis.

## License
This project is open-source and available under the MIT License.

## Contributions
Feel free to submit pull requests or report issues!
## Contact
[Email](etiegnim@gmail.com)

M1