package main 

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"

)

func ls_capture_dir() {

    files, err := os.ReadDir("captures")

    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }

}

func load_env_vars() {

    fmt.Println("Testing .env module...")
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file.")
    }

}

func main() {


    load_env_vars()

    access_token_key := os.Getenv("ACCESS_TOKEN")
    mp4_compiles_album := os.Getenv("MP4_COMPILES_ALBUM_ID")
    hourly_compiles_album := os.Getenv("HOURLY_CAPTURES_ALBUM_ID")
     

    fmt.Println("Access Token Key: ", access_token_key)
    fmt.Println(mp4_compiles_album)
    fmt.Println(hourly_compiles_album)


    ls_capture_dir()
}
