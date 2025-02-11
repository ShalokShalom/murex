# `pt`

> Pipe telemetry. Writes data-types and bytes written

## Description

Pipe telemetry, `pt`, writes statistics about the pipeline. The telemetry is written
directly to the OS's STDERR so to preserved the pipeline.

## Usage

```
<stdin> -> pt -> <stdout>
```

## Examples

```
curl -s https://example.com/bigfile.bin -> pt -> > bigfile.bin
```

(though Murex does also have it's own HTTP clients, `get`, `post` and
`getfile`)

## See Also

* [`>>` (append file)](../commands/greater-than-greater-than.md):
  Writes STDIN to disk - appending contents if file already exists
* [`>` (truncate file)](../commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [`get`](../commands/get.md):
  Makes a standard HTTP request and returns the result as a JSON object
* [`getfile`](../commands/getfile.md):
  Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.
* [`post`](../commands/post.md):
  HTTP POST request with a JSON-parsable return

<hr/>

This document was generated from [builtins/core/io/file_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/io/file_doc.yaml).