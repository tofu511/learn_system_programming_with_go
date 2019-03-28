package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	// クライアントにgzipを返すwriterを定義
	compressor := gzip.NewWriter(w)

	// MultiWriterでクライアントへのgzipと標準出力をするmultiWriterを定義
	multiWriter := io.MultiWriter(compressor, os.Stdout)

	// multiWriterへ書き込む encoderを定義
	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "	")

	// 標準出力にjsonを書き込む
	encoder.Encode(source)

	// gzipを書き込む
	compressor.Flush()
	compressor.Close()

}

func main()  {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}