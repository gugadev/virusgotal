package requests

import (
	"log"
	"os"

	"github.com/gugadev/virusgotal/models"
	"github.com/levigross/grequests"
)

const uri = "https://www.virustotal.com/vtapi/v2/file/"

/* type ProgressReader struct {
	io.Reader
	Reporter func(r int64)
}

func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	n, err = pr.Reader.Read(p)
	pr.Reporter(int64(n))
	return
} */

/*
Upload  send a file to VirusTotal
*/
func Upload(key string, file *os.File) models.Scan {
	var result models.Scan
	// get info of file
	fileInfo, _ := file.Stat()

	// prepare payload
	ro := &grequests.RequestOptions{
		Data: map[string]string{
			"apikey": key,
		},
		Files: []grequests.FileUpload{
			{
				FileName:     fileInfo.Name(),
				FileContents: file,
				FieldName:    "file",
			},
		},
	}
	response, err := grequests.Post(uri+"scan", ro)
	if err != nil {
		log.Fatal(err)
	}
	response.JSON(&result)
	return result
}

/*
GetReport get the report of the
previously uploaded/scanned file
*/
func GetReport(key, resource string) models.Report {
	var result models.Report
	ro := &grequests.RequestOptions{
		Params: map[string]string{
			"apikey":   key,
			"resource": resource,
		},
	}
	response, err := grequests.Get(uri+"report", ro)
	if err != nil {
		log.Fatal(err)
	}
	response.JSON(&result)
	return result
}
