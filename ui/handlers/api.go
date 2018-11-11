package handlers

import (
  "net/http"
  "io"
  "log"
  "os"
  "math/rand"
  "time"
  "path/filepath"
  "os/exec"
)

func ApiPrimitive(w http.ResponseWriter, r *http.Request) {

    keys, ok := r.URL.Query()["key"]

    if !ok || len(keys[0]) < 1 {
        log.Println("Url Param 'key' is missing")
        return
    }

    // Query()["key"] will return an array of items,
    // we only want the single item.
    key := keys[0]

    log.Println("Url Param 'key' is: " + string(key))

    fileUrl := key

    download_filepath := "public/static/shared/" + RandString(10) + GetFileExt(fileUrl)
    err := DownloadFile(download_filepath, fileUrl)
    if err != nil {
        panic(err)
    }

    filepath := build_primitive(download_filepath, )


    img, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err) // perhaps handle this nicer
    }
    defer img.Close()
    w.Header().Set("Content-Type", "image/svg+xml") // <-- set the content-type header
    io.Copy(w, img)

    err_r := os.Remove(filepath)

     if err_r != nil {
         log.Println(err_r)
         return
     }

     err_r2 := os.Remove(download_filepath)

      if err_r2 != nil {
          log.Println(err_r2)
          return
      }
}


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


const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func RandString(length int) string {
  return StringWithCharset(length, charset)
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


// primitive console params
// https://github.com/fogleman/primitive
type PrimitiveApiParams struct {
 i string //input file
 o string //output file
 n int //number of shapes
 m int //mode: 0=combo, 1=triangle, 2=rect, 3=ellipse, 4=circle, 5=rotatedrect, 6=beziers, 7=rotatedellipse, 8=polygon
}

func build_primitive(in string) string {
    binary, lookErr := exec.LookPath("./vender/primitive")
    if lookErr != nil {
        log.Println(binary)
        panic(lookErr)
    }

    primitive_filepath := "public/static/shared/primitive/" + RandString(10) + ".svg"
    args := []string{"-i", in, "-o", primitive_filepath, "-n", "25"}

    cmd := "./vender/primitive"

    cmd_command := exec.Command(cmd, args...)
    cmd_command.Start()

    if err := cmd_command.Wait(); err != nil {
        log.Println(err)
    } else {
        log.Println("Importing complete")
    }

    return primitive_filepath
}
