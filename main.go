package main

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type OEmbedResponse struct {
	Version       string `json:"version"`
	Type          string `json:"type"`
	Title         string `json:"title,omitempty"`
	AuthorName    string `json:"author_name,omitempty"`
	AuthorURL     string `json:"author_url,omitempty"`
	ProviderName  string `json:"provider_name,omitempty"`
	ProviderURL   string `json:"provider_url,omitempty"`
	ThumbnailURL  string `json:"thumbnail_url,omitempty"`
	ThumbnailWidth  int    `json:"thumbnail_width,omitempty"`
	ThumbnailHeight int    `json:"thumbnail_height,omitempty"`
	URL           string `json:"url,omitempty"`
}

func main() {
	r := gin.Default()

	r.GET("/oembed", func(c *gin.Context) {
		pageURL := c.Query("url")

		// Validate URL
		parsedURL, err := url.ParseRequestURI(pageURL)
		if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
			return
		}

		// Fetch the webpage
		resp, err := http.Get(pageURL)
		if err != nil || resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the page"})
			return
		}
		defer resp.Body.Close()

		// Parse the page
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse the page"})
			return
		}

		// Build oEmbed response
		oembed := OEmbedResponse{
			Version:      "1.0",
			Type:         "link", // Default type, can be more specific based on content type
			ProviderName: parsedURL.Host,
			ProviderURL:  parsedURL.Scheme + "://" + parsedURL.Host,
		}

		// Extract metadata
		doc.Find("meta").Each(func(i int, s *goquery.Selection) {
			if name, exists := s.Attr("name"); exists {
				content, _ := s.Attr("content")
				switch name {
				case "title", "og:title":
					oembed.Title = content
				case "author", "og:author":
					oembed.AuthorName = content
				case "og:url":
					oembed.URL = content
				case "og:image":
					oembed.ThumbnailURL = content
				}
			}

			if property, exists := s.Attr("property"); exists {
				content, _ := s.Attr("content")
				switch property {
				case "og:title":
					oembed.Title = content
				case "og:author":
					oembed.AuthorName = content
				case "og:url":
					oembed.URL = content
				case "og:image":
					oembed.ThumbnailURL = content
				}
			}
		})

		c.JSON(http.StatusOK, oembed)
	})

	r.Run(":8080") // Run on port 8080
}
