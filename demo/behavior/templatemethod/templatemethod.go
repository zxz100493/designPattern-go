package templatemethod

import "fmt"

type DownLoader interface {
	DownLoad(uri string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func NewTemplete(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) DownLoad(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("download finished\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}

type HttpDownLoader struct {
	*template
}

func NewHttpDownloader() DownLoader {
	downloader := &HttpDownLoader{}
	template := NewTemplete(downloader)
	downloader.template = template
	return downloader
}

func (h *HttpDownLoader) download() {
	fmt.Printf("download %s via http\n", h.uri)
}

func (*HttpDownLoader) save() {
	fmt.Printf("http save\n")
}

type FtpDownLoader struct {
	*template
}

func NewFtpDownloader() DownLoader {
	downloader := &FtpDownLoader{}
	template := NewTemplete(downloader)
	downloader.template = template
	return downloader
}

func (f *FtpDownLoader) download() {
	fmt.Printf("download %s via ftp\n", f.uri)
}
