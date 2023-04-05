# Go library for [indexnow](https://www.indexnow.org/)

IndexNow is a simple ping so that search engines know that a URL has been changed.

This is a go library that allows to index ping search engines for website update.

So far bing and yandex supports the website, and google [confirmed](https://www.searchenginejournal.com/google-will-be-testing-indexnow/426602/) that they will be testing it 

In order to use this library first you have to:
1) generate an api key and host an api key as specified in https://www.bing.com/indexnow

2) After that request indexing using this go sdk

Here is a go snippet on how to request indexing a single url
```golang
s, _ := indexnow.NewIndexer(indexnow.YANDEX, "brainya.com", "123412345263245cf12451", nil)
res, err := s.IndexSinglePage(context.Background(), "https://brainya.com/u/HwL692uxCTMOqhILPPqEg5AZ5IJ2/a/360857755814199888")
if err != nil {
	log.Fatal(err)
}
fmt.Println("status code", res.StatusCode)
```
and for bulk update
```golang
s, _ := indexnow.NewIndexer(indexnow.BING, "brainya.com", "123412345263245cf12451", nil)
res, err := s.BulkIndexPages(context.Background(), &[]string{
	"https://brainya.com/u/HwL692uxCTMOqhILPPqEg5AZ5IJ2/a/360857755814199888",
})
if err != nil {
	log.Fatal(err)
}
fmt.Println("status code", res.StatusCode)
```
