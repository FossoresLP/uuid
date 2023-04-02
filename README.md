UUID
====

[![Licensed under: Boost Software License](https://img.shields.io/github/license/FossoresLP/uuid?color=28D)](https://github.com/FossoresLP/uuid/blob/main/LICENSE)
[![Current release](https://img.shields.io/github/v/release/FossoresLP/uuid?display_name=tag&sort=semver)](https://github.com/FossoresLP/uuid/releases)
[![Documentation](https://img.shields.io/badge/Docs-pkg.go.dev-blue)](https://pkg.go.dev/github.com/fossoreslp/uuid)

This package implements the UUID versions defined in [RFC4122bis](https://www.ietf.org/archive/id/draft-ietf-uuidrev-rfc4122bis-02.html).

Versions 1 through 5 are defined in the original [RFC4122](https://www.ietf.org/rfc/rfc4122.html), while versions 6 through 8 are new additions.

Versions 6 and 7 are designed as timestamp-based, sortable IDs for use e.g. in database indexes. Version 7 is recommended for new applications while version 6 is a sortable replacement for version 1.

Version 8 is highly customizable, allowing for 122 custom bits with only version and variant being predefined.

Create a new UUID: `uuid.NewVx() (UUID, error)`

Convert an UUID to a string: `UUID.String() string`

Convert a string to an UUID: `uuid.Parse(string) (UUID, error)`

Convert a byte slice to an UUID: `uuid.ParseBytes([]byte) (UUID, error)`

Check if UUID contains only zeros: `UUID.IsNil() bool`

Check if UUID contains only ones: `UUID.IsMax() bool`

Support for `encoding.Text(Un)Marshaler`, `encoding.Binary(Un)Marshaler`, `database/sql.Scanner` and `database/sql/driver.Valuer` is built in, so the IDs can be used with most data exchange formats and databases.
