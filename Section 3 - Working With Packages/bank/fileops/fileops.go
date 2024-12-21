package fileops
// new packages must be on your own folder
// the folder and the the package must have the same name
// the file itself does not need to have the same name

import (
    "errors"
    "fmt"
    "os"
    "strconv"
)

// Its only possible use (import) function/variables/constants from other packages if
// the name starts with an uppercase.
// Go will export everything that starts with and uppercase so it can be
// imported in other files
func GetFloatFromFile(fileName string) (value float64, err error) {
    data, errReadFile := os.ReadFile(fileName)
    if errReadFile != nil {
        err = errors.New("failed to find file")
    }

    value, errParseFloat := strconv.ParseFloat(string(data), 64)
    if errParseFloat != nil {
        err = errors.New("failed to parse stored value")
    }
    return value, err
}

func WriteFloatToFile(value float64, fileName string) {
    valueText := fmt.Sprint(value)
    os.WriteFile(fileName, []byte(valueText), 0644)
    // 0644 is a file permission notation where the owner of the file will be
    // able to write and read the file, whereas other users can only read it.
    // https://www.redhat.com/en/blog/linux-file-permissions-explained
}
