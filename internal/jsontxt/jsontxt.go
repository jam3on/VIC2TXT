package jsontxt

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Declaration of the structure of the JSON provided by the VIC database
type VicData struct {
	Context string `json:"@odata.context"`
	Value   []VIC  `json:"value"`
}

type VIC struct {
	MediaID            int       `json:"MediaID"`
	Category           int       `json:"Category"`
	MD5                string    `json:"MD5"`
	SHA1               string    `json:"SHA1,omitempty"`
	MediaSize          int       `json:"MediaSize"`
	DateUpdated        time.Time `json:"DateUpdated"`
	IsPrecategorized   bool      `json:"IsPrecategorized,string"`
	Exifs              []Exif    `json:"Exifs,omitempty"`
	Series             string    `json:"Series,omitempty"`
	OffenderIdentified bool      `json:"OffenderIdentified,string,omitempty"`
	VictimIdentified   bool      `json:"VictimIdentified,string,omitempty"`
	IsDistributed      bool      `json:"IsDistributed,string,omitempty"`
	PhotoDNA           string    `json:"PhotoDNA,omitempty"`
}

type Exif struct {
	MD5           string `json:"MD5"`
	PropertyName  string `json:"PropertyName"`
	PropertyValue string `json:"PropertyValue"`
}

// Reading the JSON
func readerJson(inFile string) []VIC {
	document, err := os.Open(inFile) // File to open
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer document.Close()

	var data VicData
	decoder := json.NewDecoder(document)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	return data.Value
}

// Function to write hashes to a file
func writeHashes(filename string, hashes <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating/opening file:", filename, err)
		return
	}
	defer file.Close()

	for hash := range hashes {
		_, err := file.WriteString(hash + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", filename, err)
		}
	}
}

// Function to create TXT files
func makeTxtFiles(outDir string, vics []VIC) {
	md5Chan := make(chan string, len(vics))
	sha1Chan := make(chan string, len(vics))
	var wg sync.WaitGroup

	fileMd5 := filepath.Join(time.Now().Format("20060102_150405") + "_md5_hashes.txt")
	fileSha1 := filepath.Join(time.Now().Format("20060102_150405") + "_sha1_hashes.txt")

	wg.Add(2)
	go writeHashes(filepath.Join(outDir, fileMd5), md5Chan, &wg)
	go writeHashes(filepath.Join(outDir, fileSha1), sha1Chan, &wg)

	for _, vic := range vics {
		if vic.MD5 != "" {
			md5Chan <- vic.MD5
		}
		if vic.SHA1 != "" {
			sha1Chan <- vic.SHA1
		}
	}

	close(md5Chan)
	close(sha1Chan)
	wg.Wait()
}

// Launch the module
func RunJsontoTxt(inFile string, outDir string) {
	vics := readerJson(inFile)
	if vics != nil {
		makeTxtFiles(outDir, vics)
	}
}
