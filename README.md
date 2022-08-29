# xk6-xml

This is a [k6](https://github.com/grafana/k6) extension using the
[xk6](https://github.com/grafana/xk6) system.

## Build

To build a `k6` binary with this plugin, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```shell
  go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  xk6 build --with github.com/rammstein4o/xk6-xml
  ```

## Example

```javascript
// script.js
import sql from 'k6/x/xml';

export default function () {
    
}
```