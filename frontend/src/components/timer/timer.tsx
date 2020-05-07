// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import {
  Box,
  Button,
  createStyles,
  FormControl,
  IconButton,
  InputLabel,
  List,
  MenuItem,
  Select,
  TextField,
  Theme,
  Typography,
  WithStyles,
  withStyles,
} from "@material-ui/core";
import { green, red } from "@material-ui/core/colors";
import { Delete, PlayArrow, Stop } from "@material-ui/icons";
import MaterialTable from "material-table";
import React from "react";
import { withRouter } from "react-router-dom";
import {
  AllProjectsComponent,
  AllTimerComponent,
  CreateProjectComponent,
  DeleteProjectComponent,
  Project,
  Timer as TimerEntry,
  UpdateProjectComponent,
} from "../../graphql";
import project from "../../lib/project";
import timer, { Timer as TimerSrv } from "../../lib/timer";
// ----------------------------------------------------------------------------

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
    formControl: {
      margin: spacing(1),
      paddingRight: spacing(1),
      paddingLeft: spacing(1),
      minWidth: 120,
      width: "100%",
    },
    topButtons: {
      marginLeft: spacing(2),
      paddingRight: spacing(1),
    },
    selectEmpty: {
      // marginTop: 13,
    },
  });

interface IState {
  startTime: Date;
  endTime: Date;
  isRunning: boolean;
  elapsed: number;
  sum: number;
  columns: any[];
  description: string;
  currentProject: string;
  timefilter: string;
}
interface IProps {
  history?: any;
}
export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class Timer extends React.PureComponent<PROPS_WITH_STYLES, IState> {
  interval?: NodeJS.Timeout;

  filterSelect: any = [
    // { key: "0", value: "all" },
    { key: "1", value: "Today" },
    { key: "2", value: "Yesterday" },
    { key: "7", value: "Week" },
    { key: "30", value: "Month" },
  ];
  /**
   *
   *
   */
  constructor(props: PROPS_WITH_STYLES, state: IState) {
    super(props, state);
    this.state = {
      currentProject: timer.getTimer().project,
      description: timer.getTimer().description,
      isRunning: timer.getTimer().isRunning,
      startTime: timer.getTimer().timerStart,
      endTime: timer.getTimer().timerEnd,
      elapsed: timer.getTimer().elapsed(),
      columns: [],
      timefilter: "1",
      sum: 0,
    };
  }

  componentDidMount() {
    this.setState({
      columns: this.getColumns(),
      description: timer.getTimer().description,
      currentProject: timer.getTimer().project,
    });
    this.interval = setInterval(() => {
      this.checkTimer();
    }, 500);
  }
  componentWillUnmount() {
    clearInterval(this.interval!);
  }
  handleDescriptionField = (event: any) => {
    timer.currentTimer.description = event.target.value!;
    timer.save();
    this.setState({ description: event.target.value });
    event.preventDefault();
  };

  render() {
    const {
      description,
      currentProject,
      isRunning,
      columns,
      timefilter,
      sum,
    } = this.state;
    const { classes } = this.props;
    const seconds = timer.elapsed();
    return (
      <div>
        <h3>Timer</h3>
        <AllProjectsComponent>
          {({ data, loading, error }) => {
            const allProjects = data;
            if (columns) {
              const f = columns.find((c) => c.title === "Project");
              if (f) {
                f.lookup = this.projectsDict(
                  allProjects?.allProjects as Project[]
                );
                console.log("PROJECTS");
              }
            }
            console.log("AllProjects", allProjects);
            return (
              <UpdateProjectComponent>
                {(updateProject, { data }) => {
                  console.log("Updateclient MutationData", data);
                  return (
                    <DeleteProjectComponent>
                      {(deleteProject, { data }) => {
                        return (
                          <CreateProjectComponent>
                            {(createProject, { data }) => {
                              console.log("MutationData", data);
                              return (
                                <AllTimerComponent>
                                  {({ data, loading, error }) => {
                                    // Any errors? Say so!
                                    data?.allTimer?.forEach((d) => {
                                      d &&
                                        ((d as any).projectId = d?.project?.id);
                                    });
                                    console.log("Returned data", data);
                                    if (error) {
                                      return (
                                        <div>
                                          <h1>
                                            Error retrieving Timer list &mdash;{" "}
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
                                      <>
                                        <Box
                                          display="flex"
                                          flexDirection="row"
                                          alignItems="center"
                                        >
                                          <Box flexGrow={1}>
                                            <FormControl
                                              className={classes.formControl}
                                            >
                                              <TextField
                                                required
                                                autoFocus
                                                value={description}
                                                onChange={
                                                  this.handleDescriptionField
                                                }
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
                                              className={[
                                                classes.formControl,
                                                classes.selectEmpty,
                                              ].join(" ")}
                                            >
                                              <InputLabel>Project</InputLabel>
                                              <Select
                                                className={classes.selectEmpty}
                                                label="Project"
                                                value={currentProject}
                                                onChange={(event) => {
                                                  console.log(
                                                    "Projectselection:",
                                                    event.target
                                                  );
                                                  timer.currentTimer.project = event
                                                    .target.value as string;
                                                  timer.save();
                                                  this.setState({
                                                    currentProject:
                                                      timer.currentTimer
                                                        .project,
                                                  });
                                                }}
                                                // inputProps={{
                                                //   name: "age",
                                                //   id: "age-native-simple",
                                                // }}
                                              >
                                                <MenuItem
                                                  aria-label="None"
                                                  value=""
                                                />
                                                {allProjects &&
                                                  allProjects.allProjects?.map(
                                                    (e: any) => (
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
                                          </Box>

                                          <Box alignItems="center">
                                            <FormControl
                                              className={classes.topButtons}
                                            >
                                              <IconButton
                                                onClick={() =>
                                                  this.startStopTimer()
                                                }
                                                edge="start"
                                                color="secondary"
                                                aria-label="menu"
                                              >
                                                {!isRunning && (
                                                  <PlayArrow
                                                    style={{
                                                      color: green[500],
                                                    }}
                                                  />
                                                )}
                                                {isRunning && (
                                                  <Stop
                                                    style={{ color: red[500] }}
                                                  />
                                                )}
                                              </IconButton>
                                            </FormControl>
                                          </Box>
                                          <Box>
                                            <FormControl
                                              className={classes.topButtons}
                                            >
                                              <IconButton
                                                disabled={!isRunning}
                                                onClick={() =>
                                                  this.discardTimer()
                                                }
                                                edge="start"
                                                color="inherit"
                                                aria-label="menu"
                                              >
                                                <Delete />
                                              </IconButton>
                                            </FormControl>
                                          </Box>
                                          <Box>
                                            <Typography
                                              variant="body2"
                                              color="textSecondary"
                                              align="center"
                                            >
                                              xxxxx
                                              {/* {isRunning
                                                ? timer.currentTimer.hms()
                                                : "00:00:00"} */}
                                            </Typography>
                                          </Box>
                                        </Box>
                                        {/* ----------------------- */}
                                        <Box
                                          display="flex"
                                          flexDirection="row"
                                          alignItems="center"
                                        >
                                          <Box>
                                            <FormControl
                                              style={{ width: 250 }}
                                              className={[
                                                classes.formControl,
                                                classes.selectEmpty,
                                              ].join(" ")}
                                            >
                                              <InputLabel>Filter</InputLabel>
                                              <Select
                                                className={classes.selectEmpty}
                                                value={timefilter}
                                                // renderValue={(value) =>
                                                //   `${
                                                //     this.filterSelect.find(
                                                //       (f: any) =>
                                                //         f.key === value
                                                //     ).value
                                                //   }:  ${TimerSrv.hms(sum)}`
                                                // }
                                                onChange={(event) => {
                                                  const list = timer.Entries(
                                                    event.target
                                                      .value! as string
                                                  );
                                                  this.setState({
                                                    timefilter: event.target
                                                      .value! as string,
                                                  });
                                                }}
                                                // inputProps={{
                                                //   name: "age",
                                                //   id: "age-native-simple",
                                                // }}
                                              >
                                                {this.filterSelect.map(
                                                  (f: any) => (
                                                    <MenuItem
                                                      aria-label="None"
                                                      value={f.key}
                                                      key={f.key}
                                                    >
                                                      {f.value}
                                                    </MenuItem>
                                                  )
                                                )}
                                              </Select>
                                            </FormControl>
                                          </Box>
                                          {/* <Box>
            <Typography variant="body2" color="textSecondary" align="center">
              {TimerSrv.hms(sum)}
            </Typography>
          </Box> */}
                                        </Box>
                                        <List
                                          component="nav"
                                          aria-label="main mailbox folders"
                                        >
                                          <MaterialTable
                                            isLoading={loading}
                                            options={{
                                              actionsColumnIndex: -1,
                                              padding: "dense",
                                              sorting: true,
                                              minBodyHeight:
                                                "calc(100vh - 460px)",
                                              maxBodyHeight:
                                                "calc(100vh - 460px)",
                                              pageSize: 20,
                                              pageSizeOptions: [10, 20, 100],
                                            }}
                                            title="Timetable"
                                            columns={columns}
                                            data={
                                              data?.allTimer?.map(
                                                (u) => u
                                              ) as any[]
                                            }
                                            editable={{
                                              isEditable: (rowData) => {
                                                return true;
                                              }, // only name(a) rows would be editable
                                              isDeletable: (rowData) => true, // only name(a) rows would be deletable
                                              onRowUpdate: (newData, oldData) =>
                                                new Promise(
                                                  (resolve, reject) => {
                                                    setTimeout(() => {
                                                      {
                                                        const d = timer.update(
                                                          newData,
                                                          timefilter
                                                        );
                                                        console.log(
                                                          "New List",
                                                          d
                                                        );
                                                      }
                                                      resolve();
                                                    }, 1000);
                                                  }
                                                ),
                                              onRowDelete: (oldData) =>
                                                new Promise(
                                                  (resolve, reject) => {
                                                    setTimeout(() => {
                                                      {
                                                        const d = timer.del(
                                                          oldData,
                                                          timefilter
                                                        );
                                                        console.log(
                                                          "New List",
                                                          d
                                                        );
                                                      }
                                                      resolve();
                                                    }, 1000);
                                                  }
                                                ),
                                            }}
                                          />
                                        </List>
                                      </>
                                    );
                                  }}
                                </AllTimerComponent>
                              );
                            }}
                          </CreateProjectComponent>
                        );
                      }}
                    </DeleteProjectComponent>
                  );
                }}
              </UpdateProjectComponent>
            );
          }}
        </AllProjectsComponent>
      </div>
    );
  }
  getColumns = () => {
    const columns = [
      { width: "40%", title: "Description", field: "description" },
      // {
      //   title: "Client",
      //   field: "client",
      //   editable: () => true,
      //   lookup: client.EntriesDict(),
      // },
      {
        title: "Project",
        field: "projectId",
        width: "350px",
        editable: () => true,
        lookup: project.EntriesDict(),
      },
      {
        title: "Date",
        field: "tStart",
        width: "150px",
        defaultSort: "desc",
        editable: () => false,
        render: (data: TimerEntry) => {
          return <>{this.toLocaleDate(data.timerStart)}</>;
        },
      },
      {
        title: "Time",
        field: "tStart",
        width: "280px",
        editable: () => false,
        defaultSort: "desc",
        render: (data: TimerEntry) => {
          return (
            <>
              {this.toTime(data.timerStart)} - {this.toTime(data.timerEnd)}
            </>
          );
        },
      },
      {
        title: "Seconds",
        field: "elapsedSeconds",
        width: "100px",
        render: (data: TimerEntry) => {
          return <>{TimerSrv.hms(data.elapsedSeconds)}</>;
        },
      },
    ];
    return columns;
  };
  toTime(d: string | null | undefined): string {
    if (!d) {
      return "";
    }
    return new Date(d).toLocaleTimeString();
  }

  toLocaleDate(d: string | null | undefined): string {
    if (!d) {
      return "";
    }
    return new Date(d).toLocaleDateString();
  }
  projectsDict(data: Project[] | null | undefined): any {
    if (!data) {
      return {};
    }
    const dict: any = {};
    data.forEach((e) => e.id && (dict[+e.id] = e.name));
    return dict;
  }

  startStopTimer() {
    this.state.isRunning ? timer.endTimer() : timer.startTimer();
    this.setState({
      currentProject: timer.getTimer().project,
      description: timer.getTimer().description,
    });
  }
  discardTimer() {
    timer.discardTimer();
    this.setState({
      description: timer.getTimer().description,
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

  // recalcSummary(list: TimeEntry[]): number {
  //   let sum = 0;
  //   list.forEach((l) => (sum += l.elapsedSeconds));
  //   return sum;
  // }
}

export default withStyles(styles as any)(withRouter(Timer as any));
