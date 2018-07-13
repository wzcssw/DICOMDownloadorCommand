package main

import (
	"DICOMDownloadorCommand/lib"
	"encoding/json"
	"fmt"
	"os"
)

const DICOMServerURL string = "http://dicomup.tongxinyiliao.com/api/getByFilmNo"

var con int = 30

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请输入影像号")
		return
	}
	m := make(map[string]string)
	m["filmno"] = os.Args[1]
	responsedDataBody := lib.SendDicomAPIRequest(DICOMServerURL, m)
	var responsedData lib.DicomAPIRequest
	json.Unmarshal([]byte(responsedDataBody), &responsedData)
	if len(responsedData.List) < 1 {
		fmt.Println("No DICOM Files exist.")
		return
	}
	go lib.ShowProcess(lib.CountSeriesFile(responsedData.List))

	fmt.Println("FilmNo:", m["filmno"], "        DICOM Files:", lib.CountSeriesFile(responsedData.List))
	lib.DownloadSeriesFile(responsedData.List, m["filmno"], con)

}
