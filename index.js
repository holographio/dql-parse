const { parse } = require('./build/Release/dql_parse')

module.exports = {
	parse: (query, variables) => JSON.parse(parse(query, JSON.stringify(variables)))
}
