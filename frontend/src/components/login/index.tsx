// ReactQL example page - delete this folder for your own project!

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
import { Button, Container, FormControl, Grid, IconButton, Paper, Slide } from "@material-ui/core";
import Snackbar from '@material-ui/core/Snackbar';
import { TransitionProps } from '@material-ui/core/transitions/transition';
import CloseIcon from '@material-ui/icons/Close';
import React from "react";
import TimerSrv from "../../lib/timer";


/* Local */

// ----------------------------------------------------------------------------

interface IIndexState {
  dynamic: React.SFC | null;
  openNotification: boolean;
}

interface IProps {
  history: any;
}

// Say hello from GraphQL, along with a HackerNews feed fetched by GraphQL
class Index extends React.PureComponent<IProps, IIndexState> {
  public state = {
    dynamic: null,
    openNotification: false,
  };

  public componentDidMount = async () => {
    // Fetch the component dynamically

    // ... and keep ahold of it locally
    this.setState({});
    const permission = Notification.permission;
    console.log("Allowed ", permission)
    this.setState({ openNotification: true });
    if (permission !== "granted") {
      this.setState({ openNotification: true });
    }

  };

  public render() {
    // const DynamicComponent = this.state.dynamic || (() => <h2>Loading...</h2>);
    const { openNotification } = this.state;
    return (
      <>
        <Container maxWidth="sm">
          {/* Note: The <h1> style will have a yellow background due to @/global/styles.ts! */}
          <Paper  >
            <h1>Login/Signup</h1>
            <Grid container spacing={3} justify="center" direction="row" alignItems="center">
              <Grid item>
                <FormControl>
                  <Button
                    variant="outlined"
                    onClick={() => this.loginClick(0)}
                  >
                    <img
                      style={{ width: "250px" }}
                      src="btn_google_signin_dark_normal_web@2x.png"></img>
                  </Button>

                </FormControl>
              </Grid>

            </Grid>

            <Grid container spacing={3} justify="center" direction="row" alignItems="center">
              <Grid item>
                <FormControl >
                  <Button
                    variant="outlined"
                    fullWidth={true}
                    onClick={() => this.loginClick(2)}
                  >
                    <img
                      style={{ width: "250px" }}
                      src="ms-symbollockup_signin_dark.svg"></img>
                  </Button>
                </FormControl>
              </Grid>
            </Grid>
          </Paper>
        </Container>
        <Snackbar
          open={openNotification}
          onClose={() => this.closeGrant()}
          TransitionComponent={this.TransitionUp}
          message="Allow Notifications for long running timers."
          action={
            <React.Fragment>
              <Button color="primary" size="small" onClick={() => this.closeGrant(true)}>
                ALLOW
              </Button>
              <IconButton
                aria-label="close"
                color="inherit"
                onClick={() => this.closeGrant()}
              >
                <CloseIcon />
              </IconButton>
            </React.Fragment>
          }
        />
      </>
    );
  }

  async closeGrant(allow: boolean = false) {
    this.setState({ openNotification: false });
    if (allow) {
      const notify = await Notification.requestPermission();
      console.log("Allowed ", notify)
      TimerSrv.notifyMessage("Thx for allowing notifications","We will notify you if some important event occured.");
    }
  }

  TransitionUp(props: TransitionProps) {
    return <Slide {...props} direction="up" />;
  }
  async loginClick(mode: number) {
    console.log("Redirect");
    switch (mode) {
      case 0:
        window.location.href = window.location.origin + "/auth/login";
        break;
      case 1:
        window.location.href = window.location.origin + "/auth/login?provider=apple";
        break;
      case 2:
        window.location.href = window.location.origin + "/auth/login?provider=microsoft";
        break;
    }
    // this.props.history.replace("/auth/login");
    // const ok = await Security.login("MyUsername@hotmail.com", "MyPassword");
    // console.log("Redirecting");
    // if (ok) {
    //   this.props.history.replace("/auth/login");
    // }
  }
}

export default Index;
