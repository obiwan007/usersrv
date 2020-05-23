// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from '@apollo/react-common';
import { Button, createStyles, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, Fab, FormControl, IconButton, InputLabel, List, ListItem, ListItemIcon, ListItemSecondaryAction, ListItemText, MenuItem, Select, TextField, Theme, WithStyles, withStyles, withTheme } from "@material-ui/core";
import { Add as AddIcon, Assignment, Delete } from "@material-ui/icons";
import React from "react";
import { withRouter } from "react-router-dom";
import { AllClientsComponent, AllProjectsComponent, Client, CreateProjectComponent, CreateProjectMutation, CreateProjectMutationVariables, DeleteProjectComponent, DeleteProjectMutation, DeleteProjectMutationVariables, Project, refetchAllProjectsQuery, UpdateProjectComponent, UpdateProjectMutation, UpdateProjectMutationVariables } from "../../graphql";
import client from "../../lib/client";
// import project, { ProjectEntry } from "../../lib/project";
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
      padding: spacing(1),
      margin: spacing(1),
      backgroundColor: palette.background.default,
      color: palette.primary.main,
    },
    fabButton: {
      position: 'absolute',
      zIndex: 10,
      bottom: spacing(2),
      right: spacing(2),
      margin: '0 auto',
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

  currentProject?: Project;
  columns: any[];
}
interface IProps {
  history?: any;
}

export class Projects extends React.PureComponent<PROPS_WITH_STYLES, IState> {
  getColumns = () => {
    return [
      { title: "Title", field: "name", editable: () => true, width: "40%" },
      {
        title: "Client",
        field: "clientId",
        editable: () => true,
        lookup: client.EntriesDict(),
        // render: (r: any) => {
        //   return <>{r.client?.name}</>;
        // },
      },
      { title: "Status", field: "status", editable: () => false },
      { title: "Seconds", field: "elapsedSeconds", editable: () => false },
      {
        title: "Team",
        field: "team",
        editable: () => true,
        lookup: client.EntriesDict(),
      },
    ];
  };

  /**
   *
   *
   */
  constructor(props: PROPS_WITH_STYLES, state: IState) {
    super(props, state);
    this.state = {
      columns: [],
      addOpen: false,
      editOpen: false,
    };
  }

