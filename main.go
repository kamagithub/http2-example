package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var httpAddr = flag.String("http", ":8080", "Listen address")

func main() {
	flag.Parse()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		pusher, ok := w.(http.Pusher)
		if ok {
			// Push is supported. Try pushing rather than
			// waiting for the browser request these static assets.

			if err := pusher.Push("/static/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app1.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app2.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app3.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app5.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app6.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app7.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app8.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app9.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app10.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app11.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}

			if err := pusher.Push("/static/app12.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err := pusher.Push("/static/app13.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		fmt.Fprintf(w, indexHTML)
	})
	log.Fatal(http.ListenAndServeTLS(*httpAddr, "cert.pem", "key.pem", nil))
}

const indexHTML = `<html>
<head>
	<title>Hello World</title>
	<script src="/static/app.js"></script>
	<script src="/static/app1.js"></script>
	<script src="/static/app2.js"></script>
	<script src="/static/app3.js"></script>
	<script src="/static/app4.js"></script>
	<script src="/static/app5.js"></script>
	<script src="/static/app6.js"></script>
	<script src="/static/app7.js"></script>
	<script src="/static/app8.js"></script>
	<script src="/static/app9.js"></script>
	<script src="/static/app10.js"></script>
	<script src="/static/app11.js"></script>
	<script src="/static/app12.js"></script>
	<script src="/static/app13.js"></script>
	<script src="/static/app14.js"></script>
</head>
<body>
Hello, gopher!
</body>
</html>
`
