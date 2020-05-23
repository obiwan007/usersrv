// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from "@apollo/react-common";
import DateFnsUtils from '@date-io/date-fns';
import { Box, Button, Collapse, createStyles, FormControl, InputLabel, List, ListItemIcon, ListItemSecondaryAction, ListItemText, MenuItem, Select, Theme, Typography, WithStyles, withStyles } from "@material-ui/core";
import ListItem from "@material-ui/core/ListItem";
import { Assignment, ExpandLess, ExpandMore, StarBorder } from "@material-ui/icons";
import { KeyboardDatePicker, MuiPickersUtilsProvider } from '@material-ui/pickers';
import * as _ from "lodash";
import moment from "moment";
import React from "react";
import { withRouter } from "react-router-dom";
import { AllTimerComponent, DeleteTimerMutation, DeleteTimerMutationVariables, Project, refetchAllTimerQuery, Timer as TimerEntry } from "../../graphql";
import { Timer as TimerSrv } from "../../lib/timer";
import theme from '../../theme';
import ProjectSelect from "../projects/projectSelect";

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
    nested: {
      paddingLeft: theme.spacing(4),
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
  isOpen: { [id: string]: boolean };
}
interface IProps {
  history?: any;
}

type TimerMoment = TimerEntry & {
  t1: moment.Moment,
  t2: moment.Moment,
}

class ProjectGroup {
  project?: Project;
  timeEntries: TimerEntry[] = [];
  elapsed: number = 0;
}

export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class Summary extends React.PureComponent<PROPS_WITH_STYLES, IState> {

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
    const { isOpen, filterProject, filterTimerStart, filterTimerEnd, timefilter, addOpen, editOpen } = this.state;
    const { classes } = this.props;
    let allTimer: TimerMoment[] | null | undefined = [];

    let allProjects: { [id: string]: ProjectGroup } = {}
    let open = true;
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
            }
            );

            console.log('All', allProjects);

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
                          <ListItemText style={{width: "80%"}} primary={entry!.project!.name! + " (" + entry!.timeEntries.length + ")"} ></ListItemText>
                          <ListItemText primary={TimerSrv.hms(entry!.elapsed)} ></ListItemText>                          
                          <ListItemSecondaryAction>
                            <Typography
                              // component="h1"
                              // variant="h2"
                              align="right"
                              color="textPrimary"
                              gutterBottom
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
                                  <StarBorder />
                                </ListItemIcon>
                                <ListItemText style={{width: "80%"}}  primary={t.description} />
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
              </>
            )
          }}
        </AllTimerComponent>


      </div >
    );
  }

  private handleExpand(entry: ProjectGroup, isOpen: { [id: string]: boolean; }) {
    const id = entry!.project!.id!;
    isOpen[id] = !isOpen[id];
    this.setState({ isOpen });
  }


  private getIsOpen(entry: ProjectGroup): boolean {
    const id = entry!.project!.id!;
    return this.state.isOpen[id];
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

export default withStyles(styles as any)(withRouter(Summary as any));
