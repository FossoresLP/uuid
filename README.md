# UUID for Go

[![Licensed under: Boost Software License](https://img.shields.io/github/license/FossoresLP/uuid?color=28D)](https://github.com/FossoresLP/uuid/blob/main/LICENSE)
[![Current release](https://img.shields.io/github/v/release/FossoresLP/uuid?display_name=tag&sort=semver)](https://github.com/FossoresLP/uuid/releases)
[![Go Version](https://img.shields.io/badge/go-1.24+-00ADD8.svg)](https://go.dev/dl/)
[![Documentation](https://img.shields.io/badge/Docs-pkg.go.dev-blue)](https://pkg.go.dev/github.com/fossoreslp/uuid)
[![Go Report Card](https://goreportcard.com/badge/github.com/FossoresLP/uuid)](https://goreportcard.com/report/github.com/FossoresLP/uuid)

This Go package provides tools for generating and parsing Universally Unique Identifiers (UUIDs) according to the **[RFC 9562](https://www.ietf.org/rfc/rfc9562.html)** standard, which obsoletes RFC 4122.

## Key Features

*   **Full RFC 9562 Implementation:** Supports UUID versions 1, 3, 4, 5, 6, 7, and 8.
*   **Thread-Safe:** Generation of time-based UUIDs (V1, V6, V7) is thread-safe.
*   **Sortable UUIDs:** V6 and V7 provide time-sortable UUIDs, ideal for database keys. V7 is generally recommended for new applications.
*   **High-Precision V7:** Version 7 implementation uses millisecond timestamp precision plus additional fractional bits for better ordering within the same millisecond.
*   **Configurable V1/V6 MAC:** Use system hardware MAC, a custom MAC, or the default randomly generated MAC address for V1 and V6 UUIDs.
*   **Robust Parsing:** Parse canonical string representation or raw binary bytes.
*   **Standard Interface Support:** Natively implements `fmt.Stringer`, `encoding.Text(Un)Marshaler`, `encoding.Binary(Un)Marshaler`, `database/sql.Scanner`, and `database/sql/driver.Valuer` for seamless integration.

## Installation

```bash
go get github.com/fossoreslp/uuid@v1.0.0
```

Requires **Go 1.24** or later.

## Usage

### Generating UUIDs

```go
package main

import (
	"fmt"
	"github.com/fossoreslp/uuid"
)

func main() {
	// Version 4 (Random)
	idV4 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", idV4)

	// Version 7 (Timestamp + Random, Recommended Sortable)
	idV7 := uuid.NewV7()
	fmt.Printf("UUIDv7: %s\n", idV7)

	// Version 1 (Timestamp + MAC)
	idV1 := uuid.NewV1()
	fmt.Printf("UUIDv1: %s\n", idV1)

	// Version 6 (Timestamp + MAC, Sortable)
	idV6 := uuid.NewV6()
	fmt.Printf("UUIDv6: %s\n", idV6)

	// Version 3 (MD5 Hash)
	idV3 := uuid.NewV3(uuid.NamespaceDNS(), "example.com")
	fmt.Printf("UUIDv3: %s\n", idV3)

	// Version 5 (SHA1 Hash)
	idV5 := uuid.NewV5(uuid.NamespaceDNS(), "example.com")
	fmt.Printf("UUIDv5: %s\n", idV5)

	// Version 8 (Custom Data)
	customData := []byte{
		0xDE, 0xAD, 0xBE, 0xEF, 0xCA, 0xFE, 0xBA, 0xBE,
		0xFE, 0xED, 0xFA, 0xCE, 0xBA, 0xAD, 0xF0, 0x0D,
	}
	idV8 := uuid.NewV8(customData)
	fmt.Printf("UUIDv8: %s\n", idV8) // Note: Version/Variant bits are overwritten
}
```

### Parsing UUIDs

```go
package main

import (
	"fmt"
	"github.com/fossoreslp/uuid"
)

func main() {
	// Parse from string
	s := "1ec9414c-232a-6b00-b3c8-9f6bdeced846" // A UUIDv6 example
	id, err := uuid.Parse(s)
	if err != nil {
		fmt.Printf("Error parsing string: %v\n", err)
	} else {
		fmt.Printf("Parsed from string: %s (Version %d)\n", id, id.Version())
	}

	// Parse from binary bytes (e.g., from database)
	binaryBytes := []byte{
		0x01, 0x7f, 0x22, 0xe2, 0x79, 0xb0, // Timestamp (ms)
		0x7c, 0xc3,                         // Version 7 + rand_a
		0x98, 0xc4, 0xdc, 0x0c, 0x0c, 0x07, 0x39, 0x8f, // Variant + rand_b
	} // A UUIDv7 example
	idFromBin, err := uuid.ParseBytes(binaryBytes)
	if err != nil {
		fmt.Printf("Error parsing binary bytes: %v\n", err)
	} else {
		fmt.Printf("Parsed from binary: %s (Version %d)\n", idFromBin, idFromBin.Version())
	}

	// Parse from string bytes (e.g., from database storing as CHAR(36))
	stringBytes := []byte("919108f7-52d1-4320-9bac-f847db4148a8") // A UUIDv4 example
	idFromStringBytes, err := uuid.ParseBytes(stringBytes)
	if err != nil {
		fmt.Printf("Error parsing string bytes: %v\n", err)
	} else {
		fmt.Printf("Parsed from string bytes: %s (Version %d)\n", idFromStringBytes, idFromStringBytes.Version())
	}
}
```

### Checking Special UUIDs

```go
package main

import (
	"fmt"
	"github.com/fossoreslp/uuid"
)

func main() {
	nilUUID := uuid.UUID{} // Zero value is the Nil UUID
	maxUUID := uuid.UUID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF} // All bits set to 1

	fmt.Printf("Is %s Nil? %t\n", nilUUID, nilUUID.IsNil()) // true
	fmt.Printf("Is %s Max? %t\n", maxUUID, maxUUID.IsMax()) // true

	idV4 := uuid.NewV4()
	fmt.Printf("Is %s Nil? %t\n", idV4, idV4.IsNil())       // false
	fmt.Printf("Is %s Max? %t\n", idV4, idV4.IsMax())       // false
}
```

### Using Interfaces

The built-in interface support makes common tasks easy:

```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/fossoreslp/uuid"
	_ "github.com/mattn/go-sqlite3" // Example DB driver
)

type Record struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

func main() {
	id := uuid.NewV7()

	// fmt.Stringer (used by Println, etc.)
	fmt.Printf("UUID: %s\n", id)

	// encoding.Text(Un)Marshaler / encoding.Binary(Un)Marshaler (e.g., JSON)
	rec := Record{ID: id, Name: "Example"}
	jsonData, _ := json.Marshal(rec)
	fmt.Printf("JSON: %s\n", jsonData) // Output: {"id":"01...","name":"Example"}

	var decodedRec Record
	json.Unmarshal(jsonData, &decodedRec)
	fmt.Printf("Decoded ID: %s\n", decodedRec.ID)

	// database/sql.Scanner / driver.Valuer (Database interaction)
	// Assuming a DB connection `db *sql.DB` and table `records (id BLOB PRIMARY KEY, name TEXT)`
	// db, _ := sql.Open("sqlite3", ":memory:")
	// db.Exec("CREATE TABLE records (id BLOB PRIMARY KEY, name TEXT)")
	// _, err := db.Exec("INSERT INTO records (id, name) VALUES (?, ?)", id, "DB Example")
	// if err == nil {
	// 	var dbRec Record
	// 	row := db.QueryRow("SELECT id, name FROM records WHERE id = ?", id)
	// 	err = row.Scan(&dbRec.ID, &dbRec.Name) // Scan automatically handles uuid.UUID
	// 	if err == nil {
	// 		fmt.Printf("Read from DB: ID=%s, Name=%s\n", dbRec.ID, dbRec.Name)
	// 	}
	// }
}
```

### Configuring V1/V6 MAC Address

*Call these functions early in your application initialization, before generating V1 or V6 UUIDs.* They are not thread-safe during configuration.

```go
package main

import (
	"fmt"
	"github.com/fossoreslp/uuid"
	"log"
	"net"
)

func main() {
	// Option 1: Try to use a hardware MAC address from the system
	err := uuid.UseHardwareMAC()
	if err != nil {
		log.Printf("Could not set hardware MAC, using random: %v", err)
		// Default random MAC will be used
	} else {
		fmt.Println("Using hardware MAC address for V1/V6.")
	}

	// Option 2: Set a specific MAC address
	// customMAC, _ := net.ParseMAC("01:02:03:04:05:06")
	// err = uuid.SetMACAddress(customMAC)
	// if err != nil {
	//     log.Fatalf("Failed to set custom MAC: %v", err)
	// }
	// fmt.Println("Using custom MAC address for V1/V6.")

	// Now generate V1 or V6 UUIDs
	idV6 := uuid.NewV6()
	fmt.Printf("Generated V6 with configured MAC: %s\n", idV6)
}
```

## UUID Versions Overview

*   **Version 1 (Timestamp, MAC):** Based on current time and a node MAC address. Time component order is not suitable for direct sorting.
*   **Version 3 (Name-Based, MD5):** Generated by hashing a namespace UUID and a name using MD5.
*   **Version 4 (Random):** Generated from cryptographically secure random numbers. Most common version when sortability is not needed.
*   **Version 5 (Name-Based, SHA-1):** Generated by hashing a namespace UUID and a name using SHA-1. Preferred over V3.
*   **Version 6 (Reordered Timestamp, MAC):** Like V1 but with time bits rearranged for sortability. A sortable alternative if V1 compatibility/semantics are needed.
*   **Version 7 (Unix Epoch Timestamp, Random):** Combines a high-precision Unix timestamp with random data. Recommended for new applications needing time-sortable, collision-resistant IDs without exposing a MAC address.
*   **Version 8 (Custom/Experimental):** Allows custom data layout, defined by RFC 9562 for experimental or vendor-specific use.

## License

This package is licensed under the Boost Software License 1.0. See the [LICENSE](https://github.com/FossoresLP/uuid/blob/main/LICENSE) file for details.
