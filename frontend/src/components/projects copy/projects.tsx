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
  list: ProjectEntry[];
  addOpen: boolean;

  currentProject: ProjectEntry;
}
interface IProps {
  history?: any;
}

export class Projects extends React.PureComponent<PROPS_WITH_STYLES, IState> {
  columns = [
    { title: "Project", field: "project" },
    { title: "Client", field: "client" },
    { title: "Status", field: "status" },
    { title: "Seconds", field: "elapsedSeconds" },
    { title: "Team", field: "team" },
  ];
  interval?: NodeJS.Timeout;

  /**
   *
   *
   */
  constructor(props: PROPS_WITH_STYLES, state: IState) {
    super(props, state);
    this.state = {
      list: project.Entries(),
      addOpen: false,
      currentProject: new ProjectEntry(),
    };
  }

  componentDidMount() {}
  componentWillUnmount() {}
  render() {
    const { classes } = this.props;
    const { currentProject } = this.state;
    console.log(this.props);
    const { list, addOpen } = this.state;
    return (
      <div>
        <h3>Projects</h3>
        <Button onClick={() => this.addProject()}>Add Project</Button>

        <List component="nav" aria-label="main mailbox folders">
          <MaterialTable
            options={{
              minBodyHeight: "calc(100vh - 360px)",
              maxBodyHeight: "calc(100vh - 360px)",
            }}
            title="Projects"
            columns={this.columns}
            data={list?.map((u) => u) as any[]}
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
          onClose={() => this.setState({ addOpen: false })}
          aria-labelledby="simple-modal-title"
          aria-describedby="simple-modal-description"
        >
          <DialogTitle id="simple-dialog-title">Add new Project</DialogTitle>
          <DialogContent>
            <DialogContentText>
              To subscribe to this website, please enter your email address
              here. We will send updates occasionally.
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
            <Button onClick={() => this.handleClose(false)} color="primary">
              Cancel
            </Button>
            <Button onClick={() => this.handleClose(true)} color="primary">
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
  }
  addProject = () => {
    this.setState({ addOpen: true, currentProject: new ProjectEntry() });
  };
  handleClose = (state: any) => {
    console.log(state);
    if (state) {
      project.addProject(this.state.currentProject);
    }
    this.setState({ addOpen: false, list: project.Entries() });
  };
}

export default withTheme(
  withStyles(styles as any)(withRouter(Projects as any))
);
