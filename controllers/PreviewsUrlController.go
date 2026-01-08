package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type PreviewRequest struct {
	URL string `json:"url" binding:"required"`
}

type PreviewResponse struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	SiteName    string `json:"site_name"`
}

func HandlePreview(c *gin.Context) {
	var body PreviewRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format or 'url' field is missing"})
		return
	}

	targetURL := body.URL

	if !strings.HasPrefix(targetURL, "http") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL must start with http:// or https://"})
		return
	}

	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(targetURL)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Failed to reach the target URL"})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Target website rejected the request"})
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse website content"})
		return
	}

	preview := PreviewResponse{URL: targetURL}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		prop, _ := s.Attr("property")
		name, _ := s.Attr("name")
		content, _ := s.Attr("content")

		if prop == "og:title" || name == "twitter:title" {
			preview.Title = content
		}
		if prop == "og:description" || name == "description" || name == "twitter:description" {
			preview.Description = content
		}
		if prop == "og:image" || name == "twitter:image" {
			preview.Image = content
		}
		if prop == "og:site_name" {
			preview.SiteName = content
		}
	})

	if preview.Title == "" {
		preview.Title = doc.Find("title").First().Text()
	}

	c.JSON(http.StatusOK, preview)
}
