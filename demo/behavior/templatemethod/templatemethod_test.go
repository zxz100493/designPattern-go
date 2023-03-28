package templatemethod

func ExampleHttpDownloader() {
	var downloader DownLoader = NewHttpDownloader()
	downloader.DownLoad("http://www.baidu.com")
	// Output:
	// prepare downloading
	// download http://example.com/abc.zip via http
	// http save
	// finish downloading
}

func ExampleFtpDownloader() {
	var downloader DownLoader = NewFtpDownloader()
	downloader.DownLoad("http://www.baidu.com")
	// Output:
	// prepare downloading
	// download ftp://example.com/abc.zip via ftp
	// default save
	// finish downloading
}
