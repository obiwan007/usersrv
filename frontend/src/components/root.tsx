// Root entry point

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
// By default, pull in the ReactQL example. In your own project, just nix
// the `src/components/example` folder and replace the following line with
// your own React components
// Global styles
// import globalStyles from "@/global/styles";
// import { Global } from "@emotion/core";
import React from "react";
import { Helmet } from "react-helmet";
import { Route, Switch } from "react-router-dom";
import Example from "./example";
/* Local */
// Components
import ScrollTop from "./helpers/scrollTop";
import Login from "./login";

// ----------------------------------------------------------------------------

const Root: React.FunctionComponent = () => (
  <div>
    {/* <Global styles={globalStyles} /> */}
    <Helmet>
      <title>Userservice Admin!</title>
    </Helmet>
    <ScrollTop>
      <Switch>
        <Route path="/" exact component={Example} />
        <Route path="/login" component={Login} />
      </Switch>
    </ScrollTop>
  </div>
);

export default Root;
// export default hot(Root);
