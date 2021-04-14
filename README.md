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
