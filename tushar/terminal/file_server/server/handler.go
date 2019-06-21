package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FileStats struct {
	Name     string `json:"fileName"`
	Path     string `json:"path"`
	FileSize int64  `"json:"fileSize"`
	FileExt  string `"json:"fileExt"`
}

type FileStatsInfo struct {
	FileCount         int            `json:"fileCount"`
	MaxFileSize       int64          `json:"maxFileSize"`
	AvgFileSize       int64          `json:"avgFileSize"`
	FileExtension     map[string]int `json:"countFileExtension"`
	FreqFileExtension string         `json:"freqFileExtension"`
	RecentFilePaths   []string       `json:"recentFilePaths"`
}

var FileInfoCounter FileStatsInfo

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to file stats server")
}

func ShowFileStat(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(FileInfoCounter); err != nil {
		panic(err)
	}
}

func ProcessFileInfo(w http.ResponseWriter, r *http.Request) {
	var req FileStats
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Json decoding error "+err.Error(), http.StatusInternalServerError)
		return
	}

	FileInfoCounter.FileCount = FileInfoCounter.FileCount + 1
	if FileInfoCounter.MaxFileSize < req.FileSize {
		FileInfoCounter.MaxFileSize = req.FileSize
	}
	FileInfoCounter.AvgFileSize = (FileInfoCounter.AvgFileSize + req.FileSize) / int64(FileInfoCounter.FileCount)
	FileInfoCounter.FileExtension[req.FileExt] = FileInfoCounter.FileExtension[req.FileExt] + 1
	max := 0
	var key string
	for k, v := range FileInfoCounter.FileExtension {
		if v > max {
			max = v
			key = k
		}
	}
	FileInfoCounter.FreqFileExtension = key + fmt.Sprintf("=%v", max)
	FileInfoCounter.RecentFilePaths[FileInfoCounter.FileCount%10] = req.Path

	w.WriteHeader(http.StatusOK)
}
