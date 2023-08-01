package main

import (
	"io"
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"net/http"
	"net/http/httptest"
	"testing"
	"mime/multipart"
)

/* UNIT TESTOVI */

func generateFakeJPEG() ([]byte, error) {
	// Pravimo sliku dimenzija 200x200 sa belom pozadinom
	width, height := 200, 200
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	// Kreiramo bafer za slike
	var buf bytes.Buffer
	// Kodiramo sliku u JPEG format
	if err := jpeg.Encode(&buf, img, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func TestUploadFile(t *testing.T) {
	// Generišemo lažni JPEG sadržaj
	fakeJPEG, err := generateFakeJPEG()
	fakeJPEGReader := bytes.NewReader(fakeJPEG)
	if err != nil {
		t.Fatal(err)
	}
	// Kreiramo lažni HTTP zahtev sa lažnom slikom
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	defer writer.Close()

	// Dodajemo lažni fajl sa imenom "photo" i MIME tipom "image/jpeg"
	part, err := writer.CreateFormFile("photo", "fake_image.jpg")
	if err != nil {
		t.Fatal(err)
	}

	// Kopiramo sadržaj lažne slike u deo zahteva
	if _, err := io.Copy(part, fakeJPEGReader); err != nil {
		t.Fatal(err)
	}
	// Završavamo kreiranje zahteva
	writer.Close()

	// Pravimo HTTP zahtev
	req, err := http.NewRequest("POST", "/upload", &requestBody)
	if err != nil {
		t.Fatal(err)
	}
	// Postavljamo odgovarajući Content-Type header
	req.Header.Set("Content-Type", writer.FormDataContentType())

    rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uploadFile)
	handler.ServeHTTP(rr, req)


    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status %v, but got %v", http.StatusOK, status)
    }
	expectedResponse := "File upload endpoint hit!\n"+"Successfully uploaded the file!\n"
	if body := rr.Body.String(); body != expectedResponse {
		t.Errorf("Expected response '%v', but got '%v'", expectedResponse, body)
	}
}