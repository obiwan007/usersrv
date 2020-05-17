// ReactQL example page - delete this folder for your own project!

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */

import React from "react";
import Home from "./home";

// ----------------------------------------------------------------------------

interface IIndexState {
  dynamic: React.SFC | null;
}

interface IProps {
  history: any;
}
// Say hello from GraphQL, along with a HackerNews feed fetched by GraphQL
class Index extends React.PureComponent<IProps, IIndexState> {
  public state = {
    dynamic: null,
  };

  public componentDidMount = async () => {};

  public render() {
    return (
      <>
        <Home />
      </>
    );
  }
}

export default Index;
