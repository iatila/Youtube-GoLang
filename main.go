package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/download", downloadHandler)
	fmt.Println("Sunucu başlatıldı: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, nil)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		videoURL := r.FormValue("videoURL")
		if videoURL == "" {
			http.Error(w, "Lütfen bir URL girin", http.StatusBadRequest)
			return
		}
		ytDlpPath := "C:\\bin\\yt-dlp.exe"

		// YouTube videosunun başlığını alın
		cmd := exec.Command(ytDlpPath, "--get-title", videoURL)
		output, err := cmd.Output()
		if err != nil {
			http.Error(w, "Video başlığı alınamadı", http.StatusInternalServerError)
			return
		}

		// Başlığı temizleyin
		videoTitle := strings.TrimSpace(string(output))
		cleanTitle := cleanFileName(videoTitle)

		// Temiz başlığı kullanarak dosya adı oluşturun
		videoFileName := fmt.Sprintf("%s.mp4", cleanTitle)

		// Videoyu temizlenmiş adla indirin
		cmd = exec.Command(ytDlpPath, "-f", "mp4", "-o", videoFileName, videoURL)
		err = cmd.Run()
		if err != nil {
			writeJSONResponse(w, http.StatusInternalServerError, "Video indirilirken bir hata oluştu")
			log.Println("İndirme hatası:", err)
			return
		}

		// MP3 olarak dönüştür ve music klasörüne kaydet
		ffmpegPath := "C:\\bin\\ffmpeg.exe"
		mp3FileName := filepath.Join("music", fmt.Sprintf("%s.mp3", cleanTitle))
		cmd = exec.Command(ffmpegPath, "-i", videoFileName, mp3FileName)
		output, err = cmd.CombinedOutput()
		if err != nil {
			writeJSONResponse(w, http.StatusInternalServerError, "MP3 dönüştürme işlemi başarısız oldu")
			log.Println("Dönüştürme hatası:", err)
			log.Println("ffmpeg çıktısı:", string(output))
			return
		}

		// Video dosyasını sil
		err = os.Remove(videoFileName)
		if err != nil {
			writeJSONResponse(w, http.StatusInternalServerError, "Video dosyası silinemedi")
			log.Println("Video silme hatası:", err)
			return
		}

		writeJSONResponse(w, http.StatusOK, "Video başarıyla indirildi, MP3'e dönüştürüldü ve video dosyası silindi!")
	}
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func cleanFileName(fileName string) string {
	// Dosya adını küçük harfe çevir
	fileName = strings.ToLower(fileName)

	// Harf ve sayılar dışındaki karakterleri "_" ile değiştir
	re := regexp.MustCompile("[^a-z0-9]+")
	cleanName := re.ReplaceAllStringFunc(fileName, func(s string) string {
		// Eğer bir karakter boşluksa, bir adet "_" ekle
		if unicode.IsSpace([]rune(s)[0]) {
			return "_"
		}
		return "_"
	})

	return cleanName
}
