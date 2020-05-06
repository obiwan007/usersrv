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
  WithStyles,
  withStyles,
  withTheme,
} from "@material-ui/core";
import MaterialTable from "material-table";
import React from "react";
import { withRouter } from "react-router-dom";
import { AllProjectsComponent } from "../../graphql";
import client from "../../lib/client";
import project, { ProjectEntry } from "../../lib/project";
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

  currentProject: ProjectEntry;
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
        field: "client.name",
        editable: () => true,
        lookup: client.EntriesDict(),
        render: (r: any) => {
          return <>{r.client.name}</>;
        },
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
      currentProject: new ProjectEntry(),
    };
  }

  componentDidMount() {
    this.setState({ columns: this.getColumns() });
  }
  componentWillUnmount() {}
  render() {
    const { classes } = this.props;
    const { currentProject, columns } = this.state;
    console.log(this.props);
    const { addOpen } = this.state;
    return (
      <AllProjectsComponent>
        {({ data, loading, error }) => {
          // Any errors? Say so!
          console.log("Returned data", data);
          if (error) {
            return (
              <div>
                <h1>Error retrieving users list &mdash; {error.message}</h1>
                <Button
                  variant="contained"
                  color="primary"
                  onClick={() => this.props.history.push("/login")}
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
            <div>
              <h3>Projects</h3>

              <List component="nav" aria-label="main mailbox folders">
                <MaterialTable
                  isLoading={loading}
                  options={{
                    padding: "dense",
                    minBodyHeight: "calc(100vh - 360px)",
                    maxBodyHeight: "calc(100vh - 360px)",
                  }}
                  title="Projects"
                  columns={columns}
                  data={data?.allProjects?.map((u) => u) as any[]}
                  // data={list?.map((u) => u) as any[]}
                  editable={{
                    isEditable: (rowData) => {
                      return true;
                    }, // only name(a) rows would be editable
                    isDeletable: (rowData) => true, // only name(a) rows would be deletable
                    onRowAdd: (newData) =>
                      new Promise((resolve, reject) => {
                        setTimeout(() => {
                          {
                            const d = project.add(newData);
                            console.log("New List", d);
                            // this.setState({ list: d }, () => resolve());
                          }
                          resolve();
                        }, 1000);
                      }),
                    onRowUpdate: (newData, oldData) =>
                      new Promise((resolve, reject) => {
                        setTimeout(() => {
                          {
                            const d = project.update(newData);
                            console.log("New List", d);
                            // this.setState({ list: d }, () => resolve());
                          }
                          resolve();
                        }, 1000);
                      }),
                    onRowDelete: (oldData) =>
                      new Promise((resolve, reject) => {
                        setTimeout(() => {
                          {
                            const d = project.del(oldData);
                            console.log("New List", d);
                            // this.setState({ list: d }, () => resolve());
                          }
                          resolve();
                        }, 1000);
                      }),
                  }}
                />
              </List>
              <Dialog
                open={addOpen}
                onClose={() => this.setState({ addOpen: false })}
                aria-labelledby="simple-modal-title"
                aria-describedby="simple-modal-description"
              >
                <DialogTitle id="simple-dialog-title">
                  Add new Project
                </DialogTitle>
                <DialogContent>
                  <DialogContentText>
                    To subscribe to this website, please enter your email
                    address here. We will send updates occasionally.
                  </DialogContentText>
                  <form noValidate autoComplete="off">
                    <TextField
                      required
                      autoFocus
                      defaultValue={currentProject.name}
                      onChange={(data) => {
                        currentProject.name = data.target.value!;
                        console.log(currentProject);
                        this.setState({ currentProject: currentProject });
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
                      defaultValue={currentProject.description}
                      onChange={(data) => {
                        currentProject.description = data.target.value!;
                        console.log(currentProject);
                        this.setState({ currentProject: currentProject });
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
              </Dialog>
            </div>
          );
        }}
      </AllProjectsComponent>
    );
  }
  addProject = () => {
    this.setState({ addOpen: true, currentProject: new ProjectEntry() });
  };
  handleClose = (state: any) => {
    console.log(state);
    if (state) {
      project.addProject(this.state.currentProject);
    }
    this.setState({ addOpen: false });
  };
}

export default withTheme(
  withStyles(styles as any)(withRouter(Projects as any))
);
