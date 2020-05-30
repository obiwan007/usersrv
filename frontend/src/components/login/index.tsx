// ReactQL example page - delete this folder for your own project!

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
import { Button, Container, FormControl, Grid, Paper } from "@material-ui/core";
import React from "react";
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
    dynamic: null,
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
      </>
    );
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
