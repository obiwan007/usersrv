// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from "@apollo/react-common";
import DateFnsUtils from '@date-io/date-fns';
import { Button, Collapse, createStyles, FormControl, FormControlLabel, Grid, InputLabel, List, ListItemIcon, ListItemSecondaryAction, ListItemText, MenuItem, Select, Switch, Theme, Typography, WithStyles, withStyles } from "@material-ui/core";
import ListItem from "@material-ui/core/ListItem";
import { Assignment, ExpandLess, ExpandMore, Money as MoneyIcon, Timer as TimerIcon } from "@material-ui/icons";
import { KeyboardDatePicker, MuiPickersUtilsProvider } from '@material-ui/pickers';
import * as _ from "lodash";
import * as Moment from 'moment';
import { extendMoment } from 'moment-range';
import React from "react";
import { withRouter } from "react-router-dom";
import { AllTimerComponent, DeleteTimerMutation, DeleteTimerMutationVariables, Project, refetchAllTimerQuery, Timer as TimerEntry } from "../../graphql";
import { Timer as TimerSrv } from "../../lib/timer";
import theme from '../../theme';
import ProjectSelect from "../projects/projectSelect";
import BarChart from "./barchart";

const moment = extendMoment(Moment);
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
      // margin: spacing(2),
      // paddingRight: spacing(1),
      // paddingLeft: spacing(1),
      // minWidth: 120,
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
    nested: {
      paddingLeft: theme.spacing(4),
    },
    scroll: {
      overflowY: "auto",
      height: "calc(100vh - 248px)",
      [theme.breakpoints.down('sm')]: {
        height: "calc(100vh - 325px)",
      },
    },
  });

interface IState {
  sum: number;
  timefilter: string;
  addOpen: boolean;
  editOpen: boolean;
  filterProject?: Project;
  currentTimer: TimerEntry;
  filterTimerEnd: any;
  filterTimerStart: any;
  filterIsBilled: boolean;
  filterIsUnbilled: boolean;
  isOpen: { [id: string]: boolean };
}
interface IProps {
  history?: any;
}

type TimerMoment = TimerEntry & {
  t1: Moment.Moment,
  t2: Moment.Moment,
}

class ProjectGroup {
  project?: Project;
  timeEntries: TimerMoment[] = [];
  elapsed: number = 0;
  ranges: number[] = [];
  dates: Moment.Moment[] = [];
}

