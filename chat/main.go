package main

import (
	"html/template"
	"net/http"
	"log"
	"path/filepath"
	"sync"
	"flag"
	"os"
	"go_programing/trace"
)


//templは一つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

//serveHTTPはHTTPリクエストを返します
func (t *templateHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() //フラグを解釈します
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	//ルート
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	//チャットを開始します
	go r.run()
	// サーバーを開始します
	log.Println("webサーバーを起動します。ポート：", *addr)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}



