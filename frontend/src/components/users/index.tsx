// ReactQL example page - delete this folder for your own project!

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */

import React from "react";
import Users from "./users";

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
    dynamic: null
  };

  public componentDidMount = async () => {
    // Fetch the component dynamically
    const dynamic = await import("./dynamic");

    // ... and keep ahold of it locally
    this.setState({
      dynamic: dynamic.default
    });
  };

  public render() {
    const DynamicComponent = this.state.dynamic || (() => <h2>Loading...</h2>);

    return (
      <>
        {/* Note: The <h1> style will have a yellow background due to @/global/styles.ts! */}
        <h1>Hi from ReactQL</h1>
        <Users />
        {/* <DynamicComponent />
        <Count />
        <HackerNews />
        <button onClick={() => this.props.history.push("/login")}>Login</button> */}
      </>
    );
  }
}

export default Index;
