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
import { Assignment, Brightness7, Home as HomeIcon, Label, MonetizationOn, People, ShowChart } from "@material-ui/icons";
import MenuIcon from "@material-ui/icons/Menu";
import React from "react";
import CookieConsent from "react-cookie-consent";
import { Link as RouterLink, Redirect, Route, Switch, useHistory, useLocation, withRouter } from "react-router-dom";
import { Timer as TimerEntry } from "../graphql";
import security from "../lib/security";
import { Timer as TimerSrv } from "../lib/timer";
import Clients from "./clients/clients";
import ScrollTop from "./helpers/scrollTop";
import Home from "./home";
import Login from "./login";
import Projects from "./projects/projects";
import Reports from "./reports/reports";
import "./root.css";
import Timer from "./timer/timer";
import TimerMenuEntry from "./timerMenuEntry";
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
          <Hidden xsDown implementation="css">
            My Time Tracker
          </Hidden>
          <Hidden smUp implementation="css">
            MTT
          </Hidden>
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
            <Hidden xsDown implementation="css">
              <Button component={RouterLink} to="/login" color="inherit">
                {props.username}
              </Button>
            </Hidden>
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
  let history = useHistory();
  const { container, mobileOpen, handleDrawerToggle, closeDrawer } = props;
  const theme = useTheme();

  let currentTimer: TimerEntry | undefined = undefined;

  const getIcon = (index: number) => {
    switch (index) {
      case 0:
        return (<Assignment />)
      case 1:
        return (<MonetizationOn />)
      case 2:
        return (<People />)
      case 3:
        return (<Label />)
      case 4:
        return (<Brightness7 />)

      default:
        return <></>
    }
  }
  const activeRouteTimer = activeRoute("/timer", location);
  const drawer = (
    <div>

      <Divider />
      <TimerMenuEntry isActiveRoute={activeRouteTimer} onClick={() => {

        closeDrawer();
        history.push("/timer");
      }} />
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
            onClick={closeDrawer}
          >
            <ListItemIcon>
              {index % 2 === 0 ? <HomeIcon /> : <ShowChart />}
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
            onClick={closeDrawer}
            selected={activeRoute(o.link, location)}
            key={o.txt}
            component={RouterLink}
            to={o.link}
          >
            <ListItemIcon>
              {
                getIcon(index)
              }

              {/* {index % 2 === 0 ? <InboxIcon /> : <MailIcon />} */}
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
          onClose={closeDrawer}
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

interface IPropsRoot { }
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
    console.log('Mounting Root:', this.isRunningStandalone())
    let isLoggedIn = false;
    try {
      isLoggedIn = await security.refresh();
      console.log("Refreshed", isLoggedIn);
      // isLoggedIn=false;
      this.setState({
        isLoggedIn,
        picture: security.picture,
        name: security.username,
      });
    } catch (ex) {
      console.log("Logout");
    }

    if (isLoggedIn) {
      setInterval(async () => {
        let isLoggedIn = false;
        try {
          isLoggedIn = await security.refresh();
          console.log('Refreshed token', isLoggedIn);
        } catch (ex) {

        }
        this.setState({
          isLoggedIn,
        });

      }, 1000 * 3000);
    } else {
      if (!activeRoute("/home", window.location)) {
        console.log("Redirect to home");
        window.location.replace("/home")
      }

    }
  };
  checkTimer = () => {
    this.setState({
      elapsed: this.state.elapsed + 1
    });
  };

  isRunningStandalone() {
    return (navigator as any).standalone || (window.matchMedia('(display-mode: standalone)').matches);
  }


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
              closeDrawer={() =>
                this.setState({ menuVisible: false })
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
                <Route path="/home" exact >
                  <Home history={window.history} isAuthorized={isLoggedIn}  />
                </Route>
                <Route path="/timer" exact component={Timer} />
                <Route path="/reports" exact component={Reports} />
                <Route path="/projects" exact component={Projects} />
                <Route path="/clients" exact component={Clients} />
                <Route path="/user" exact component={Users} />
                <Route path="/login" component={Login} />
              </Switch>
            </ScrollTop>
          </div>
        </div>
        {/* <Copyright></Copyright> */}
        <CookieConsent
          location="bottom"
          buttonText="Accept Cookies"
          cookieName="cookieAccepted"
          style={{ zIndex: 2000, background: "#2B373B" }}
          buttonStyle={{ color: "#4e503b", fontSize: "20px" }}
          expires={150}
        >
          Diese Webseite verwendet Cookies.<br></br>
          <span style={{ fontSize: "16px" }}>
            Why We Use Cookies
            This site uses cookies to make your browsing experience more convenient and personal.
            Cookies store useful information on your computer to help us improve the efficiency and relevance of our site for you.
            In some cases, they are essential to making the site work properly. By accessing this site, you consent to the use of cookies.
For more information, refer to our privacy policy and cookie policy.<br></br>
            Wir verwenden Cookies, um Inhalte und Anzeigen diser Website zu personalisieren und die aktuelle eingeloggte Sitzung zu speichern.
           </span>
        </CookieConsent>

      </div>
    );
  }
}

export default withRouter(Root as any);
// export default hot(Root);
