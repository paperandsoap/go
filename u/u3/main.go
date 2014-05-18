package u3

import (
	"net/http"

	"gist.github.com/7390843.git"

	"github.com/shurcooL/go/u/u1"
	"github.com/shurcooL/go/u/u4"
)

// Displays given Markdown in a new browser window/tab.
func DisplayMarkdownInBrowser(markdown []byte) {
	stopServerChan := make(chan struct{})

	handler := func(w http.ResponseWriter, req *http.Request) {
		u1.WriteMarkdownGfmAsHtmlPage(w, markdown)

		stopServerChan <- struct{}{}
	}

	http.HandleFunc("/index", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// TODO: Aquire a free port similarly to using ioutil.TempFile() for files.
	u4.Open("http://localhost:7044/index")

	err := gist7390843.ListenAndServeStoppable("localhost:7044", nil, stopServerChan)
	if err != nil {
		panic(err)
	}
}