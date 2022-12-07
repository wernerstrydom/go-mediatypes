# Media Types

This is a go library that contains a list of media types and their extensions. 

## Usage

```go
package main

import (
    "fmt"
    "github.com/wernerstrydom/go-mediatypes"
)

func main() {
    mediaTypes := mediatypes.ByExtension("gif")
    for _, mediaType := range mediaTypes {
        fmt.Println(mediaType)
    }
}
```
This produces

```text
image/gif
```

## Contributing

Most of the repository was generated using the [media types](https://github.com/wernerstrydom/mediatypes) project.
If you'd like to add, remove or update media types, or change how the source is
generated, please see the [media types](https://github.com/wernerstrydom/mediatypes) project.


## License

MIT

## Author

Werner Strydom

## Acknowledgements

This library is based on the following sources:

- [IANA Media Types](https://www.iana.org/assignments/media-types/media-types.xhtml)
- https://s-randomfiles.s3.amazonaws.com/mime/allMimeTypes.txt




