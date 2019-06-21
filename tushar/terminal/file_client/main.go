package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	client "github.com/tushar/terminal/file_client/client"
	server "github.com/tushar/terminal/file_server/server"
)

var (
	serverAddr = "localhost"
	serverPort = "8000"
)

func visit_file_path(path string, f os.FileInfo, err error) error {
	if f.IsDir() || filepath.HasPrefix(f.Name(), ".") {
		return nil
	}
	var req server.FileStats
	fmt.Printf("Visited File Path:%s", path)
	req.Path = path
	req.FileSize = f.Size()
	req.FileExt = filepath.Ext(path)
	req.Name = filepath.Base(path)
	// fmt.Printf("\t{File name:%s,", filepath.Base(path))
	// fmt.Printf("\tExtension:%s,", filepath.Ext(path))
	// fmt.Printf("\tSize:%d bytes,", f.Size())
	// fmt.Printf("\tMode:%d bytes,", f.Mode())
	// fmt.Printf("\tModification time:%s}\n", f.ModTime())
	err = sendInfo(req)
	if err != io.EOF && err != nil {
		fmt.Printf("Send err %v", err)
		return err
	}
	return nil
}

func sendInfo(fileInfo server.FileStats) error {
	baseURL, err := url.Parse("http://" + serverAddr + ":" + serverPort)
	if err != nil {
		return err
	}
	c := &client.Client{
		BaseUrl:    baseURL,
		HttpClient: http.DefaultClient,
	}
	return c.SendFileInfo(fileInfo)
}

func main() {
	addr := flag.String("addr", "localhost", "Server address to send info")
	port := flag.String("port", "8000", "port of rest server")
	path := flag.String("path", "/", "path to traverse")
	flag.Parse()

	serverAddr = *addr
	serverPort = *port
	root_dir := *path
	if _, err := os.Stat(root_dir); os.IsNotExist(err) {
		fmt.Println("Path:", root_dir, "does not exists.")
	} else {
		filepath.Walk(root_dir, visit_file_path)
	}

}
