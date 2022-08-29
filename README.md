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
import xml from 'k6/x/xml';

export default function () {
    const body = `
    <?xml version="1.0"?>
    <catalog>
        <book id="bk101">
            <author>Gambardella, Matthew</author>
            <title>XML Developer's Guide</title>
            <genre>Computer</genre>
            <price>44.95</price>
            <publish_date>2000-10-01</publish_date>
            <description>
                <![CDATA[An in-depth look at creating applications with XML.]]>
            </description>
        </book>
        <book id="bk102">
            <author>Ralls, Kim</author>
            <title>Midnight Rain</title>
            <genre>Fantasy</genre>
            <price>5.95</price>
            <publish_date>2000-12-16</publish_date>
            <description>
                <![CDATA[A former architect battles corporate zombies, an evil sorceress, and her own childhood to become queen of the world.]]>
            </description>
        </book>
    </catalog>
    `

    const result = xml.parse(body);

    result["catalog"]["book"].forEach(book => {
        console.log(book["-id"], book["title"]);
    });
}
```