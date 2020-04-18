// ReactQL example page - delete this folder for your own project!

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */

import React from "react";
import Security from "../../lib/security";
/* Local */

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

    // ... and keep ahold of it locally
    this.setState({});
  };

  public render() {
    // const DynamicComponent = this.state.dynamic || (() => <h2>Loading...</h2>);

    return (
      <>
        {/* Note: The <h1> style will have a yellow background due to @/global/styles.ts! */}
        <h1>Login</h1>
        <button onClick={() => this.loginClick()}>Login</button>
      </>
    );
  }

  async loginClick() {
    const ok = await Security.login("MyUsername", "MyPassword");
    console.log("Redirecting");
    if (ok) {
      const ok2 = await Security.login("MyUsername", "MyPassword");
      this.props.history.replace("/");
    }
  }
}

export default Index;
