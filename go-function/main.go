package main

import (
    "github.com/nuclio/nuclio-sdk-go"
)

func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
    return "Hello from reference: refs/heads/go-func", nil
}

