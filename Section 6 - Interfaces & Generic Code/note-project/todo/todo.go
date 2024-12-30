package todo
import (
    "fmt"
    "errors"
    "os"
    "encoding/json"
)

type Todo struct {
    /*
    Metadata added to other packages (in this case the json) be able to read and
    do something with it.
    In this scenario will replace the field`s name with the word between "" when encoding.
    */
    Text string `json:"text"`
}

func New(content string) (Todo, error) {
    if content=="" {
        return Todo{}, errors.New("invalid input")
    }
    return Todo{
        Text: content,
    }, nil
}
func (t Todo) Print() {
    fmt.Printf("\nText: %v\n", t.Text)
}

func (t Todo) Save() error {
    fileName := "todo.json"
    json, err := json.Marshal(t)
    if err!=nil {
        return err
    }
    return os.WriteFile(fileName, json, 0644)
}