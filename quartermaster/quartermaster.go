package quartermaster

import (
    "fmt"
    "github.com/attamusc/quartermaster/config"
    "os"
)

const configName string = "quartermaster.json"

func Install() {
    wd, _ := os.Getwd()
    fmt.Println(wd)
}

func Bundle() {
    configFile, err := config.Read()

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(configFile)

    configFile.InstallDependencies()
}
