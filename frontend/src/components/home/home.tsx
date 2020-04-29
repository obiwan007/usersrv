// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import React from "react";
import { withRouter } from "react-router-dom";
// ----------------------------------------------------------------------------


interface IProps {
  history?: any;
}
// Say hello from GraphQL, along with a HackerNews feed fetched by GraphQL
export class Home extends React.PureComponent<IProps, any> {
  render() {
    return (

      <h3>Home</h3>
    );
  }
}

export default withRouter(Home as any);
