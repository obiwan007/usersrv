// Helper function that creates a new Apollo client per request
// Create browser history, for navigation a la single page apps
import { ApolloProvider as ApolloHooksProvider } from '@apollo/react-hooks';
import { createBrowserHistory } from "history";
import React from 'react';
// HOC for enabling Apollo GraphQL `<Query>` and `<Mutation>`
import { ApolloProvider } from "react-apollo";
import ReactDOM from 'react-dom';
import { Router } from "react-router-dom";
// Our main component, and the starting point for server/browser loading
import Root from "./components/root";
import './index.css';
import { createClient } from "./lib/apollo";
import * as serviceWorker from './serviceWorker';

// Create Apollo client
const client = createClient();
console.log("Apolleclient created", client);


function MyComponent() {
  return (
    <ApolloProvider client={client}>
      <ApolloHooksProvider client={client}>
        <Router history={history}>
          <Root />
        </Router>
      </ApolloHooksProvider>
    </ApolloProvider>
  )
}

// Create a browser history
const history = createBrowserHistory();
ReactDOM.render(
  // <React.StrictMode>
  <MyComponent />
  // </React.StrictMode>,
  , document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
