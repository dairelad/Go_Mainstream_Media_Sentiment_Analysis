// Initial template developed with claude ai

package main

import (
	"fmt"

	//"log"
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
	AuthorRole  string
	Tags        []string
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
	fmt.Printf("Visiting %s\n", baseURL)

	// Create a temporary variable to store the current article being processed
	var currentArticle Article
	// Add a flag at the top of your ScrapeArticles function
	var firstArticleFound bool = false

	// Set up callbacks for the collector
	s.collector.OnHTML("article", func(e *colly.HTMLElement) {

		if firstArticleFound { // stop scraping after the first article
			return
		}
		// Extract article details
		currentArticle = Article{
			Title:    strings.TrimSpace(e.ChildText("h3")),
			URL:      e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
			Category: category,
		}

		// Set flag to true after finding first article
		firstArticleFound = true

		//fmt.Print(currentArticle)

		// Only proceed if we found a valid article
		if currentArticle.Title != "" && currentArticle.URL != "" {
			fmt.Printf("Visiting article URL: %s\n", currentArticle.URL)
			e.Request.Visit(currentArticle.URL)
		}
	})

	// Handle individual article pages
	s.collector.OnHTML("div[itemtype='http://schema.org/Article']", func(e *colly.HTMLElement) {
		var articleParagraphs []string

		// Extract all paragraphs within the article body
		e.ForEach("p", func(_ int, p *colly.HTMLElement) {
			text := strings.TrimSpace(p.Text)
			if text != "" {
				articleParagraphs = append(articleParagraphs, text)
			}
		})
		//runtime.Breakpoint()
		title := e.ChildText("h1.headline")
		author := e.ChildText(".byline-link strong")
		authorRole := e.ChildText(".byline .article-meta p")

		fmt.Println(articleParagraphs)
		// Get publish date
		dateStr := e.ChildText(".modified-date")

		//fmt.Println(dateStr)
		dateStr = strings.TrimSpace(strings.Split(dateStr, ", ")[1])

		publishDate, err := time.Parse("02 Jan 2006 15:04", dateStr)
		if err != nil {
			// If date parsing fails, use current time and log the error
			publishDate = time.Now()
			fmt.Printf("Error parsing date '%s': %v\n", dateStr, err)
		}

		// Extract tags/categories
		var tags []string
		e.ForEach(".tags-container .tags li a", func(_ int, t *colly.HTMLElement) {
			tag := strings.TrimSpace(t.Text)
			if tag != "" {
				tags = append(tags, tag)
				fmt.Printf("Tag found: %s\n", tag)
			}
		})

		// Create the complete article
		article := Article{
			Title:       strings.TrimSpace(title),
			Content:     strings.Join(articleParagraphs, "\n"),
			URL:         e.Request.URL.String(),
			Author:      author,
			AuthorRole:  authorRole,
			PublishDate: publishDate,
			Tags:        tags,
			Category:    category,
		}

		fmt.Printf("Processed article: %s\n", article.Title)
		s.articles = append(s.articles, article)
	})

	// Error handling
	s.collector.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error scraping %s: %v\n", r.Request.URL, err)
	})

	// Start the scraping
	err := s.collector.Visit(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to start scraping: %w", err)
	}

	// Wait for scraping to finish
	s.collector.Wait()

	fmt.Printf("Scraped %d articles\n", len(s.articles))
	return s.articles, nil
}

func main() {
	scraper := NewRTEScraper()
	fmt.Print("Starting RTE Scraper..\n")

	// Scrape Political Articles on the RTE Website
	articles, err := scraper.ScrapeArticles("politics")
	if err != nil {
		fmt.Printf("Failed to scrape articles: %v", err)
	}

	// Print results
	for _, article := range articles {
		fmt.Printf("Title: %s\nURL: %s\nDate: %s\n\n",
			article.Title,
			article.URL,
			article.PublishDate.Format("2006-01-02"))
	}
}
