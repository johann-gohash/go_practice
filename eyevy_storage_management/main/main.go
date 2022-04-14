package main 

import (
    "log"
    "os"
    "strings"
    "errors"
    "github.com/joho/godotenv"
    "hourly_upload"
)

func lsCaptureDir() (ret [] string) {

    files, err := os.ReadDir("../captures")

    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        if strings.HasSuffix(file.Name(), ".jpeg" ) || strings.HasSuffix(file.Name(), ".jpg" ) {
            ret = append(ret, file.Name())
        }
    }
    return
}

func loadEnvVars() {

    err := godotenv.Load("../.env")

    if err != nil {
        log.Fatal("Error loading .env file.")
    }
}


func checkArgs() (mode string, err error) {
    // First, we expect one arg to either be hourly_cap or mp4_compile.

    if len(os.Args) != 2 {
        return "Bad arg", errors.New("Please provide one argument. [hourly_cap or mp4_compile]")
    } else if strings.Compare(os.Args[1], "hourly_cap") == 0 || strings.Compare(os.Args[1], "mp4_compile") == 0 {
        // Strings.Compare returns 0 if there is a match,
        // So if one of the statements above is 0, that means
        // the provided argument is valid.
        mode = os.Args[1]
    } else {
        return "Bad arg", errors.New("Argument not recognized")
    }

    return mode, nil
}

func main() {

    mode, err := checkArgs()
    loadEnvVars()

    if err != nil {
        log.Fatal(err)
    }
    
    if strings.Compare(mode, "hourly_cap") == 0 {
        log.Printf("Mode is: %s\n", mode)
        hourly_upload.HourlyUpload(lsCaptureDir(), 
                                   os.Getenv("ACCESS_TOKEN"),
                                   os.Getenv("DISCORD_WEBHOOK_URL"))
    } else {
        // Assuming check_args( ) didn't return an error,
        // the only other possible value for mode is "mp4_compile"
        log.Printf("Mode is: %s\n", mode)
    }
}
