package go_tour

func Start() {
	webCrawler()
}

type Fetcher interface {
	Fetch(url string) (body string, urlList[]string, err error)
}

type URL struct {
	url string
	depth int
}

func webCrawler() {

}