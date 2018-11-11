package net

import (
  "net/http"
  "io"
  "os"
  "path/filepath"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}

func GetFileContentType(out *os.File) (string, error) {

    // Only the first 512 bytes are used to sniff the content type.
    buffer := make([]byte, 512)

    _, err := out.Read(buffer)
    if err != nil {
        return "", err
    }

    // Use the net/http package's handy DectectContentType function. Always returns a valid
    // content-type by returning "application/octet-stream" if no others seemed to match.
    contentType := http.DetectContentType(buffer)

    return contentType, nil
}

func GetFileExt(filename string) string {

    return filepath.Ext(filename)
    // name := strings.TrimSuffix(basename, filepath.Ext(basename))
    // return filename[len(extension):len(filename)]
}
