// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import {
  Button,
  createStyles,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  List,
  TextField,
  Theme,
  withStyles,
  WithStyles,
  withTheme,
} from "@material-ui/core";
import MaterialTable from "material-table";
import React from "react";
import { withRouter } from "react-router-dom";
import {
  AllClientsComponent,
  CreateClientComponent,
  DeleteClientComponent,
  refetchAllClientsQuery,
  UpdateClientComponent,
} from "../../graphql";
import client, { ClientEntry } from "../../lib/client";
// ----------------------------------------------------------------------------

// const styles = (theme: any) => {
//   root: {
//     "& > *": {
//       margin: theme.spacing(1),
//       width: "25ch",
//     },
//   },
// };
// Non-dependent styles

// Theme-dependent styles
const styles = ({ palette, spacing }: Theme) =>
  createStyles({
    root: {
      display: "flex",
      width: "300px",
      flexDirection: "column",
      padding: spacing,
      margin: spacing(1),
      backgroundColor: palette.background.default,
      color: palette.primary.main,
    },
  });

export type PROPS_WITH_STYLES = IProps & WithStyles<"root">;
// withStyles(({ palette, spacing }) => ({
//   root: {
//     display: "flex",
//     flexDirection: "column",
//     padding: spacing.unit,
//     backgroundColor: palette.background.default,
//     color: palette.primary.main,
//   },
// }));

interface IState {
  addOpen: boolean;

  currentClient: ClientEntry;
}
interface IProps {
  history?: any;
}

export class Clients extends React.PureComponent<PROPS_WITH_STYLES, IState> {
  columns = [
    { title: "Name", field: "name", width: "30%" },
    { title: "Description", field: "description", width: "30%" },
    { title: "Address", field: "address", width: "40%" },
  ];
  interval?: NodeJS.Timeout;

  /**
   *
   *
   */
  constructor(props: PROPS_WITH_STYLES, state: IState) {
    super(props, state);
    this.state = {
      addOpen: false,
      currentClient: new ClientEntry(),
    };
  }

