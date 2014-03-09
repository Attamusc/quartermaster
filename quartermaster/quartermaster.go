package quartermaster

import (
    "fmt"
    "os"
)

func Install() {
    wd, _ := os.Getwd()
    fmt.Println(wd)

}

func Bundle() {
    wd, _ := os.Getwd()
    fmt.Println(wd)
}
