## How to create a new cmd verb

1. Create a dir/pkg w/ the new verb as the name
2. Create a type within the new package that implements the cmd.Cmd interface
3. Export an instance of this type
    - `daemon` & `status` verbs use the `var Cmd` as this exported instance
4. Register the new verbs package in `register.go`
