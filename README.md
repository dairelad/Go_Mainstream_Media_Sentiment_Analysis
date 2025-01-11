### Initial template generated with claude ai

# Media Sentiment Analysis

A tool for analyzing political bias in media coverage of companies through sentiment analysis. This project fetches news articles about specified companies and analyzes their sentiment to identify potential political leanings in their coverage.

## Features

- Automated news article collection from multiple sources
- Sentiment analysis to detect political bias
- Data visualization of sentiment trends over time
- Support for multiple companies/topics
- Export results to JSON for further analysis

## Prerequisites

- Go 1.19 or later
- News API key (register at https://newsapi.org)
- Basic understanding of sentiment analysis concepts

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/media-sentiment-analysis
cd media-sentiment-analysis
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up your environment variables:
```bash
export NEWS_API_KEY='your_api_key_here'
```

## Usage

1. Run the basic analysis:
```bash
go run main.go -company "Example Corp"
```

2. Advanced usage with custom parameters:
```bash
go run main.go -company "Example Corp" -days 30 -sources "nytimes,wsj,reuters"
```

### Configuration Options

- `-company`: Name of the company to analyze (required)
- `-days`: Number of days of articles to analyze (default: 30)
- `-sources`: Comma-separated list of news sources (default: all available)
- `-output`: Output file path for visualization (default: "sentiment_trends.png")

## Output

The program generates several outputs:

1. A PNG visualization of sentiment trends (`sentiment_trends.png`)
2. A JSON file containing detailed analysis results (`results.json`)
3. Console output showing analysis progress and summary statistics

## Project Structure

```
.
├── main.go              # Main application entry point
├── collector/           # News article collection logic
├── analyzer/           # Sentiment analysis implementation
├── visualizer/         # Data visualization components
└── models/             # Data structures and interfaces
```

## How It Works

1. **Data Collection**: The program uses the News API to fetch recent articles about the specified company.

2. **Sentiment Analysis**: Each article is processed using sentiment analysis to determine its political leaning:
   - Negative scores (-1 to 0) suggest left-leaning coverage
   - Positive scores (0 to 1) suggest right-leaning coverage

3. **Visualization**: Results are plotted showing trends over time and by news source.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [News API](https://newsapi.org) for providing news data access
- [sentiment](https://github.com/cdipaolo/sentiment) for sentiment analysis capabilities
- [gonum/plot](https://github.com/gonum/plot) for visualization tools

## Contact

Your Name - [@yourtwitter](https://twitter.com/yourtwitter)

Project Link: [https://github.com/yourusername/media-sentiment-analysis](https://github.com/yourusername/media-sentiment-analysis)

## Disclaimer

This tool provides a simplified analysis of media bias and should not be considered definitive. Results should be interpreted as indicative rather than conclusive, and users are encouraged to perform their own verification of findings.