package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"sort"
	"path/filepath"
)

type Photo struct {
	Title	string	`json:"title"`
	Hex		string	`json:"hex"`
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "File upload endpoint hit!\n")

	r.ParseMultipartForm(1 << 30)

	file, handler, err := r.FormFile("photo")
	if err != nil {
		fmt.Println("Error Retrieveing the File")
		fmt.Println(err)
		return
	}
	defer file.Close()


	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fileType := http.DetectContentType(fileBytes)
	
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	fmt.Printf("FILE TYPE!!: %+v\n", fileType)

	if fileType != "image/png" && fileType != "image/jpeg" {
		fmt.Fprintf(w, "File can not be uploaded!\n")
		fmt.Fprintf(w, "File must be of png or jpeg type!\n")
		return
	}
	fileExt := filepath.Ext(handler.Filename)

	h := sha1.New()
	h.Write([]byte(handler.Filename))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	tempFileStr := sha1_hash+fileExt

	m := make(map[string]int)
	var photos = getPhotoTitles()
	var photo = Photo{Title: handler.Filename, Hex: sha1_hash}

	for i := 0; i < len(photos); i++ {
		m[photos[i].Hex] = 1
	}

	val := m[sha1_hash]

	if val == 1 {
		fmt.Fprintf(w, "That photo was already uploaded!\n")
		return 
	}

	photos = append(photos, photo)
	data, _ := json.Marshal(photos)
	ioutil.WriteFile("photo_titles.json", data, 0644)


	tempFile, err := os.Create(fmt.Sprintf("./photos/%s", tempFileStr))
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully uploaded the file!\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/list", listPhotoTitles)
	http.HandleFunc("/list/sort", listSortedPhotoTitles)
	http.HandleFunc("/delete/", deleteAPhoto)
	
	http.HandleFunc("/download/", downloadAPhoto)
	http.ListenAndServe(":8080", nil)
}

func downloadAPhoto(w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	h.Write([]byte(r.URL.Path[10:]))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	fileExt := filepath.Ext(r.URL.Path[10:])
	tempFileStr := sha1_hash+fileExt
	fmt.Println("downloadAPhoto: ", "./photos/"+tempFileStr)
	fileBytes, err := ioutil.ReadFile("./photos/"+tempFileStr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Photo with such name does not exist!\n")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func remove(s []Photo, i int) []Photo {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func deleteAPhoto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE A PHOTO! ", r.URL.Path)
	var slug string = r.URL.Path[len("/delete/"):]
	fmt.Println("DELETE A PHOTO! SLUG: ", slug)
	fileExt := filepath.Ext(r.URL.Path[8:])

	var photos = getPhotoTitles()
	for i := 0; i < len(photos); i++ {
		if slug == photos[i].Title{
			h := sha1.New()
			h.Write([]byte(slug))
			sha1_hash := hex.EncodeToString(h.Sum(nil))
			tempFileStr := sha1_hash+fileExt
			fmt.Println("DELETE AtempFileStrtempFileStr! : ", tempFileStr)

			err := os.Remove("./photos/"+tempFileStr)
			
			if err != nil {
			   fmt.Fprintln(w, "Error: ", err) 
			} else {
			   fmt.Fprintln(w, "Successfully deleted file: ", photos[i].Title)
			   photos = remove(photos, i)
			   data, _ := json.Marshal(photos)
			   ioutil.WriteFile("photo_titles.json", data, 0644)
			}
			return
		}
	}

	fmt.Fprintln(w, "Photo " + slug + " does not exist!")
}

func listPhotoTitles(w http.ResponseWriter, r *http.Request) {
	var photos = getPhotoTitles()

	for i := 0; i < len(photos); i++ {
		fmt.Fprintf(w, "Slika "+strconv.FormatInt(int64(i+1), 10)+": " + photos[i].Title + " \n")
	}
}

func listSortedPhotoTitles(w http.ResponseWriter, r *http.Request) {
	var photos = getPhotoTitles()

	A := make([]string, 0)
	for i := 0; i < len(photos); i++ {
		A = append(A, photos[i].Title)
	}

	sort.Strings(A)

	for i := 0; i < len(A); i++ {
		fmt.Fprintf(w, "Slika "+strconv.FormatInt(int64(i+1), 10)+": " + A[i] + " \n")
	}
}

func getPhotoTitles()[]Photo {
	jsonFile, err := os.Open("photo_titles.json")
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var photos []Photo

	json.Unmarshal(byteValue, &photos)
	return photos
}

func main() {
	fmt.Println("Server started!")
	setupRoutes()
}
