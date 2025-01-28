# Antimetal APIs

This repository contains interface definitions as well as generated Go libraries for public Antimetal APIs.

The generated libraries and clients are primarily used in other public Antimetal repositories.
But you can also use these definitions to generate your own client libraries and other artifacts for your own purposes.

The interface definitions are defined using [Protocol Buffers](https://protobuf.dev/) verison 3 (proto3).
We use Protocol Buffers to generate gRPC versions of our APIs.
We use [Buf](https://buf.build/) to manage the Protobuf files and generate source code.

## Repository Structure

`antimetal` is the root directory for all Antimetal APIs where each API is its own subdirectory.
Each API directory then has subdirectories for each major version of the API.
So for example, the path for `foo` API is `antimetal/foo/v1`. A v2 version would be `antimetal/foo/v2`.

The proto package names match this directory structure.
The Go package isn't defined in the protobuf file; this is handled by buf, which automatically adds `github.com/antimetal/apis` as the package prefix when generating the Go libraries.
This makes it easy to locate the proto definitions for a particular API and ensures that the Go import paths are correct.

## Building Go Libraries

To generate Go libaries from the Protobuf interface definitions, simple run:

```
make proto
```

This will generate the Go source files within the Protobuf directories using Buf.
