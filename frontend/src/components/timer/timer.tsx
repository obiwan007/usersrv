// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from "@apollo/react-common";
import { Button, CircularProgress, createStyles, FormControl, Grid, Hidden, IconButton, InputLabel, ListItemIcon, ListItemSecondaryAction, ListItemText, MenuItem, Select, TextField, Theme, WithStyles, withStyles } from "@material-ui/core";
import { green, red } from "@material-ui/core/colors";
import ListItem from "@material-ui/core/ListItem";
import { Delete, PlayArrow, Stop, Timer as TimerIcon } from "@material-ui/icons";
import React from "react";
import { withRouter } from "react-router-dom";
import Autosizer from "react-virtualized-auto-sizer";
import { FixedSizeList, ListChildComponentProps } from 'react-window';
import { AllProjectsComponent, AllTimerComponent, CreateTimerComponent, CreateTimerMutation, CreateTimerMutationVariables, DeleteTimerComponent, DeleteTimerMutation, DeleteTimerMutationVariables, Project, refetchAllTimerQuery, StartTimerComponent, StartTimerMutation, StartTimerMutationVariables, StopTimerComponent, StopTimerMutation, StopTimerMutationVariables, Timer as TimerEntry, UpdateTimerComponent } from "../../graphql";
import timer, { Timer as TimerSrv } from "../../lib/timer";
import theme from '../../theme';
import { RunningClock } from "./runningClock";
import TimerEdit from "./timerEdit";
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
    container: {
      paddingTop: spacing(1),
      paddingBottom: spacing(1),
      width: "100%"
    },
    formControl: {
      // margin: spacing(1),
      //paddingRight: spacing(1),
      //paddingLeft: spacing(1),
      //minWidth: 120,
      width: "100%",
    },
    list: {
      listStyle: "none",
    },
    topButtons: {
      marginLeft: spacing(2),
      paddingRight: spacing(1),
    },
    selectEmpty: {
      // marginTop: 13,
    },
    scroll: {
      listStyle: "none",
      overflowY: "auto",
      height: "calc(100vh - 230px)",
      [theme.breakpoints.down('sm')]: {
        height: "calc(100vh - 320px)",
      },
    },
  });

