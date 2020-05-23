// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from "@apollo/react-common";
import DateFnsUtils from '@date-io/date-fns';
import { Box, Button, createStyles, FormControl, Hidden, IconButton, InputLabel, ListItemIcon, ListItemSecondaryAction, ListItemText, MenuItem, Select, Theme, WithStyles, withStyles } from "@material-ui/core";
import ListItem from "@material-ui/core/ListItem";
import { Delete, Timer as TimerIcon } from "@material-ui/icons";
import { KeyboardDatePicker, MuiPickersUtilsProvider } from '@material-ui/pickers';
import moment from "moment";
import React from "react";
import { withRouter } from "react-router-dom";
import Autosizer from "react-virtualized-auto-sizer";
import { FixedSizeList, ListChildComponentProps } from 'react-window';
import { AllTimerComponent, DeleteTimerComponent, DeleteTimerMutation, DeleteTimerMutationVariables, Project, refetchAllTimerQuery, Timer as TimerEntry } from "../../graphql";
import { Timer as TimerSrv } from "../../lib/timer";
import ProjectSelect from "../projects/projectSelect";
import TimerEdit from "../timer/timerEdit";
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
      // paddingRight: spacing(1),
      // paddingLeft: spacing(1),
      minWidth: 120,
      width: "100%",
    },
    container: {
      paddingTop: spacing(2),
    },
    box: {
      paddingRight: spacing(1),
      paddingLeft: spacing(2),
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
  });

interface IState {
  sum: number;
  timefilter: string;
  addOpen: boolean;
  editOpen: boolean;
  filterProject: Project;
  currentTimer: TimerEntry;
  filterTimerEnd: any;
  filterTimerStart: any;
}
interface IProps {
  history?: any;
}

type TimerMoment = TimerEntry & { 
  t1: moment.Moment,
  t2: moment.Moment,
}

export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class TimerList extends React.PureComponent<PROPS_WITH_STYLES, IState> {

  filterSelect: any = [
    // { key: "0", value: "all" },
    { key: "1", value: "Today" },
    { key: "2", value: "Yesterday" },
    { key: "7", value: "Week" },
    { key: "30", value: "Month" },
    { key: "90", value: "3 Month" },
    { key: "thismonth", value: "This Month" },
    { key: "lastmonth", value: "Last Month" },
  ];
  /**
   *
   *
   */
  constructor(props: PROPS_WITH_STYLES, state: IState) {
    super(props, state);
    this.state = {
      timefilter: "7",
      sum: 0,
      addOpen: false,
      editOpen: false,
      filterProject: {},
      currentTimer: {},
      filterTimerEnd: null,
      filterTimerStart: null,
    };
  }

  componentDidMount() {
    // this.setState({
    //   filterProject: {},
    // });
    // this.interval = setInterval(() => {
    //   this.checkTimer();
    // }, 500);
    this.setFilterTimerange("7")
  }
  componentWillUnmount() {
    //clearInterval(this.interval!);
  }

  render() {
    const { filterProject, filterTimerStart, filterTimerEnd, timefilter, addOpen, editOpen } = this.state;
    const { classes } = this.props;
    let allTimer: TimerMoment[] | null | undefined = [];

    return (
      <div>

        <DeleteTimerComponent>
          {(deleteTimer, { data }) => {
            return (
              <AllTimerComponent
                variables={{
                  d: {
                    dayrange: "100",
                  },
                }}
              >
                {({ data, loading, error }) => {
                  // Any errors? Say so!
                  // allTimer = _.cloneDeep(data);

                  allTimer = data?.allTimer ? data?.allTimer?.map(t => {
                    return {
                      ...t,
                      t1: moment(t?.timerStart!),
                      t2: moment(t?.timerEnd!)
                    };
                  }) : [];

                  console.log(allTimer)
                  if (filterProject?.id && !loading && allTimer) {
                    allTimer = allTimer.filter((a: TimerMoment) => a?.project?.id === filterProject.id);
                  }
                  if (filterTimerStart && filterTimerEnd) {
                    allTimer = allTimer?.filter((a: TimerMoment) => a.t1.isBefore(filterTimerEnd) && a.t1.isAfter(filterTimerStart) );
                  }

                  const count = allTimer ? allTimer.length : 0;
                  console.log("C", count)
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
                        className={classes.container}
                      >
                        <Box className={classes.box}>
                          <FormControl
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
                                this.setFilterTimerange(event.target.value! as string)

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


                        <Box className={classes.box}>
                          <FormControl
                            className={[
                              classes.formControl,
                              classes.selectEmpty,
                            ].join(" ")}
                          >
                            <ProjectSelect
                              project={filterProject}
                              onChanged={(p: Project) => this.setState({ filterProject: p })}
                            >
                            </ProjectSelect>
                          </FormControl>
                        </Box>
                        <MuiPickersUtilsProvider utils={DateFnsUtils}>
                          <Box className={classes.box}>
                            <KeyboardDatePicker
                              margin="dense"
                              id="date-picker-dialog"
                              label="Start"
                              value={filterTimerStart}
                              onChange={(date) => {
                                let filterTimerStart = date?.toISOString()!;

                                this.setState({
                                  filterTimerStart,
                                });
                              }
                              }
                              KeyboardButtonProps={{
                                'aria-label': 'change date',
                              }}
                            />
                          </Box>
                          <Box className={classes.box}>
                            <KeyboardDatePicker
                              margin="dense"
                              id="time-picker"
                              label="End"
                              value={filterTimerEnd}
                              onChange={(date) => {
                                let filterTimerEnd = date?.toISOString()!;
                                console.log(filterTimerEnd);
                                this.setState({
                                  filterTimerEnd: filterTimerEnd,
                                });
                              }
                              }
                              KeyboardButtonProps={{
                                'aria-label': 'change time',
                              }}
                            />


                          </Box>
                        </MuiPickersUtilsProvider>
                      </Box>
                      <div
                        className={classes.list}
                        style={{
                          height: "calc(100vh - 200px)",
                          minHeight: "calc(100vh - 200px)",
                        }}
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
                  )
                }}
              </AllTimerComponent>

            );
          }}
        </DeleteTimerComponent>


        <TimerEdit timefilter={timefilter} timer={this.state.currentTimer} addOpen={addOpen} editOpen={editOpen} onClose={() => { this.setState({ addOpen: false, editOpen: false }) }}>

        </TimerEdit>
      </div >
    );
  }
  renderRow = (props: ListChildComponentProps) => {
    const { index, style } = props;
    if (!props.data.allTimer) {
      return <></>
    }
    const entry = props.data.allTimer[index];
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

  setFilterTimerange(days: string) {
    let t1 = moment();
    let t2 = moment().add("days", -100);

    switch (days) {
      case "1":
        t2 = moment().hour(0);
        break;
      case "2":
        t2 = moment().add("days", -1);
        break;
      case "7":
        t2 = moment().add("days", -7);
        break;
      case "30":
        t2 = moment().add("days", -30);
        break;
      case "90":
        t2 = moment().add("days", -90);
        break;
      case "thismonth":
        t2 = moment().date(1);
        break;
      case "lastmonth":
        t1 = moment().date(1).add("days", -1);
        t2 = moment().add("month", -1).date(1);
        break;

    }

    this.setState({
      timefilter: days, filterTimerStart: t2, filterTimerEnd: t1
    });
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

  // recalcSummary( TimeEntry[]): number {
  //   let sum = 0;
  //   list.forEach((l) => (sum += l.elapsedSeconds));
  //   return sum;
  // }
}

export default withStyles(styles as any)(withRouter(TimerList as any));
