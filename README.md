This is a function for parsing dql strings into js objects. It is a wrapper over dgraph's own dql parser, compiled from Go to C with a binding to Node.

It should run on Node>=10 in an environment with GLIBC>=2.23, x86_64.

## usage

```js
const { parse } = require('@holographio/dql-parse')

parse(
	`
		query all($some_value: int) {
			all(func: eq(some_field, $some_value)) {
				foo
				bar
			}
		}
	`,
	{ $some_value: '123' }
)
```
