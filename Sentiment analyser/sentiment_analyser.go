package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cdipaolo/sentiment"
	"github.com/go-resty/resty/v2"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Article represents a news article with its metadata and analysis
type Article struct {
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Source         string    `json:"source"`
	PublishedDate  time.Time `json:"published_date"`
	SentimentScore *float64  `json:"sentiment_score,omitempty"`
}

// NewsCollector handles fetching articles from news APIs
type NewsCollector struct {
	apiKey string
	client *resty.Client
}

// NewNewsCollector creates a new NewsCollector instance
func NewNewsCollector(apiKey string) *NewsCollector {
	return &NewsCollector{
		apiKey: apiKey,
		client: resty.New(),
	}
}

// FetchArticles retrieves articles about a specific company
func (nc *NewsCollector) FetchArticles(company string) ([]Article, error) {
	// TODO: Implement actual API call
	// This is where you'd make the HTTP request to your chosen news API
	// Example using NewsAPI endpoint:
	/*
		resp, err := nc.client.R().
			SetQueryParams(map[string]string{
				"q": company,
				"apiKey": nc.apiKey,
			}).
			Get("https://newsapi.org/v2/everything")
	*/

	// Placeholder return
	return []Article{}, nil
}

// SentimentAnalyzer handles the sentiment analysis of text
type SentimentAnalyzer struct {
	model *sentiment.Models
}

// NewSentimentAnalyzer creates a new SentimentAnalyzer instance
func NewSentimentAnalyzer() (*SentimentAnalyzer, error) {
	model, err := sentiment.Restore()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize sentiment model: %w", err)
	}
	return &SentimentAnalyzer{model: model}, nil
}

// AnalyzeText performs sentiment analysis on the given text
func (sa *SentimentAnalyzer) AnalyzeText(text string) float64 {
	// This is a simplified approach - you'd want to tune this
	analysis := sa.model.SentimentAnalysis(text, sentiment.English)
	// Convert to a -1 to 1 scale where negative means left-leaning
	score := float64(analysis.Score)*2 - 1
	return score
}

// Visualizer handles the creation of visualizations
type Visualizer struct{}

// CreateTrendChart generates a visualization of sentiment trends
func (v *Visualizer) CreateTrendChart(articles []Article, outputPath string) error {
	p := plot.New()
	p.Title.Text = "Media Sentiment Analysis"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Sentiment Score"

	// Create scatter plot data
	pts := make(plotter.XYs, len(articles))
	for i, article := range articles {
		if article.SentimentScore != nil {
			pts[i].X = float64(article.PublishedDate.Unix())
			pts[i].Y = *article.SentimentScore
		}
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return fmt.Errorf("failed to create scatter plot: %w", err)
	}
	p.Add(s)

	// Save the plot
	if err := p.Save(6*vg.Inch, 4*vg.Inch, outputPath); err != nil {
		return fmt.Errorf("failed to save plot: %w", err)
	}

	return nil
}

func main() {
	// Get API key from environment
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Fatal("NEWS_API_KEY environment variable is required")
	}

	// Initialize components
	collector := NewNewsCollector(apiKey)
	analyzer, err := NewSentimentAnalyzer()
	if err != nil {
		log.Fatalf("Failed to initialize sentiment analyzer: %v", err)
	}
	visualizer := &Visualizer{}

	// Fetch articles
	articles, err := collector.FetchArticles("Example Corp")
	if err != nil {
		log.Fatalf("Failed to fetch articles: %v", err)
	}

	// Analyze sentiment for each article
	for i := range articles {
		score := analyzer.AnalyzeText(articles[i].Content)
		articles[i].SentimentScore = &score
	}

	// Create visualization
	if err := visualizer.CreateTrendChart(articles, "sentiment_trends.png"); err != nil {
		log.Fatalf("Failed to create visualization: %v", err)
	}

	// Optionally save results to JSON
	jsonData, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal results: %v", err)
	}
	if err := os.WriteFile("results.json", jsonData, 0644); err != nil {
		log.Fatalf("Failed to save results: %v", err)
	}
}
