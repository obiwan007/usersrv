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
  list: ClientEntry[];
  addOpen: boolean;

  currentClient: ClientEntry;
}
interface IProps {
  history?: any;
}

export class Clients extends React.PureComponent<PROPS_WITH_STYLES, IState> {
  columns = [
    { title: "Name", field: "name" },
    { title: "Description", field: "description" },
  ];
  interval?: NodeJS.Timeout;

  /**
   *
   *
   */
  constructor(props: PROPS_WITH_STYLES, state: IState) {
    super(props, state);
    this.state = {
      list: client.Entries(),
      addOpen: false,
      currentClient: new ClientEntry(),
    };
  }

  componentDidMount() {}
  componentWillUnmount() {}
  render() {
    const { classes } = this.props;
    const { currentClient } = this.state;

    const { list, addOpen } = this.state;
    console.log("Current list", list);
    return (
      <div>
        <h3>Clients</h3>
        {/*         

        <Button onClick={() => this.addProject()}>Add Client</Button> */}

        <List component="nav" aria-label="main mailbox folders">
          <MaterialTable
            options={{
              minBodyHeight: "calc(100vh - 360px)",
              maxBodyHeight: "calc(100vh - 360px)",
            }}
            title="Clients"
            columns={this.columns}
            data={list?.map((u) => u) as any[]}
            editable={{
              isEditable: (rowData) => true, // only name(a) rows would be editable
              isDeletable: (rowData) => true, // only name(a) rows would be deletable
              onRowAdd: (newData) =>
                new Promise((resolve, reject) => {
                  setTimeout(() => {
                    {
                      const d = client.add(newData);
                      console.log("New List", d);
                      this.setState({ list: d }, () => resolve());
                    }
                    resolve();
                  }, 1000);
                }),
              onRowUpdate: (newData, oldData) =>
                new Promise((resolve, reject) => {
                  setTimeout(() => {
                    {
                      const d = client.update(newData);
                      console.log("New List", d);
                      this.setState({ list: d }, () => resolve());
                    }
                    resolve();
                  }, 1000);
                }),
              onRowDelete: (oldData) =>
                new Promise((resolve, reject) => {
                  setTimeout(() => {
                    {
                      const d = client.del(oldData);
                      console.log("New List", d);
                      this.setState({ list: d }, () => resolve());
                    }
                    resolve();
                  }, 1000);
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
          onClose={() => this.setState({ addOpen: false })}
          aria-labelledby="simple-modal-title"
          aria-describedby="simple-modal-description"
        >
          <DialogTitle id="simple-dialog-title">Add new Client</DialogTitle>
          <DialogContent>
            <DialogContentText>
              To subscribe to this website, please enter your email address
              here. We will send updates occasionally.
            </DialogContentText>
            <form noValidate autoComplete="off">
              <TextField
                required
                autoFocus
                defaultValue={currentClient.name}
                onChange={(data) => {
                  currentClient.name = data.target.value!;
                  console.log(currentClient);
                  this.setState({ currentClient: currentClient });
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
                  this.setState({ currentClient: currentClient });
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
    this.setState({ addOpen: true, currentClient: new ClientEntry() });
  };
  handleClose = (state: any) => {
    console.log(state);
    if (state) {
      client.add(this.state.currentClient);
    }
    this.setState({ addOpen: false, list: client.Entries() });
  };
}

export default withTheme(withStyles(styles as any)(withRouter(Clients as any)));
