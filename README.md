# GO Types

Generate typescript types from Go struct for easy interop between Go and Typescript.

## What is this?

Using Go for the backend is the dream. Yet, when you *have* to use Typescript based frameworks for your frontend, you end up writing the same types in both languages. This is a waste of time and error prone.

So, this project aims to provide a way to generate Typescript types from Go types.

Now, I could have made a cli that actually parsed the Go code and generated the types, but that would have been a lot of work. Instead, go comes with the handy `reflect` package that allows us to get the type of a variable at runtime. This means that we can just write a simple function that takes a variable and returns a string that is the Typescript type.

So, in a way this is a Typescript type generator that uses Go as a DSL.

## I need your help

I am not a Typescript expert. I have only used it for a few months. So, if you see something that can be improved, please open an issue or a PR.

Moreover, if you have any suggestions as to how I can better structure this project, handle the transpilation, etc. please let me know.

This is the first time I am writing something so wacky, so I am sure there are a lot of things that can be improved.

## Usage

```go
package main

import (
    "fmt"

    gotypes "github.com/newtoallofthis123/gotypes"
)

type User struct {
    Name string
    Age  int
}

func main() {
    user := User{
        Name: "John Doe",
        Age:  20,
    }

    fmt.Println(gotypes.GenerateType("user.ts", user))
}
```

This will generate `user.ts` with the following contents:

```ts
export type User {
    name: string;
    age: number;
}
```

By default, gotypes will infer the typescript name from the Go type name. If you want to override this, you can use the `json` tag:

```go
type User struct {
    Name string `json:"username"`
    Age  int
}
```

This will print out the following:

```ts
export type User {
    username: string;
    age: number;
}
```

Cool right?

## Supported Types

For now, gotypes supports the following typescript types:

- `string`
- `number`
- `bool`

Along with their array counterparts:

- `string[]`
- `number[]`
- `bool[]`

## MiniDoc

### `GenerateType`

```go
func GenerateType(filename string, v interface{}) string
```

This function takes a variable and returns a string that is the Typescript type.

### `GenerateTypes`

```go
func GenerateTypes(filename string, v interface{}) string
```

This function takes a variable and returns a string that is the Typescript type. It also generates the types for all the fields of the variable.

### More

There are a lot more small functions that are used internally.

For more information, check out the [docs](https://pkg.go.dev/github.com/newtoallofthis123/gotypes).

## TODO

- [ ] Add support for more types
- [ ] Add support for nested structs
- [ ] Add support for nested arrays
- [ ] Add support for typescript enums
- [ ] Add support for typescript unions
- [ ] Add support for typescript interfaces

## Contributing

Contributions are welcome! Just open an issue or a PR.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
