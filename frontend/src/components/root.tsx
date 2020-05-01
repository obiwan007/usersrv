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
import {
  AppBar,
  Button,
  Divider,
  IconButton,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  makeStyles,
  Toolbar,
  useTheme,
} from "@material-ui/core";
import Avatar from "@material-ui/core/Avatar";
import Drawer from "@material-ui/core/Drawer";
import Hidden from "@material-ui/core/Hidden";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import MailIcon from "@material-ui/icons/Mail";
import MenuIcon from "@material-ui/icons/Menu";
import InboxIcon from "@material-ui/icons/MoveToInbox";
import React from "react";
import { Link as RouterLink, Route, Switch } from "react-router-dom";
import security from "../lib/security";
import ScrollTop from "./helpers/scrollTop";
import Home from "./home";
import Login from "./login";
import "./root.css";
import Users from "./users";

// ----------------------------------------------------------------------------
const drawerWidth = 200;
const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
    [theme.breakpoints.up("sm")]: {
      display: "none",
    },
  },
  title: {
    flexGrow: 1,
  },
  appBar: {
    [theme.breakpoints.up("sm")]: {
      width: `calc(100% - ${drawerWidth}px)`,
      marginLeft: drawerWidth,
    },
  },
  drawer: {
    [theme.breakpoints.up("sm")]: {
      width: drawerWidth,
      flexShrink: 0,
    },
  },
  drawerPaper: {
    width: drawerWidth,
  },
  drawerHeader: {
    display: "flex",
    alignItems: "center",
    padding: theme.spacing(0, 1),
    // necessary for content to be below app bar
    ...theme.mixins.toolbar,
    justifyContent: "flex-end",
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
  toolbar: theme.mixins.toolbar,
}));

function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {"Copyright © "}
      <Link color="inherit" href="https://material-ui.com/">
        Markus Miertschink
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}
export function ButtonAppBar(props: any) {
  const classes = useStyles();
  const { picture, handleDrawerToggle, isLoggedIn } = props;
  return (
    // <div className={classes.root}>
    <AppBar
      position="fixed"
      className={classes.appBar}
      style={!isLoggedIn ? { width: "100%" } : {}}
    >
      <Toolbar>
        {isLoggedIn && (
          <IconButton
            onClick={handleDrawerToggle}
            edge="start"
            className={classes.menuButton}
            color="inherit"
            aria-label="menu"
          >
            <MenuIcon />
          </IconButton>
        )}
        <Typography variant="h6" className={classes.title}>
          MyPlaner
        </Typography>
        {/* <Link
            component={RouterLink} to="/login"
            variant="body1"
            color="inherit"
            href="/login"
          >
            Login
          </Link> */}
        {!props.isLoggedIn && (
          <Button component={RouterLink} to="/login" color="inherit">
            Login
          </Button>
        )}
        {props.isLoggedIn && (
          <>
            <Button component={RouterLink} to="/user" color="inherit">
              {props.username}
            </Button>
            <Avatar alt="Remy Sharp" src={picture}>
              {props.username}
            </Avatar>
          </>
        )}
      </Toolbar>
    </AppBar>
    // </div >
  );
}

export function MainMenu(props: any) {
  const classes = useStyles();
  const { container, mobileOpen, handleDrawerToggle } = props;
  const theme = useTheme();
  console.log("MobileOpen", mobileOpen);
  const drawer = (
    <div>
      {/* <div className={classes.toolbar} /> */}
      {/* <div className={classes.drawerHeader}>
        <IconButton onClick={handleDrawerToggle}>
          {theme.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
        </IconButton>
      </div> */}
      <Divider />
      <List>
        {[{ txt: "Timer", link: "/timer" }].map((o, index) => (
          <ListItem component={RouterLink} to={o.link} button key={o.txt}>
            <ListItemIcon>
              {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
            </ListItemIcon>
            <ListItemText primary={o.txt} />
          </ListItem>
        ))}
      </List>
      <Divider />
      <List>
        {[
          { txt: "Home", link: "/" },
          { txt: "Reports", link: "/reports" },
        ].map((o, index) => (
          <ListItem component={RouterLink} to={o.link} button key={o.txt}>
            <ListItemIcon>
              {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
            </ListItemIcon>
            <ListItemText primary={o.txt} />
          </ListItem>
        ))}
      </List>
      <Divider />
      <List>
        {[
          { txt: "Projects", link: "/projects" },
          { txt: "Clients", link: "/clients" },
          { txt: "Team", link: "/team" },
          { txt: "Tags", link: "/tags" },
          { txt: "Settings", link: "/settings" },
          { txt: "User Admin", link: "/user" },
        ].map((o, index) => (
          <ListItem button key={o.txt} component={RouterLink} to={o.link}>
            <ListItemIcon>
              {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
            </ListItemIcon>
            <ListItemText primary={o.txt} />
          </ListItem>
        ))}
      </List>
    </div>
  );

  return (
    <nav className={classes.drawer} aria-label="mailbox folders">
      {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
      <Hidden smUp implementation="css">
        <Drawer
          container={container}
          variant="temporary"
          anchor={theme.direction === "rtl" ? "right" : "left"}
          open={mobileOpen}
          onClose={handleDrawerToggle}
          classes={{
            paper: classes.drawerPaper,
          }}
          ModalProps={{
            keepMounted: true, // Better open performance on mobile.
          }}
        >
          {drawer}
        </Drawer>
      </Hidden>
      <Hidden xsDown implementation="css">
        <Drawer
          classes={{
            paper: classes.drawerPaper,
          }}
          variant="permanent"
          open
        >
          {drawer}
        </Drawer>
      </Hidden>
    </nav>
  );
}

interface IStateRoot {
  dynamic: React.SFC | null;
  isLoggedIn: boolean;
  picture: string;
  name: string;
  menuVisible: boolean;
}

interface IPropsRoot {}
// Say hello from GraphQL, along with a HackerNews feed fetched by GraphQL
class Root extends React.PureComponent<IPropsRoot, IStateRoot> {
  public state = {
    dynamic: null,
    isLoggedIn: false,
    picture: "",
    name: "",
    menuVisible: false,
  };
  public componentDidMount = async () => {
    // Fetch the component dynamically

    // ... and keep ahold of it locally
    const isLoggedIn = await security.refresh();
    console.log("Refreshed", isLoggedIn);

    this.setState({
      isLoggedIn,
      picture: security.picture,
      name: security.username,
    });
  };

  public render() {
    const { menuVisible, isLoggedIn } = this.state;
    return (
      <div>
        {/* <Global styles={globalStyles} /> */}
        {/* <Helmet>
          <title>Userservice Admin!</title>
        </Helmet> */}

        <div style={{ display: "flex" }}>
          <ButtonAppBar
            handleDrawerToggle={() =>
              this.setState({ menuVisible: !menuVisible })
            }
            username={this.state.name}
            picture={this.state.picture}
            isLoggedIn={this.state.isLoggedIn}
          ></ButtonAppBar>
          {isLoggedIn && (
            <MainMenu
              handleDrawerToggle={() =>
                this.setState({ menuVisible: !menuVisible })
              }
              mobileOpen={menuVisible}
            ></MainMenu>
          )}

          <div className="content" style={{ flexGrow: 1, padding: 0 }}>
            <ScrollTop>
              <Switch>
                <Route path="/" exact component={Home} />
                <Route path="/user" exact component={Users} />
                <Route path="/login" component={Login} />
              </Switch>
            </ScrollTop>
          </div>
        </div>
        <Copyright></Copyright>
      </div>
    );
  }
}

export default Root;
// export default hot(Root);
