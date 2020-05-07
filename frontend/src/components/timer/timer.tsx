// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from "@apollo/react-common";
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
  CreateTimerComponent,
  CreateTimerMutation,
  CreateTimerMutationVariables,
  DeleteTimerComponent,
  Project,
  refetchAllTimerQuery,
  StartTimerComponent,
  StartTimerMutation,
  StartTimerMutationVariables,
  StopTimerComponent,
  StopTimerMutation,
  StopTimerMutationVariables,
  Timer as TimerEntry,
  UpdateTimerComponent,
} from "../../graphql";
import project from "../../lib/project";
import timer, { Timer as TimerSrv } from "../../lib/timer";
import { RunningClock } from "./runningClock";
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
      description: "",
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
      currentProject: "",
    });
    // this.interval = setInterval(() => {
    //   this.checkTimer();
    // }, 500);
  }
  componentWillUnmount() {
    //clearInterval(this.interval!);
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
      columns,
      timefilter,
      sum,
    } = this.state;
    const { classes } = this.props;
    const seconds = timer.elapsed();
    let currentTimer: TimerEntry | undefined = undefined;

    const isRunning = false;
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
              }
            }
            return (
              <UpdateTimerComponent>
                {(updateTimer, { data }) => {
                  return (
                    <DeleteTimerComponent>
                      {(deleteTimer, { data }) => {
                        return (
                          <CreateTimerComponent>
                            {(createTimer, { data }) => {
                              return (
                                <AllTimerComponent>
                                  {({ data, loading, error }) => {
                                    // Any errors? Say so!
                                    currentTimer = undefined;
                                    data?.allTimer?.forEach((d) => {
                                      if (d) {
                                        (d as any).projectId = d?.project?.id;
                                        if (d.isRunning === true) {
                                          currentTimer = d;
                                        }
                                      }
                                    });
                                    console.log(
                                      "currentTimer",
                                      currentTimer,
                                      data?.allTimer
                                    );
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
                                              <StopTimerComponent>
                                                {(stopTimer, { data }) => {
                                                  return (
                                                    <StartTimerComponent>
                                                      {(
                                                        startTimer,
                                                        { data }
                                                      ) => {
                                                        return (
                                                          <IconButton
                                                            onClick={() =>
                                                              this.startStopTimer(
                                                                createTimer,
                                                                startTimer,
                                                                stopTimer,
                                                                currentTimer
                                                              )
                                                            }
                                                            edge="start"
                                                            color="secondary"
                                                            aria-label="menu"
                                                          >
                                                            {!currentTimer && (
                                                              <PlayArrow
                                                                style={{
                                                                  color:
                                                                    green[500],
                                                                }}
                                                              />
                                                            )}
                                                            {currentTimer && (
                                                              <Stop
                                                                style={{
                                                                  color:
                                                                    red[500],
                                                                }}
                                                              />
                                                            )}
                                                          </IconButton>
                                                        );
                                                      }}
                                                    </StartTimerComponent>
                                                  );
                                                }}
                                              </StopTimerComponent>
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
                                            <RunningClock
                                              currentTimer={currentTimer}
                                            ></RunningClock>
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
                                                updateTimer({
                                                  refetchQueries: [
                                                    refetchAllTimerQuery(),
                                                  ],
                                                  fetchPolicy: "no-cache",
                                                  variables: {
                                                    d: {
                                                      id: newData.id,
                                                      description:
                                                        newData.description,
                                                      project:
                                                        newData.projectId,
                                                    },
                                                  },
                                                }).catch((ex) => {
                                                  console.log(
                                                    "Error in mutation",
                                                    ex
                                                  );
                                                }),
                                              onRowDelete: (oldData) =>
                                                deleteTimer({
                                                  refetchQueries: [
                                                    refetchAllTimerQuery(),
                                                  ],
                                                  fetchPolicy: "no-cache",
                                                  variables: {
                                                    timerId: oldData.id,
                                                  },
                                                }).catch((ex) => {
                                                  console.log(
                                                    "Error in mutation",
                                                    ex
                                                  );
                                                }),
                                            }}
                                          />
                                        </List>
                                      </>
                                    );
                                  }}
                                </AllTimerComponent>
                              );
                            }}
                          </CreateTimerComponent>
                        );
                      }}
                    </DeleteTimerComponent>
                  );
                }}
              </UpdateTimerComponent>
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
        editable: () => false,
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
  // showElapsed(t: TimerEntry | null | unknown): string {
  //   if (t) {
  //     const time1 = new Date(t.timerStart!);
  //     const time2 = new Date();
  //     return TimerSrv.hms((time2.getTime() - time1.getTime()) / 1000);
  //   }
  //   return "00:00:00";
  // }
  async startStopTimer(
    createTimer: MutationFunction<
      CreateTimerMutation,
      CreateTimerMutationVariables
    >,
    startTimer: MutationFunction<
      StartTimerMutation,
      StartTimerMutationVariables
    >,
    stopTimer: MutationFunction<StopTimerMutation, StopTimerMutationVariables>,
    currentTimer?: TimerEntry
  ) {
    console.log("CurrentTimer:", currentTimer);
    if (!currentTimer) {
      const newTimer = await createTimer({
        refetchQueries: [refetchAllTimerQuery()],
        fetchPolicy: "no-cache",
        variables: {
          d: {
            description: "test", // newData.description,
            project: "0", // newData.projectId,
          },
        },
      }).catch((ex) => {
        console.log("Error in mutation", ex);
      });
      console.log("NewTimer", newTimer);
      if (newTimer) {
        startTimer({
          refetchQueries: [refetchAllTimerQuery()],
          fetchPolicy: "no-cache",
          variables: {
            timerId: newTimer.data!.createTimer!.id!,
          },
        }).catch((ex) => {
          console.log("Error in mutation", ex);
        });
      }
    } else {
      stopTimer({
        refetchQueries: [refetchAllTimerQuery()],
        fetchPolicy: "no-cache",
        variables: {
          timerId: currentTimer.id!,
        },
      }).catch((ex) => {
        console.log("Error in mutation", ex);
      });
    }

    // this.state.isRunning ? timer.endTimer() : timer.startTimer();
    // this.setState({
    //   currentProject: timer.getTimer().project,
    //   description: timer.getTimer().description,
    // });
  }
  discardTimer() {
    timer.discardTimer();
    this.setState({
      description: timer.getTimer().description,
    });
  }

  checkTimer = () => {
    this.setState({
      elapsed: new Date().getTime(),
    });
  };

  // recalcSummary( TimeEntry[]): number {
  //   let sum = 0;
  //   list.forEach((l) => (sum += l.elapsedSeconds));
  //   return sum;
  // }
}

export default withStyles(styles as any)(withRouter(Timer as any));