export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class Summary extends React.PureComponent<PROPS_WITH_STYLES, IState> {


  options: any = {
    tooltips: {
      callbacks: {
        label: (tooltipItem: any, data: any) => {
          return TimerSrv.hms(tooltipItem.yLabel * 3600);;
        }
      }
    }
  };

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
      currentTimer: {},
      filterTimerEnd: null,
      filterTimerStart: null,
      filterIsBilled: false,
      filterIsUnbilled: true,
      isOpen: {},
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
    const { isOpen, filterIsBilled, filterIsUnbilled, filterProject, filterTimerStart, filterTimerEnd, timefilter, addOpen, editOpen } = this.state;
    const { classes } = this.props;
    let allTimer: TimerMoment[] | null | undefined = [];

    let allProjects: { [id: string]: ProjectGroup } = {}
    let open = true;
    let barchart = {};
    return (
      <div>
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
              allTimer = allTimer?.filter((a: TimerMoment) => a.t1.isBefore(filterTimerEnd) && a.t1.isAfter(filterTimerStart));
            }
            if (filterIsBilled == false || filterIsUnbilled === false) {
            allTimer = allTimer?.filter((a: TimerMoment) => 
            (filterIsBilled && (a.isBilled === true))
            || (filterIsUnbilled && (a.isBilled === false))
            
            );
            }
            allTimer.forEach(t => {
              const id = t.project?.id!;
              if (!allProjects[id]) {
            allProjects[id] = new ProjectGroup();
                if (t!.project) {
            allProjects[id].project = t!.project
          } else {
            allProjects[id].project = { id: "-1", name: "Unknown" }
          }
              }
              allProjects[id].timeEntries.push(t);
              allProjects[id].elapsed += t!.elapsedSeconds!;
            });
            const count = allTimer ? allTimer.length : 0;
            if (allProjects && count > 0 && allTimer) {
            this.calculateBarchart(allProjects, timefilter, filterTimerStart, filterTimerEnd)
              console.log('All', allProjects);
              const id = allTimer!.find(t => t?.project !== null)!.project!.id!;
              console.log('ID:', id);
              barchart = {
            labels: allProjects[id]?.dates.map((d: Moment.Moment) => d?.format("D MMM ddd")),
                backgroundColors: _.map(allProjects, p => this.getRandomColor()),
                datasets: _.map(allProjects, (p, key) => ({
            label: p.project?.name,
                  data: p.ranges.map(s => s / 3600),
                  backgroundColor: this.getRandomColor(),

                }))
              }

              console.log("C", barchart)
            }
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
            <MuiPickersUtilsProvider utils={DateFnsUtils}>

              <Grid container spacing={3}
                className={classes.container}
              >
                <Grid item sm={3} xs={6}>
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
                </Grid>

                <Grid item sm={3} xs={6}>

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
                </Grid>
                <Grid item sm={2} xs={4}>
                  <FormControl
                    style={{ marginTop: -5 }}
                    className={[
                      classes.formControl,
                    ].join(" ")}>
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
                  </FormControl>
                </Grid>
                <Grid item sm={2} xs={4}>
                  <FormControl
                    style={{ marginTop: -5 }}
                    className={[
                      classes.formControl,
                    ].join(" ")}>

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
                  </FormControl>

                </Grid>

                {/* // Billed Unbilled */}

                <Grid item sm={1} xs={2}>
                  <FormControl
                    style={{ marginTop: -5 }}
                    className={[
                      classes.formControl,
                    ].join(" ")}>
                    <FormControlLabel
                      control={
                        <Switch
                          checked={this.state.filterIsBilled}
                          onChange={(event) => {
                            let filterIsBilled = event.target.checked;
                            console.log(filterIsBilled);
                            this.setState({
                              filterIsBilled: filterIsBilled,
                            });
                          }
                          }
                          name="checkedB"
                          color="primary"
                        />
                      }
                      label="Billed"
                    />

                  </FormControl>

                  {/* </Grid>
                    <Grid item sm={3} xs={6}> */}
                  <FormControl
                    style={{ marginTop: -5 }}
                    className={[
                      classes.formControl,
                    ].join(" ")}>
                    <FormControlLabel
                      control={
                        <Switch
                          checked={this.state.filterIsUnbilled}
                          onChange={(event) => {
                            let filterIsUnbilled = event.target.checked;
                            this.setState({
                              filterIsUnbilled: filterIsUnbilled,
                            });
                          }
                          }
                          name="checkedB"
                          color="primary"
                        />
                      }
                      label="Unbilled"
                    />

                  </FormControl>

                </Grid>

              </Grid>
            </MuiPickersUtilsProvider>
            <div
              className={classes.scroll}
            >
              <div
                style={{
                  overflowY: "auto",
                  height: "50%",
                }}>
                <BarChart data={barchart} options={this.options}></BarChart>

              </div>
              <div
                className={classes.list}
                style={{
                  overflowY: "auto",
                  height: "50%",
                }}
              >

                <List
                  component="nav"
                  aria-label="main mailbox folders"
                >
                  {_.map(allProjects, (entry: ProjectGroup, index) => (
                    <>
                      <ListItem button key={index} onClick={() => {
                        this.handleExpand(entry, isOpen);
                      }}>
                        <ListItemIcon>
                          <Assignment />
                        </ListItemIcon>
                        <ListItemText style={{ width: "70%" }} primary={entry!.project!.name! + " (" + entry!.timeEntries.length + ")"} ></ListItemText>
                        <ListItemText style={{ width: "80px" }} primary={TimerSrv.hms(entry!.elapsed)} ></ListItemText>
                        <ListItemSecondaryAction>
                          <Typography
                            // component="h1"
                            // variant="h2"
                            align="right"
                            color="textPrimary"
                          //gutterBottom
                          >

                            {this.getIsOpen(entry) ? <ExpandLess /> : <ExpandMore />}
                          </Typography>

                        </ListItemSecondaryAction>
                      </ListItem>
                      <Collapse in={this.getIsOpen(entry)} timeout="auto" unmountOnExit>
                        <List dense={true} component="div" disablePadding>
                          {entry.timeEntries.map(t =>
                            (<ListItem button className={classes.nested}>
                              <ListItemIcon>
                                <TimerIcon />
                              </ListItemIcon>
                              <ListItemText style={{ width: "40%" }} primary={t.description} />
                              <ListItemIcon>
                                {t.isBilled &&
                                  <MoneyIcon />
                                }

                              </ListItemIcon>
                              <ListItemText style={{ width: "15%" }} primary={this.toLocaleDate(t.timerStart)} />

                              <ListItemText primary={TimerSrv.hms(t.elapsedSeconds)} />
                              <ListItemSecondaryAction>
                              </ListItemSecondaryAction>


                            </ListItem>)
                          )
                          }

                        </List>
                      </Collapse>
                    </>
                  ))}

                </List>
              </div>
            </div>
          </>
            )
          }}
        </AllTimerComponent>


      </div >
    );
  }
  calculateBarchart(allProjects: { [id: string]: ProjectGroup; }, timefilter: string, filterTimerStart: Moment.Moment, filterTimerEnd: Moment.Moment) {
    if (!filterTimerStart && !filterTimerEnd) {
      return;
    }
    const r = moment.range(filterTimerStart, filterTimerEnd);
    const days = r.duration("days") + 2;
    console.log('Diff:', r.duration("days"))
    _.map(allProjects, p => {
      if (days < 15) {
        let start = r.start;
        p.dates = new Array<Moment.Moment>(days);
        p.ranges = new Array<number>(days);
        p.dates.fill(moment(), 0, days);
        p.ranges.fill(0, 0, days);

        for (let interval = 0; interval <= days; interval++) {
          const r = moment.rangeFromInterval("days", 1, start);

          p.dates[interval] = start;
          p.timeEntries.forEach(t => {
            if (t && t.elapsedSeconds && r.contains(t!.t1!)) {

              p.ranges[interval + 1] += t.elapsedSeconds!;
            }
          })

          start = r.end;
        }

      }
      if (days > 14) {
        const weeks = r.duration("weeks") + 1;
        let start = r.start;
        p.dates = new Array<Moment.Moment>(weeks);
        p.ranges = new Array<number>(weeks);
        p.dates.fill(moment(), 0, weeks);
        p.ranges.fill(0, 0, weeks);
        for (let interval = 0; interval <= weeks; interval++) {
          const r = moment.rangeFromInterval("weeks", 1, start);

          p.dates[interval] = start;
          p.timeEntries.forEach(t => {
            if (t && t.elapsedSeconds && r.contains(t!.t1!)) {

              p.ranges[interval] += t.elapsedSeconds!;
            }
          })

          start = r.end;
        }

      }
    })
  }
  getRandomColor() {
    var letters = '0123456789ABCDEF'.split('');
    var color = '#';
    for (var i = 0; i < 6; i++) {
      color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
  }
  private handleExpand(entry: ProjectGroup, isOpen: { [id: string]: boolean; }) {
    const id = entry!.project!.id!;
    const newOpen = _.clone(isOpen);
    newOpen[id] = !newOpen[id];
    this.setState({ isOpen: newOpen });
  }


  private getIsOpen(entry: ProjectGroup): boolean {
    const id = entry!.project!.id!;
    return this.state.isOpen[id];
  }

  setFilterTimerange(days: string) {
    let t1: Moment.Moment = moment();
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

export default withStyles(styles as any)(withRouter(Summary as any));
