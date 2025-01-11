// Initial template developed with claude ai

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// Article represents a news article from RTE
type Article struct {
	Title       string
	Content     string
	URL         string
	PublishDate time.Time
	Category    string
	Author      string
}

// RTEScraper handles the scraping of RTE news articles
type RTEScraper struct {
	collector *colly.Collector
	articles  []Article
}

// NewRTEScraper creates a new scraper instance
func NewRTEScraper() *RTEScraper {
	c := colly.NewCollector(
		colly.AllowedDomains("www.rte.ie", "rte.ie"),
		// Respect robots.txt
		colly.UserAgent("NewsScraperBot/1.0"),
	)

	// Add rate limiting to be respectful
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 5 * time.Second,
		Parallelism: 2,
	})

	return &RTEScraper{
		collector: c,
		articles:  make([]Article, 0),
	}
}

// ScrapeArticles scrapes articles from the RTE news section
func (s *RTEScraper) ScrapeArticles(category string) ([]Article, error) {
	baseURL := fmt.Sprintf("https://www.rte.ie/news/%s/", category)

	// Set up callbacks for the collector
	s.collector.OnHTML("article", func(e *colly.HTMLElement) {
		// Extract article details
		article := Article{
			Title:    strings.TrimSpace(e.ChildText("h1")),
			URL:      e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
			Category: category,
		}

		// Only proceed if we found a valid article
		if article.Title != "" && article.URL != "" {
			// Visit the article page to get full content
			s.collector.Visit(article.URL)
		}
	})

	// Handle individual article pages
	s.collector.OnHTML("div.article-body", func(e *colly.HTMLElement) {
		content := []string{}
		e.ForEach("p", func(_ int, p *colly.HTMLElement) {
			text := strings.TrimSpace(p.Text)
			if text != "" {
				content = append(content, text)
			}
		})

		// Get publish date
		dateStr := e.ChildText("time")
		publishDate, _ := time.Parse("2006-01-02 15:04:05", dateStr)

		article := Article{
			Title:       e.ChildText("h1"),
			Content:     strings.Join(content, "\n"),
			URL:         e.Request.URL.String(),
			PublishDate: publishDate,
			Category:    category,
			Author:      e.ChildText(".author"),
		}

		s.articles = append(s.articles, article)
	})

	// Error handling
	s.collector.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping %s: %v", r.Request.URL, err)
	})

	// Start the scraping
	err := s.collector.Visit(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to start scraping: %w", err)
	}

	// Wait for scraping to finish
	s.collector.Wait()

	return s.articles, nil
}

func main() {
	scraper := NewRTEScraper()
	
	// Example: Scrape business articles
	articles, err := scraper.ScrapeArticles("business")
	if err != nil {
		log.Fatalf("Failed to scrape articles: %v", err)
	}

	// Print results
	for _, article := range articles {
		fmt.Printf("Title: %s\nURL: %s\nDate: %s\n\n",
			article.Title,
			article.URL,
			article.PublishDate.Format("2006-01-02"))
	}
}