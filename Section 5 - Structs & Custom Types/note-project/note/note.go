package note
import (
    "fmt"
    "time"
    "errors"
    "os"
    "strings"
    "encoding/json"
)

type Note struct {
    /*
    Metadata added to other packages (in this case the json) be able to read and
    do something with it.
    In this scenario will replace the field`s name with the word between "" when encoding.
    */
    Title string `json:"title"`
    Content string `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}

func New(t, c string) (Note, error) {
    if t=="" || c=="" {
        return Note{}, errors.New("invalid input")
    }
    return Note{
        Title: t,
        Content: c,
        CreatedAt: time.Now(),
    }, nil
}
func (n Note) Print() {
    fmt.Printf("\nTitle: %v\nContent: %v\n", n.Title, n.Content)
}

func (n Note) Save() error {
    fileName := strings.ReplaceAll(n.Title, " ", "-")
    fileName = strings.ToLower(fileName) + ".json"
    
    json, err := json.Marshal(n)
    if err!=nil {
        return err
    }
    return os.WriteFile(fileName, json, 0644)
}