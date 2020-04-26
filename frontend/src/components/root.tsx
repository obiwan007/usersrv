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
import { AppBar, Button, IconButton, makeStyles, Toolbar } from '@material-ui/core';
import Link from '@material-ui/core/Link';
import Typography from '@material-ui/core/Typography';
import MenuIcon from '@material-ui/icons/Menu';
import React from "react";
import { Link as RouterLink, Route, Switch } from 'react-router-dom';
import security from '../lib/security';
import ScrollTop from "./helpers/scrollTop";
import Login from "./login";
import './root.css';
import Users from "./users";
// ----------------------------------------------------------------------------

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
}));


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
export function ButtonAppBar(props: any) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" className={classes.title}>
            News
          </Typography>
          {/* <Link
            component={RouterLink} to="/login"
            variant="body1"
            color="inherit"
            href="/login"
          >
            Login
          </Link> */}
          {!props.isLoggedIn && <Button component={RouterLink} to="/login" color="inherit" >Login</Button>}
          {props.isLoggedIn && <Button component={RouterLink} to="/user" color="inherit" >User</Button>}

        </Toolbar>
      </AppBar>
    </div >
  );
}


interface IStateRoot {
  dynamic: React.SFC | null;
  isLoggedIn: boolean;
}

interface IPropsRoot {

}
// Say hello from GraphQL, along with a HackerNews feed fetched by GraphQL
class Root extends React.PureComponent<IPropsRoot, IStateRoot> {
  public state = {
    dynamic: null,
    isLoggedIn: false,
  };

  public componentDidMount = async () => {
    // Fetch the component dynamically

    // ... and keep ahold of it locally
    const isLoggedIn = await security.refresh();
    console.log("Refreshed", isLoggedIn);
    this.setState({
      isLoggedIn
    });
  };

  public render() {
    const DynamicComponent = this.state.dynamic || (() => <h2>Loading...</h2>);

    return (
      <div>
        {/* <Global styles={globalStyles} /> */}
        {/* <Helmet>
          <title>Userservice Admin!</title>
        </Helmet> */}

        <ButtonAppBar isLoggedIn={this.state.isLoggedIn}></ButtonAppBar>

        <div className="content">

          <ScrollTop>
            <Switch>
              <Route path="/" exact component={Users} />
              <Route path="/login" component={Login} />
            </Switch>

          </ScrollTop>
        </div>
        <Copyright></Copyright>
      </div>
    );
  }
}



export default Root;
// export default hot(Root);
