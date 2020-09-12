const { ApolloServer, gql } = require('apollo-server');

// The GraphQL schema
const typeDefs = gql`
  type Query {
    "A simple type for getting started!"
    hello: String
    error: String
  }
`;

// A map of functions which return data for the schema.
const resolvers = {
    Query: {
	hello: () => 'world',
	error: () => {
	    throw new Error('error')
	}
    },
};

const server = new ApolloServer({
  typeDefs,
  resolvers,
});

server.listen().then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});
