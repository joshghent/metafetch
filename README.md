# MetaFetch

MetaFetch is a lightweight Go-based API that fetches and returns metadata from any URL in the **oEmbed** format. It’s built using the blazing-fast [Gin](https://github.com/gin-gonic/gin) framework and [GoQuery](github.com/PuerkitoBio/goquery) to fetch the page metadata.

## Features

- 🚀 **Fast & Lightweight** – Powered by Go and the Gin framework.
- 🌍 **oEmbed Compliant** – Outputs metadata in standard oEmbed JSON format.
- 📦 **Dockerized** – Easily deploy MetaFetch anywhere using Docker.
- ⚡ **Extensible** – Add more metadata fields or customize the response format.

## Getting Started

### Prerequisites

- **Go 1.20+** (if running locally)
- **Docker** (for containerization)
- A basic understanding of HTTP APIs

### Running Locally

To run the project locally, you can build and run it using Go:

1. Clone the repository:

    ```bash
    git clone https://github.com/joshghent/metafetch.git
    cd metafetch
    ```

2. Install the dependencies:

    ```bash
    go mod tidy
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

4. Access the API at `http://localhost:8080/oembed?url=<page_url>`, replacing `<page_url>` with a valid URL.

### Docker

MetaFetch is Dockerized for easy deployment. To build and run the Docker container:

1. Build the Docker image:

    ```bash
    docker build -t metafetch .
    ```

2. Run the Docker container:

    ```bash
    docker run -p 8080:8080 metafetch
    ```

3. Access the API at `http://localhost:8080/oembed?url=<page_url>`.

### GitHub Package Registry

MetaFetch is also available via the GitHub Container Registry. To pull the latest image:

```bash
docker pull ghcr.io/joshghent/metafetch:latest
```

## Usage

You can use MetaFetch to extract metadata from any publicly accessible URL. Simply make a GET request to the /oembed endpoint with the URL as a query parameter.

Example Request:
```
GET /oembed?url=https://example.com
```

Example Response:

```json
{
  "version": "1.0",
  "type": "link",
  "title": "Example Page Title",
  "author_name": "John Doe",
  "author_url": "https://example.com/author",
  "provider_name": "example.com",
  "provider_url": "https://example.com",
  "thumbnail_url": "https://example.com/thumbnail.jpg"
}
```

## GitHub Actions CI/CD

This project uses GitHub Actions for continuous integration and deployment. The Docker image is automatically built and pushed to the GitHub Container Registry on every push to the main branch.

## Contributing

Contributions are welcome! To get started:
* Fork the repository.
* Create a new feature branch (`git checkout -b feature/your-feature`).
* Make your changes.
* Push to your branch and open a pull request.
