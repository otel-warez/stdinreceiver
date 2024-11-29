# Stdin receiver

The Stdin receiver accepts logs from stdin on the current collector process.

Supported pipeline types: logs

> :construction: This receiver is in beta and configuration fields are subject to change.
## Configuration

Example:

```yaml
receivers:
  stdin:
```

## Standard in
The receiver consumes data passed in via standard input.

### Piping
If it receives data via pipe, the receiver consumes all data passed in, blocking until such time it sends it out.
It then stops.

### Interactive
If the collector is run in an interactive CLI, you can exit by entering enter, Ctrl+C or Ctrl+D.