  componentDidMount() {
    this.setState({ columns: this.getColumns() });
  }
  componentWillUnmount() { }
  render() {
    const { classes } = this.props;
    const { currentProject, columns } = this.state;
    console.log(this.props);
    const { addOpen, editOpen } = this.state;
    return (
      <>
        <DeleteProjectComponent>
          {(deleteProject, { data }) => {
            return (
              <AllProjectsComponent>
                {({ data, loading, error }) => {
                  // Any errors? Say so!
                  data?.allProjects?.forEach((d) => {
                    d && ((d as any).clientId = d?.client?.id);
                  });
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

                  return (
                    <div style={{
                      marginTop: 5,
                      overflowY: "auto",
                      height: "calc(100vh - 75px)",
                      minHeight: "calc(100vh - 75px)",
                    }}>
                      <List
                        component="nav"
                        aria-label="main mailbox folders"
                      >
                        {(data?.allProjects?.map(
                          (u) => u
                        ) as any[])?.map((entry: Project) => (
                          <ListItem button onClick={() => this.setState({ currentProject: Object.assign({}, entry), editOpen: true })}>
                            <ListItemIcon>
                              <Assignment />
                            </ListItemIcon>
                            <ListItemText primary={entry.name} secondary={entry.client?.name} ></ListItemText>
                            <ListItemSecondaryAction>
                              <IconButton onClick={() => this.deleteProject(entry, deleteProject)} edge="end" aria-label="delete">
                                <Delete />
                              </IconButton>
                            </ListItemSecondaryAction>
                          </ListItem>
                        ))}

                      </List>
                      <Fab onClick={() => this.addProject()} color="secondary" aria-label="add" className={classes.fabButton}>
                        <AddIcon />
                      </Fab>
                    </div>
                  );
                }}

              </AllProjectsComponent>
            )
          }}
        </DeleteProjectComponent>
        <CreateProjectComponent>
          {(createProject, { data }) => {
            return (
              <UpdateProjectComponent>
                {(updateProject, { data }) => {
                  return (
                    <AllClientsComponent>
                      {({ data, loading, error }) => {
                        const allClients = data;
                        if (columns) {
                          const f = columns.find((c) => c.title === "Client");
                          if (f) {
                            f.lookup = this.clientsDict(allClients?.allClients as Client[]);
                          }
                        }
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
                              {addOpen ? "Add new Project" : "Edit Project"}
                            </DialogTitle>
                            <DialogContent>
                              <DialogContentText>
                                To subscribe to this website, please
                                enter your email address here. We
                                will send updates occasionally.
                          </DialogContentText>
                              <form noValidate autoComplete="off">
                                <FormControl className={classes.formControl}>
                                  <TextField
                                    label="Name of Project"
                                    required
                                    autoFocus
                                    defaultValue={currentProject?.name}
                                    onChange={(data) => {
                                      if (currentProject) {
                                        currentProject.name = data.target.value!;
                                        console.log(currentProject);
                                        this.setState({
                                          currentProject: currentProject,
                                        });
                                      }
                                    }}
                                    margin="dense"
                                    id="name"
                                    type="text"
                                    fullWidth
                                  />
                                </FormControl>
                                <FormControl className={classes.formControl}>
                                  {/* <InputLabel id="demo-simple-select-helper-label">Description</InputLabel> */}
                                  <TextField
                                    label="Description"
                                    margin="dense"
                                    id="description"
                                    type="text"
                                    defaultValue={
                                      currentProject?.description
                                    }
                                    onChange={(data) => {
                                      if (currentProject) {
                                        currentProject.description = data.target.value!;
                                        console.log(currentProject);
                                        this.setState({
                                          currentProject: currentProject,
                                        });
                                      }
                                    }}
                                    fullWidth
                                  />
                                </FormControl>
                                <FormControl className={classes.formControl}>
                                  <InputLabel id="demo-simple-select-helper-label">Client</InputLabel>
                                  <Select
                                    // className={classes.selectEmpty}
                                    label="Client"
                                    value={
                                      allClients &&
                                        (allClients.allClients as any[])
                                          .length > 0
                                        ? currentProject?.client?.id
                                        : ""
                                    }
                                    onChange={(event) => {
                                      console.log(
                                        "Clientselection:",
                                        event.target
                                      );
                                      if (currentProject) {
                                        console.log('Selected', event.target.value as string);
                                        // </form>currentProject.client = event.target.value as string;
                                        const selected = allClients?.allClients?.find(c => c?.id === event.target.value as string);
                                        currentProject.client = selected;
                                        this.setState({ currentProject: currentProject });
                                      }
                                    }}
                                  >
                                    <MenuItem
                                      aria-label="None"
                                      value=""
                                    />
                                    {allClients &&
                                      allClients.allClients?.map((e: any) => (
                                        <MenuItem
                                          key={e.id!}
                                          value={e.id!}
                                        >
                                          {e.name}
                                        </MenuItem>
                                      )
                                      )}
                                  </Select>
                                </FormControl>

                              </form>
                            </DialogContent>
                            <DialogActions>
                              <Button
                                onClick={() =>
                                  this.handleClose(false)
                                }
                                color="primary"
                              >
                                Cancel
                          </Button>
                              <Button
                                onClick={() =>
                                  this.handleClose(true, updateProject, createProject)
                                }
                                color="primary"
                              >
                                {addOpen ? "Create" : "Update"}
                              </Button>
                            </DialogActions>
                          </Dialog>

                        );
                      }}
                    </AllClientsComponent>
                  )
                }}
              </UpdateProjectComponent>
            )
          }}
        </CreateProjectComponent>
      </>
    );
  }


  clientsDict(data: Client[] | null | undefined): any {
    if (!data) {
      return {};
    }
    const dict: any = {};
    data.forEach((e) => e.id && (dict[+e.id] = e.name));
    return dict;
  }
  deleteProject = (entity: Project, deleteProject: MutationFunction<DeleteProjectMutation, DeleteProjectMutationVariables>) => {
    deleteProject({
      refetchQueries: [
        refetchAllProjectsQuery(),
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
    this.setState({ addOpen: true, currentProject: { name: 'Projectname', description: '', client: { id: '1' } } });
  };
  handleClose = async (state: any,
    updateProject?: MutationFunction<UpdateProjectMutation, UpdateProjectMutationVariables>,
    createProject?: MutationFunction<CreateProjectMutation, CreateProjectMutationVariables>,
  ) => {
    console.log(state);
    // if (state) {
    //   project.addProject(this.state.currentProject);
    const newData = this.state.currentProject!;
    const data = {
      id: newData.id,
      name: newData.name,
      description:
        newData.description,
      team: newData.team,
      status: newData.status,
      client: newData.client?.id,
    };
    console.log('Newdata', data)
    if (this.state.editOpen && updateProject) {
      await updateProject({
        refetchQueries: [
          refetchAllProjectsQuery(),
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
    else if (this.state.addOpen && createProject) {
      await createProject({
        refetchQueries: [
          refetchAllProjectsQuery(),
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



    // }
    this.setState({ addOpen: false, editOpen: false });
  };
}

export default withTheme(
  withStyles(styles as any)(withRouter(Projects as any))
);