interface IState {
  startTime: Date;
  endTime: Date;
  isRunning: boolean;
  elapsed: number;
  sum: number;
  description: string;
  currentProject: string;
  currentTimer: TimerEntry;
  timefilter: string;
  isLoading: boolean;
  addOpen: boolean;
  editOpen: boolean;
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
      currentTimer: {},
      description: "",
      isRunning: timer.getTimer().isRunning,
      startTime: timer.getTimer().timerStart,
      endTime: timer.getTimer().timerEnd,
      elapsed: timer.getTimer().elapsed(),
      timefilter: "1",
      sum: 0,
      isLoading: false,
      addOpen: false,
      editOpen: false,
    };
  }

  componentDidMount() {
    this.setState({
      description: timer.getTimer().description,
      currentProject: timer.getTimer().project,
    });
    // this.interval = setInterval(() => {
    //   this.checkTimer();
    // }, 500);
  }
  componentWillUnmount() {
    //clearInterval(this.interval!);
  }

  render() {
    const { description, currentProject, timefilter, addOpen, editOpen } = this.state;

    let isLoading = false;

    const { classes } = this.props;
    let currentTimer: TimerEntry | undefined = undefined;

    let isRunning = false;
    let allTimer: any = [];
    return (
      <div>
        <AllProjectsComponent>
          {({ data, loading, error }) => {
            const allProjects = data;
            // if (columns) {
            //   const f = columns.find((c) => c.title === "Project");
            //   if (f) {
            //     f.lookup = this.projectsDict(
            //       allProjects?.allProjects as Project[]
            //     );
            //   }
            // }
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
                                <AllTimerComponent
                                  variables={{
                                    d: {
                                      dayrange: timefilter,
                                    },
                                  }}
                                >
                                  {({ data, loading, error }) => {
                                    // Any errors? Say so!
                                    currentTimer = undefined;
                                    // console.log("IsLoading1", this.state.isLoading, loading);
                                    isLoading = this.state.isLoading || loading;
                                    // console.log("IsLoading2", isLoading, loading);
                                    allTimer = data;
                                    const count = allTimer?.allTimer?.length;
                                    data?.allTimer?.forEach((d) => {
                                      if (d) {
                                        (d as any).projectId = d?.project?.id;
                                        if (d.isRunning === true) {
                                          currentTimer = d;
                                          isRunning = true;
                                        }
                                      }
                                    });

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
                                        <Grid className={classes.container} container alignItems="center" spacing={3}
                                        >
                                          <Grid item sm={6} xs={6}>

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
                                                label="Work Topic"
                                                type="text"
                                                fullWidth
                                              />
                                            </FormControl>
                                          </Grid>
                                          <Grid item sm={3} xs={6}>
                                            <FormControl
                                              className={[
                                                classes.formControl,
                                                classes.selectEmpty,
                                              ].join(" ")}
                                            >
                                              <InputLabel shrink id="demo-simple-select-placeholder-label-label">
                                                Project
                                              </InputLabel>
                                              <Select
                                                className={classes.selectEmpty}
                                                value={
                                                  allProjects &&
                                                    (allProjects.allProjects as any[])
                                                      .length > 0
                                                    ? currentProject
                                                    : ""
                                                }
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
                                          </Grid>
                                          <Grid item sm={1} xs={1}>
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
                                                            disabled={isLoading}
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
                                                            {isLoading && (
                                                              <CircularProgress></CircularProgress>
                                                            )}
                                                            {!isLoading &&
                                                              !currentTimer && (
                                                                <PlayArrow
                                                                  style={{
                                                                    color:
                                                                      green[500],
                                                                  }}
                                                                />
                                                              )}
                                                            {!isLoading &&
                                                              currentTimer && (
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
                                          </Grid>
                                          <Grid item sm={1} xs={2}>
                                            <FormControl
                                              className={classes.topButtons}
                                            >
                                              <IconButton
                                                disabled={!isRunning}
                                                onClick={() =>
                                                  this.deleteTimer(currentTimer!, deleteTimer)
                                                }
                                                edge="start"
                                                color="inherit"
                                                aria-label="menu"
                                              >
                                                <Delete />
                                              </IconButton>
                                            </FormControl>
                                          </Grid>
                                          <Grid item sm={1} xs={1}>
                                            <RunningClock
                                              currentTimer={currentTimer}
                                            ></RunningClock>
                                          </Grid>
                                        </Grid>
                                        {/* ----------------------- */}
                                        <Grid className={classes.container} container alignItems="center" spacing={3}
                                        >
                                          <Grid item sm={6} xs={6}>
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

                                                onChange={(event) => {
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
                                          </Grid>
                                          {/* <Box>
            <Typography variant="body2" color="textSecondary" align="center">
              {TimerSrv.hms(sum)}
            </Typography>
          </Box> */}
                                        </Grid>
                                        <div
                                          className={classes.scroll}                                          
                                        >
                                          <Autosizer>
                                            {({ height, width }) =>
                                              (<FixedSizeList
                                                height={height}
                                                width={width}
                                                itemCount={count}
                                                itemSize={80}
                                                itemData={{ allTimer, deleteTimer }}
                                              // style={{
                                              //   height: "calc(100vh - 445px)",
                                              //   minHeight: "calc(100vh - 445px)",
                                              // }}
                                              >

                                                {this.renderRow}
                                              </FixedSizeList>)
                                            }

                                          </Autosizer>
                                        </div>
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

        <TimerEdit timefilter={timefilter} timer={this.state.currentTimer} addOpen={addOpen} editOpen={editOpen} onClose={() => { this.setState({ addOpen: false, editOpen: false }) }}>

        </TimerEdit>
        {/* {this.renderDialog()} */}
      </div >
    );
  }
  renderRow = (props: ListChildComponentProps) => {
    const { index, style } = props;
    if (!props.data.allTimer) {
      return <></>
    }
    const entry = props.data.allTimer.allTimer[index];
    return (
      <div style={style} key={index}>
        <ListItem button onClick={() => this.setState({ currentTimer: Object.assign({}, entry), editOpen: true })}>
          <Hidden xsDown implementation="css">
            <ListItemIcon>
              <TimerIcon />
            </ListItemIcon>
          </Hidden>
          <ListItemText style={{ width: '50%' }} primary={entry.description} secondary={entry.project?.name} ></ListItemText>
          <ListItemText style={{ width: '30%' }} primary={this.toLocaleDate(entry.timerStart)} secondary={<>
            {entry.isRunning ? (
              <>
                {this.toTime(entry.timerStart)} - running
              </>
            ) : (
                <>
                  {this.toTime(entry.timerStart)} - {this.toTime(entry.timerEnd)}
                </>
              )}
          </>} >
          </ListItemText>
          <ListItemText style={{ width: '20%' }} primary={!entry.isRunning ? (
            <>{TimerSrv.hms(entry.elapsedSeconds)}</>
          ) : (
              "Running"
            )}></ListItemText>
          <ListItemSecondaryAction>
            <IconButton onClick={() => this.deleteTimer(entry, props.data.deleteTimer)} edge="end" aria-label="delete">
              <Delete />
            </IconButton>
            {/* <IconButton edge="end" aria-label="delete">
              <Delete />
            </IconButton> */}
          </ListItemSecondaryAction>
        </ListItem>
      </div>
    );
  }

  deleteTimer = (entity: TimerEntry, deleteTimer: MutationFunction<DeleteTimerMutation, DeleteTimerMutationVariables>) => {
    deleteTimer({
      refetchQueries: [
        refetchAllTimerQuery({ d: { dayrange: this.state.timefilter } }),
        refetchAllTimerQuery({ d: { dayrange: "0" } }),
      ],
      fetchPolicy: "no-cache",
      variables: {
        timerId: entity!.id!,
      },
    }).catch((ex) => {
      console.log(
        "Error in mutation",
        ex
      );
    });
  }

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

  handleDescriptionField = (event: any) => {
    timer.currentTimer.description = event.target.value!;
    timer.save();
    this.setState({ description: event.target.value });
    event.preventDefault();
  };
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
    console.log("StartStop CurrentTimer:", currentTimer);

    // New time will be created
    this.setState({ isLoading: true }, async () => {
      if (!currentTimer) {
        const newTimer = await createTimer({
          // refetchQueries: [
          //   refetchAllTimerQuery({ d: { dayrange: this.state.timefilter } }),
          // ],
          fetchPolicy: "no-cache",
          variables: {
            d: {
              description: timer.currentTimer.description, // newData.description,
              project: timer.currentTimer.project, // newData.projectId,
            },
          },
        }).catch((ex) => {
          console.log("Error in mutation", ex);
        });
        console.log("NewTimer", newTimer);
        if (newTimer) {
          await startTimer({
            refetchQueries: [
              refetchAllTimerQuery({ d: { dayrange: this.state.timefilter } }),
              refetchAllTimerQuery({ d: { dayrange: "0" } }),
            ],
            fetchPolicy: "no-cache",
            variables: {
              timerId: newTimer.data!.createTimer!.id!,
            },
          }).catch((ex) => {
            console.log("Error in mutation", ex);
          });
        }
      } else {
        await stopTimer({
          refetchQueries: [
            refetchAllTimerQuery({ d: { dayrange: this.state.timefilter } }),
            refetchAllTimerQuery({ d: { dayrange: "0" } }),
          ],
          fetchPolicy: "no-cache",
          variables: {
            timerId: currentTimer.id!,
          },
        }).catch((ex) => {
          console.log("Error in mutation", ex);
        });
      }
      this.setState({ isLoading: false });
    });

    // this.state.isRunning ? timer.endTimer() : timer.startTimer();
    // this.setState({
    //   currentProject: timer.getTimer().project,
    //   description: timer.getTimer().description,
    // });
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
