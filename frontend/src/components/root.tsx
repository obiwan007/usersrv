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
import Link from '@material-ui/core/Link';
import Typography from '@material-ui/core/Typography';
import React from "react";
import { Route, Switch } from "react-router-dom";
import Example from "./example";
import ScrollTop from "./helpers/scrollTop";
import Login from "./login";
import './root.css';
// ----------------------------------------------------------------------------

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Markus Miertschink
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}


const Root: React.FunctionComponent = () => (
  // <Container maxWidth="lg">
  <div>
    {/* <Global styles={globalStyles} /> */}
    {/* <Helmet>
        <title>Userservice Admin!</title>
      </Helmet> */}
    <div className="content">

      <ScrollTop>
        <Switch>
          <Route path="/" exact component={Example} />
          <Route path="/login" component={Login} />
        </Switch>

      </ScrollTop>
    </div>
    <Copyright></Copyright>
  </div>
  // </Container>

);

export default Root;
// export default hot(Root);
