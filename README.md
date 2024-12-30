# Moderation Package for Go

## Overview
The **moderation** package provides a simple and efficient client for interacting with the OpenAI Moderation API. It enables developers to send text inputs for moderation and receive categorized results with scores.

## Features
- Sends requests to the OpenAI Moderation API.
- Handles API authentication and request/response parsing.
- Provides a reusable and extensible design.

## Installation

1. Add the package to your project:
   ```bash
   go get github.com/your_username/moderation
   ```

2. Import the package into your code:
   ```go
   import "github.com/your_username/moderation"
   ```

## Usage

### Setup

Create a session and initialize the client:

```go
package main

import (
	"context"
	"fmt"
	"github.com/your_username/moderation"
)

func main() {
	session := &moderation.Session{APIKey: "your_openai_api_key"}
	client := moderation.NewClient(session, "omni-moderation-latest")

	ctx := context.Background()
	request := &moderation.Request{
		Input: "This is a test input for moderation."
	}

	response, err := client.Create(ctx, request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Response: %+v\n", response)
}
```

### API Documentation

#### `NewClient`
Creates a new moderation client.

**Parameters:**
- `session *Session`: The API session containing your API key.
- `model string`: The model to use for moderation (e.g., `"omni-moderation-latest"`).

**Returns:**
- `*Client`: The moderation client instance.

#### `Client.Create`
Sends a moderation request and returns the response.

**Parameters:**
- `ctx context.Context`: Context for the API call.
- `p *Request`: The moderation request containing the input and model.

**Returns:**
- `*Response`: The moderation API response.
- `error`: An error if the request fails.

#### `Request`
Represents the moderation request payload.

**Fields:**
- `Model string`: The model to use for moderation.
- `Input string`: The text input for moderation.

#### `Response`
Represents the moderation API response.

**Fields:**
- `ID string`: The unique ID of the response.
- `Model string`: The model used for moderation.
- `Results []Result`: A list of moderation results.

#### `Result`
Represents a single moderation result.

**Fields:**
- `Categories map[string]bool`: Categorized results (e.g., `harassment`, `hate`).
- `CategoryScores map[string]float64`: Confidence scores for each category.
- `Flagged bool`: Whether the input was flagged.

## Testing

### Writing Tests
Create a test file `moderation_test.go`:

```go
package moderation

import (
	"context"
	"testing"
)

func TestCreateModeration(t *testing.T) {
	session := &Session{APIKey: "test_api_key"}
	client := NewClient(session, "test-model")

	ctx := context.Background()
	request := &Request{
		Input: "Test moderation input",
	}

	// Mock the MakeRequest method in a real-world scenario
	response, err := client.Create(ctx, request)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response == nil || response.ID == "" {
		t.Fatalf("expected a valid response, got: %+v", response)
	}
}
```

### Running Tests
Run the tests using the following command:

```bash
go test ./...
```

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests to improve the package.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