  componentDidMount() {}
  componentWillUnmount() {}
  render() {
    const { classes } = this.props;
    const { currentClient, addOpen } = this.state;

    return (
      <UpdateClientComponent>
        {(updateClient, { data }) => {
          console.log("Updateclient MutationData", data);
          return (
            <DeleteClientComponent>
              {(deleteClient, { data }) => {
                return (
                  <CreateClientComponent>
                    {(createClient, { data }) => {
                      console.log("MutationData", data);
                      return (
                        <AllClientsComponent>
                          {({ data, loading, error }) => {
                            // Any errors? Say so!
                            console.log("Returned data", data);
                            if (error) {
                              return (
                                <div>
                                  <h1>
                                    Error retrieving users list &mdash;{" "}
                                    {error.message}
                                  </h1>
                                  <Button
                                    variant="contained"
                                    color="primary"
                                    onClick={() =>
                                      this.props.history.push("/login")
                                    }
                                  >
                                    Login
                                  </Button>
                                  {/* <Button variant="contained" color="secondary" onClick={() => this.refreshClick()}>Refresh</Button> */}
                                </div>
                              );
                            }

                            // If the data is still loading, return with a basic
                            // message to alert the user
                            // if (loading) {
                            //   return <h1>Loading Projects...</h1>;
                            // }
                            return (
                              <div>
                                <h3>Clients</h3>
                                {/*         

        <Button onClick={() => this.addProject()}>Add Client</Button> */}

                                <List
                                  component="nav"
                                  aria-label="main mailbox folders"
                                >
                                  <MaterialTable
                                    isLoading={loading}
                                    options={{
                                      padding: "dense",
                                      minBodyHeight: "calc(100vh - 306px)",
                                      maxBodyHeight: "calc(100vh - 306px)",
                                    }}
                                    title="Clients"
                                    columns={this.columns}
                                    data={
                                      data?.allClients?.map((u) => u) as any[]
                                    }
                                    editable={{
                                      isEditable: (rowData) => true, // only name(a) rows would be editable
                                      isDeletable: (rowData) => true, // only name(a) rows would be deletable
                                      onRowAdd: (newData) =>
                                        createClient({
                                          refetchQueries: [
                                            refetchAllClientsQuery(),
                                          ],
                                          fetchPolicy: "no-cache",
                                          variables: {
                                            d: {
                                              name: newData.name,
                                              description: newData.description,
                                              address: newData.address,
                                            },
                                          },
                                        }).catch((ex) => {
                                          console.log("Error in mutation", ex);
                                        }),
                                      onRowUpdate: (newData, oldData) =>
                                        updateClient({
                                          variables: {
                                            d: {
                                              id: newData.id,
                                              name: newData.name,
                                              description: newData.description,
                                              address: newData.address,
                                            },
                                          },
                                        }),
                                      onRowDelete: (oldData) =>
                                        deleteClient({
                                          refetchQueries: [
                                            refetchAllClientsQuery(),
                                          ],
                                          fetchPolicy: "no-cache",
                                          variables: {
                                            d: oldData.id,
                                          },
                                        }).catch((ex) => {
                                          console.log("Error in mutation", ex);
                                        }),
                                    }}
                                  />
                                  {/* {data!.allUsers!.map(data => (
    <ListItem button>
      <ListItemIcon>
        <Person />
      </ListItemIcon>
      <ListItemText primary={data!.name} secondary={data!.email} />
    </ListItem>
  ))} */}
                                </List>
                                <Dialog
                                  open={addOpen}
                                  onClose={() =>
                                    this.setState({ addOpen: false })
                                  }
                                  aria-labelledby="simple-modal-title"
                                  aria-describedby="simple-modal-description"
                                >
                                  <DialogTitle id="simple-dialog-title">
                                    Add new Client
                                  </DialogTitle>
                                  <DialogContent>
                                    <DialogContentText>
                                      To subscribe to this website, please enter
                                      your email address here. We will send
                                      updates occasionally.
                                    </DialogContentText>
                                    <form noValidate autoComplete="off">
                                      <TextField
                                        required
                                        autoFocus
                                        defaultValue={currentClient.name}
                                        onChange={(data) => {
                                          currentClient.name = data.target.value!;
                                          console.log(currentClient);
                                          this.setState({
                                            currentClient: currentClient,
                                          });
                                        }}
                                        margin="dense"
                                        id="name"
                                        label="Title of Project"
                                        type="text"
                                        fullWidth
                                      />
                                      <TextField
                                        margin="dense"
                                        id="description"
                                        label="Description of Project"
                                        type="text"
                                        defaultValue={currentClient.notes}
                                        onChange={(data) => {
                                          currentClient.notes = data.target.value!;
                                          console.log(currentClient);
                                          this.setState({
                                            currentClient: currentClient,
                                          });
                                        }}
                                        fullWidth
                                      />
                                    </form>
                                  </DialogContent>
                                  <DialogActions>
                                    <Button
                                      onClick={() => this.handleClose(false)}
                                      color="primary"
                                    >
                                      Cancel
                                    </Button>
                                    <Button
                                      onClick={() => this.handleClose(true)}
                                      color="primary"
                                    >
                                      Add
                                    </Button>
                                  </DialogActions>
                                  {/* <div>
            <form className={classes.root} noValidate autoComplete="off">
              <TextField id="standard-basic" label="Standard" />
              <TextField id="filled-basic" label="Filled" variant="filled" />
              <TextField
                id="outlined-basic"
                label="Outlined"
                variant="outlined"
              />
              <Button>Save</Button>
              <Button>Save</Button>
            </form>
          </div> */}
                                </Dialog>
                              </div>
                            );
                          }}
                        </AllClientsComponent>
                      );
                    }}
                  </CreateClientComponent>
                );
              }}
            </DeleteClientComponent>
          );
        }}
      </UpdateClientComponent>
    );
  }
  addProject = () => {
    this.setState({ addOpen: true, currentClient: new ClientEntry() });
  };
  handleClose = (state: any) => {
    console.log(state);
    if (state) {
      client.add(this.state.currentClient);
    }
    this.setState({ addOpen: false });
  };
}

export default withTheme(withStyles(styles as any)(withRouter(Clients as any)));
