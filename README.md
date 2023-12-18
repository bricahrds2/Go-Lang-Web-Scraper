# Invasion Club Scraper

This Go script is designed to scrape product information from the Invasion Club website, specifically focusing on the "tops" category. The collected data includes product titles, original prices, discounted prices, and image URLs.

## Prerequisites
Make sure you have Go installed on your machine. If not, you can download and install it from https://golang.org/dl/.

## Running the Script
1. Clone this repository to your local machine:

``` Go
git clone https://github.com/your-username/invasion-club-scraper.git
cd invasion-club-scraper
``` 
2. Run the Go script:

``` Go
go run main.go
``` 
This will initiate the scraping process and print the titles of the visited pages to the console.

3. Check the Output

Once the script completes, it will generate a JSON file named page-products.json, containing the scraped data.

## Output Format

The output JSON file consists of an array of objects, where each object represents a product with the following attributes:

* 'Title': Product title.
* 'OrigPrice': Original product price.
* 'DescPrice': Discounted product price.
* 'Img': URL of the product image.

Dependencies
This script utilizes the github.com/gocolly/colly package for web scraping. You can install it using:

``` Go
go get -u github.com/gocolly/colly/...
```

## Notes

The script is configured to scrape the "tops" category on the Invasion Club website (https://invasion.club/collections/tops). You can modify the URL in the c.Visit("https://invasion.club/collections/tops") line to target other categories.

The script may be subject to changes based on the structure of the website. If the website structure is updated, you may need to adjust the HTML element selectors accordingly.

Feel free to explore and modify the script according to your needs! If you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request. Happy scraping!
