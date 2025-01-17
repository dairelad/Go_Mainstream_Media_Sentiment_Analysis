# Media Sentiment Analysis

A tool for analyzing political bias in media coverage of companies through sentiment analysis. This project fetches news articles about specified companies and analyzes their sentiment to identify potential political leanings in their coverage.

## Features

- Automated news article collection from multiple sources
- Sentiment analysis to detect political bias
- Data visualization of sentiment trends over time
- Export results to JSON for further analysis

## Prerequisites
- Go 1.19 or later

## Installation

1. Clone the repository:
```bash
git clone https://github.com/dairelad/Go_Mainstream_Media_Sentiment_Analysis.git
cd Go_Mainstream_Media_Sentiment_Analysis
```

2. Install dependencies:
```bash
go mod tidy
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

1. **Data Collection**: The program scapes the target website.

2. **Sentiment Analysis**: Each article is processed using sentiment analysis to determine its political leaning:
   - Negative scores (-1 to 0) suggest left-leaning coverage
   - Positive scores (0 to 1) suggest right-leaning coverage

3. **Visualization**: Results are plotted showing trends over time and by news source.

## License

This project is licensed under the MIT License - see the LICENSE file for details.