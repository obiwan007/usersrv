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
import { AppBar, Button, Divider, IconButton, List, ListItem, ListItemIcon, ListItemText, makeStyles, Toolbar, useTheme } from "@material-ui/core";
import Avatar from "@material-ui/core/Avatar";
import Drawer from "@material-ui/core/Drawer";
import Hidden from "@material-ui/core/Hidden";
import Link from "@material-ui/core/Link";
import Typography from "@material-ui/core/Typography";
import MailIcon from "@material-ui/icons/Mail";
import MenuIcon from "@material-ui/icons/Menu";
import InboxIcon from "@material-ui/icons/MoveToInbox";
import React from "react";
import { Link as RouterLink, Redirect, Route, Switch, useLocation, withRouter } from "react-router-dom";
import { AllTimerComponent, Timer as TimerEntry } from "../graphql";
import security from "../lib/security";
import { Timer as TimerSrv } from "../lib/timer";
import Clients from "./clients/clients";
import ScrollTop from "./helpers/scrollTop";
import Home from "./home";
import Login from "./login";
import Projects from "./projects/projects";
import "./root.css";
import Timer from "./timer/timer";
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
    <Typography
      variant="body2"
      style={{ paddingTop: 10, background: "#556cd6", height: 40 }}
      color="textSecondary"
      align="center"
    >
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
          My Time Tracker
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

function activeRoute(routeName: string, location: any) {
  return location.pathname.indexOf(routeName) > -1 ? true : false;
}

function showElapsed(t: any): string {
  if (t) {
    const time1 = new Date(t.timerStart!);
    const time2 = new Date();
    return TimerSrv.hms((time2.getTime() - time1.getTime()) / 1000);
  }
  return "Timer";
}
export function MainMenu(props: any) {
  const classes = useStyles();
  let location = useLocation();
  const { container, mobileOpen, handleDrawerToggle } = props;
  const theme = useTheme();
  
  let currentTimer: TimerEntry | undefined = undefined;
  const drawer = (
    <div>
      
      <Divider />
      <List>
        <AllTimerComponent variables={{ d: { dayrange: "50" } }}>
          {({ data, loading, error }) => {
            // Any errors? Say so!
            currentTimer = undefined;
            data?.allTimer?.forEach((d) => {
              if (d) {
                (d as any).projectId = d?.project?.id;
                if (d.isRunning === true) {
                  currentTimer = d;
                }
              }
            });
            const timerText =
              currentTimer && currentTimer!.isRunning && !activeRoute("/timer", location)
                ? showElapsed(currentTimer)
                : "Timer";
            // console.log("currentTimer", currentTimer, data?.allTimer);
            if (error) {
              return (
                <div>
                  <h1>Error retrieving Timer list &mdash; {error.message}</h1>
                  {/* <Button variant="contained" color="secondary" onClick={() => this.refreshClick()}>Refresh</Button> */}
                </div>
              );
            }

            // If the data is still loading, return with a basic
            // message to alert the user

            return (
              <>
                {[{ txt: timerText, link: "/timer" }].map((o, index) => (
                  <ListItem
                    selected={activeRoute(o.link, location)}
                    component={RouterLink}
                    to={o.link}
                    button
                    key={o.txt}
                  >
                    <ListItemIcon>
                      {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
                    </ListItemIcon>
                    <ListItemText primary={o.txt} />
                  </ListItem>
                ))}
              </>
            );
          }}
        </AllTimerComponent>
      </List>
      <Divider />
      <List>
        {[
          { txt: "Home", link: "/home" },
          { txt: "Reports", link: "/reports" },
        ].map((o, index) => (
          <ListItem
            selected={activeRoute(o.link, location)}
            component={RouterLink}
            to={o.link}
            button
            key={o.txt}
          >
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
          <ListItem
            button
            selected={activeRoute(o.link, location)}
            key={o.txt}
            component={RouterLink}
            to={o.link}
          >
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
  isRunning: boolean;
  elapsed: number;
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
    isRunning: false,
    elapsed: 0,
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
    setInterval(() => {
      this.checkTimer();
    }, 1000);
  };
  checkTimer = () => {
    this.setState({
      elapsed: this.state.elapsed+1
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
              {...this.state}
              handleDrawerToggle={() =>
                this.setState({ menuVisible: !menuVisible })
              }
              mobileOpen={menuVisible}
            ></MainMenu>
          )}

          <div className="content" style={{ flexGrow: 1, padding: 10 }}>
            <ScrollTop>
              <Switch>
                <Route path="/" exact component={Home}>
                  <Redirect
                    to={{
                      pathname: "/home",
                    }}
                  />
                </Route>
                <Route path="/home" exact component={Home} />
                <Route path="/timer" exact component={Timer} />
                <Route path="/projects" exact component={Projects} />
                <Route path="/clients" exact component={Clients} />
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

export default withRouter(Root as any);
// export default hot(Root);
