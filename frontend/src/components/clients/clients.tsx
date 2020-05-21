// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from '@apollo/react-common';
import { Button, createStyles, Dialog, DialogActions, DialogContent, DialogTitle, Fab, FormControl, IconButton, List, ListItem, ListItemIcon, ListItemSecondaryAction, ListItemText, TextField, Theme, withStyles, WithStyles, withTheme } from "@material-ui/core";
import { Add as AddIcon, Delete, MonetizationOn } from '@material-ui/icons';
import React from "react";
import { withRouter } from "react-router-dom";
import { AllClientsComponent, Client, ClientInput, CreateClientComponent, CreateClientMutation, CreateClientMutationVariables, DeleteClientComponent, DeleteClientMutation, DeleteClientMutationVariables, refetchAllClientsQuery, UpdateClientComponent, UpdateClientMutation, UpdateClientMutationVariables } from "../../graphql";
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
    fabButton: {
      position: 'absolute',
      zIndex: 10,
      // bottom: 30,
      // left: 10,
      // right: 10,
      margin: '0 auto',
      bottom: spacing(2),
      right: spacing(2),
    },
    formControl: {
      margin: spacing(1),
      minWidth: 120,
      width: "100%"
    },
  });

export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
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
  editOpen: boolean;

  currentClient: Client;
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
      editOpen: false,
      currentClient: {},
    };
  }

  componentDidMount() { }
  componentWillUnmount() { }
  render() {
    const { classes } = this.props;
    const { currentClient, addOpen, editOpen } = this.state;

    return (
      <>

        <DeleteClientComponent>
          {(deleteClient, { data }) => {
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
                      {/*         

        <Button onClick={() => this.addProject()}>Add Client</Button> */}

                      <div>
                        <List
                          component="nav"
                          aria-label="main mailbox folders"
                        >
                          {(data?.allClients?.map(
                            (u) => u
                          ) as any[])?.map((entry: Client) => (
                            <ListItem button onClick={() => this.setState({ currentClient: Object.assign({}, entry), editOpen: true })}>
                              <ListItemIcon>
                                <MonetizationOn />
                              </ListItemIcon>
                              <ListItemText primary={entry.name} secondary={entry.description} ></ListItemText>
                              <ListItemSecondaryAction>
                                <IconButton onClick={() => this.deleteClient(entry, deleteClient)} edge="end" aria-label="delete">
                                  <Delete />
                                </IconButton>
                              </ListItemSecondaryAction>
                            </ListItem>
                          ))}

                        </List>
                      </div>
                      <Fab onClick={() => this.addProject()} color="secondary" aria-label="add" className={classes.fabButton}>
                        <AddIcon />
                      </Fab>

                    </div>
                  );
                }}
              </AllClientsComponent>
            );
          }}
        </DeleteClientComponent>
        <UpdateClientComponent>
          {(updateClient, { data }) => {
            return (
              <CreateClientComponent>
                {(createClient, { data }) => {
                  return (

                    <Dialog
                      open={addOpen || editOpen}
                      onClose={() =>
                        this.setState({ addOpen: false, editOpen: false })
                      }
                      aria-labelledby="simple-modal-title"
                      aria-describedby="simple-modal-description"
                    >
                      <DialogTitle id="simple-dialog-title">
                        {editOpen ? "Edit Client" : "Add new Client"}
                      </DialogTitle>
                      <DialogContent>
                        {/* <DialogContentText>
                          To subscribe to this website, please enter
                          your email address here. We will send
                          updates occasionally.
                    </DialogContentText> */}
                        <form noValidate autoComplete="off">
                          <FormControl className={classes.formControl}>
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
                          </FormControl>
                          <FormControl className={classes.formControl}>
                            <TextField
                              margin="dense"
                              id="description"
                              label="Description of Project"
                              type="text"
                              defaultValue={currentClient.description}
                              onChange={(data) => {
                                currentClient.description = data.target.value!;
                                console.log(currentClient);
                                this.setState({
                                  currentClient: currentClient,
                                });
                              }}
                              fullWidth
                            />
                          </FormControl>
                          <FormControl className={classes.formControl}>
                            <TextField
                              margin="dense"
                              id="Address"
                              multiline={true}
                              rowsMax="8"
                              label="Address of the Client"
                              type="text"
                              defaultValue={currentClient.address}
                              onChange={(data) => {
                                currentClient.address = data.target.value!;
                                this.setState({
                                  currentClient: currentClient,
                                });
                              }}
                              fullWidth
                            />
                          </FormControl>
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
                          onClick={() => this.handleClose(true, updateClient, createClient)}
                          color="primary"
                        >
                          {addOpen ? "Create" : "Update"}
                        </Button>
                      </DialogActions>

                    </Dialog>)
                }}
              </CreateClientComponent>)
          }}
        </UpdateClientComponent>
      </>
    );
  }

  deleteClient = (entity: Client, deleteProject: MutationFunction<DeleteClientMutation, DeleteClientMutationVariables>) => {
    deleteProject({
      refetchQueries: [
        refetchAllClientsQuery(),
      ],
      fetchPolicy: "no-cache",
      variables: {
        d: entity!.id!,
      },
    }).catch((ex) => {
      console.log(
        "Error in mutation",
        ex
      );
    });
  }
  addProject = () => {
    this.setState({ addOpen: true, currentClient: { name: 'Client Name', description: '' } });
  };

  handleClose = async (state: any,
    updateClient?: MutationFunction<UpdateClientMutation, UpdateClientMutationVariables>,
    createClient?: MutationFunction<CreateClientMutation, CreateClientMutationVariables>,
  ) => {
    console.log(state);
    // if (state) {
    //   project.addProject(this.state.currentProject);
    const newData = this.state.currentClient!;
    const data: ClientInput = {
      id: newData.id,
      name: newData.name,
      description:
        newData.description,
      address:
        newData.address,
    };
    console.log('Newdata', data)
    if (this.state.editOpen && updateClient) {
      await updateClient({
        refetchQueries: [
          refetchAllClientsQuery(),
        ],
        fetchPolicy: "no-cache",
        variables: {
          d: data
        },
      }).catch((ex) => {
        console.log(
          "Error in mutation",
          ex
        );
      });
    }
    else if (this.state.addOpen && createClient) {
      await createClient({
        refetchQueries: [
          refetchAllClientsQuery(),
        ],
        fetchPolicy: "no-cache",
        variables: {
          d: data
        },
      }).catch((ex) => {
        console.log(
          "Error in mutation",
          ex
        );
      });
    }
    this.setState({ addOpen: false, editOpen: false });
  };

}

export default withTheme(withStyles(styles as any)(withRouter(Clients as any)));
