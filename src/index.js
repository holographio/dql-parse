const { parse } = require('../build/Release/dql_parse')

module.exports = {
	parse: (query, variables = {}) => {
		const { error, ...rest } = JSON.parse(parse(query, JSON.stringify(variables)))
		if (error) { throw error }
		return rest
	}
}
