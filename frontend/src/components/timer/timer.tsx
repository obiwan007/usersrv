// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import {
  Box,
  createStyles,
  FormControl,
  IconButton,
  InputLabel,
  List,
  Select,
  TextField,
  Theme,
  Typography,
  withStyles,
  WithStyles,
} from "@material-ui/core";
import { PlayArrow, Stop } from "@material-ui/icons";
import MaterialTable from "material-table";
import React from "react";
import { withRouter } from "react-router-dom";
import client from "../../lib/client";
import project from "../../lib/project";
import timer, { TimeEntry } from "../../lib/timer";
// ----------------------------------------------------------------------------

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
    formControl: {
      margin: spacing(1),
      paddingRight: spacing(1),
      minWidth: 120,
      width: "100%",
    },
    selectEmpty: {
      marginTop: 13,
    },
  });

interface IState {
  startTime: Date;
  endTime: Date;
  isRunning: boolean;
  elapsed: number;
  list: TimeEntry[];
  columns: any[];
}
interface IProps {
  history?: any;
}
export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class Timer extends React.PureComponent<PROPS_WITH_STYLES, IState> {
  interval?: NodeJS.Timeout;

  /**
   *
   *
   */
  constructor(props: PROPS_WITH_STYLES, state: IState) {
    super(props, state);
    this.state = {
      isRunning: timer.getTimer().isRunning,
      startTime: timer.getTimer().timerStart,
      endTime: timer.getTimer().timerEnd,
      elapsed: timer.getTimer().elapsed(),
      list: timer.Entries(),
      columns: [],
    };
  }

  componentDidMount() {
    this.setState({ columns: this.getColumns() });
    this.interval = setInterval(() => {
      this.checkTimer();
    }, 500);
  }
  componentWillUnmount() {
    clearInterval(this.interval!);
  }
  render() {
    const { isRunning, list, columns } = this.state;
    const { classes } = this.props;
    const seconds = timer.elapsed();
    return (
      <div>
        <h3>Timer</h3>
        <Box display="flex" flexDirection="row">
          <Box flexGrow={1}>
            <FormControl className={classes.formControl}>
              <TextField
                required
                autoFocus
                defaultValue={timer.currentTimer.description}
                onChange={(data) => {
                  timer.currentTimer.description = data.target.value!;
                  timer.save();
                }}
                margin="dense"
                id="name"
                label="What are your working on"
                type="text"
                fullWidth
              />
            </FormControl>
          </Box>
          <Box>
            <FormControl
              className={[classes.formControl, classes.selectEmpty].join(" ")}
            >
              <InputLabel>Project</InputLabel>
              <Select
                className={classes.selectEmpty}
                label="Project"
                native
                defaultValue={timer.currentTimer.project}
                onChange={(event) => {
                  console.log("Projectselection:", event.target);
                  timer.currentTimer.project = event.target.value as string;

                  timer.save();
                }}
                // inputProps={{
                //   name: "age",
                //   id: "age-native-simple",
                // }}
              >
                <option aria-label="None" value="" />
                {project.Entries().map((e) => (
                  <option value={e.id}>{e.name}</option>
                ))}
              </Select>
            </FormControl>
          </Box>
          <Box>
            <FormControl
              className={[classes.formControl, classes.selectEmpty].join(" ")}
            >
              <IconButton
                onClick={() => this.startStopTimer()}
                edge="start"
                color="inherit"
                aria-label="menu"
              >
                {!isRunning && <PlayArrow />}
                {isRunning && <Stop />}
              </IconButton>
            </FormControl>
          </Box>
          <Box>
            <Typography variant="body2" color="textSecondary" align="center">
              {timer.currentTimer.hms()}
            </Typography>
          </Box>
        </Box>

        <List component="nav" aria-label="main mailbox folders">
          <MaterialTable
            options={{
              sorting: true,
              minBodyHeight: "calc(100vh - 380px)",
              maxBodyHeight: "calc(100vh - 380px)",
              pageSize: 20,
              pageSizeOptions: [10, 20, 100],
            }}
            title="Timetable"
            columns={columns}
            data={list?.map((u) => u) as any[]}
            editable={{
              isEditable: (rowData) => {
                return true;
              }, // only name(a) rows would be editable
              isDeletable: (rowData) => true, // only name(a) rows would be deletable
              onRowUpdate: (newData, oldData) =>
                new Promise((resolve, reject) => {
                  setTimeout(() => {
                    {
                      const d = timer.update(newData);
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
                      const d = timer.del(oldData);
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
      </div>
    );
  }
  getColumns = () => {
    const columns = [
      { title: "Description", field: "description" },
      {
        title: "Client",
        field: "client",
        editable: () => true,
        lookup: client.EntriesDict(),
      },
      {
        title: "Project",
        field: "project",
        editable: () => true,
        lookup: project.EntriesDict(),
      },
      {
        title: "Start",
        field: "tStart",
        defaultSort: "desc",
        render: (data: TimeEntry) => {
          return (
            <>
              {data.tStart}-{data.tEnd}
              <br></br>
              {data.timerStart.toLocaleDateString()}
            </>
          );
        },
      },
      {
        title: "Seconds",
        field: "elapsedSeconds",
        render: (data: TimeEntry) => {
          return <>{data.hms()}</>;
        },
      },
    ];
    return columns;
  };
  startStopTimer() {
    this.state.isRunning ? timer.endTimer() : timer.startTimer();
    this.setState({
      list: timer.Entries(),
    });
  }

  checkTimer = () => {
    this.setState({
      isRunning: timer.getTimer().isRunning,
      startTime: timer.getTimer().timerStart,
      endTime: timer.getTimer().timerEnd,
      elapsed: timer.getTimer().elapsed(),
    });
  };
}

export default withStyles(styles as any)(withRouter(Timer as any));
