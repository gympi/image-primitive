package handlers

import (
  "net/http"
  "io"
  "log"
  "os"

  "github.com/gympi/image-primitive/libs/string_utils"
  "github.com/gympi/image-primitive/libs/net"
  "github.com/gympi/image-primitive/libs/primitive_constructor"
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

    download_filepath := "public/static/shared/" + string_utils.RandString(10) + net.GetFileExt(fileUrl)
    err := net.DownloadFile(download_filepath, fileUrl)
    if err != nil {
        panic(err)
    }

    filepath := primitive_constructor.PrimitiveConstructor(download_filepath, )


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
