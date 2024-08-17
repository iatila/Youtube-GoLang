# YouTube Video Downloader & MP3 Converter
## Golang pratikleri
Bu proje, bir YouTube video URL'si sağlayarak videoyu indirmenizi ve MP3 formatına dönüştürmenizi sağlar. Proje, Go dili ile geliştirilmiştir ve `yt-dlp` ve `ffmpeg` araçlarını kullanır.

## Özellikler

- YouTube videolarını MP4 formatında indirir.
- İndirilen videoyu MP3 formatına dönüştürür.
- İşlem sırasında kullanıcıya bilgi veren uyarılar ve animasyonlar sağlar.
- Bootstrap 5 ile şık bir kullanıcı arayüzü sunar.

## Kurulum

### Gerekli Araçlar

1. [Go](https://golang.org/dl/) - Go dilini kurun.
2. [yt-dlp](https://github.com/yt-dlp/yt-dlp) - YouTube videolarını indirmek için kullanılan araç. [İndir](https://github.com/yt-dlp/yt-dlp/releases)
3. [ffmpeg](https://ffmpeg.org/download.html) - Video ve ses dosyalarını dönüştürmek için kullanılan araç. [İndir](https://ffmpeg.org/download.html)

### Go Projesini Çalıştırma

1. **Go'yu kurun**: Go'yu [bu bağlantıdan](https://golang.org/dl/) indirin ve kurun.

2. **Proje Dosyalarını İndirin**:
   ```bash
   git clone [[https://github.com/iatila/Youtube-GoLang.git]
   cd Youtube-GoLang

