// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from "@apollo/react-common";
import { Button, Collapse, createStyles, List, ListItemIcon, ListItemSecondaryAction, ListItemText, Theme, Typography, WithStyles, withStyles } from "@material-ui/core";
import ListItem from "@material-ui/core/ListItem";
import { Assignment, ExpandLess, ExpandMore, Money as MoneyIcon, Timer as TimerIcon } from "@material-ui/icons";
import * as _ from "lodash";
import * as Moment from 'moment';
import { extendMoment } from 'moment-range';
import React from "react";
import { AllTimerComponent, DeleteTimerMutation, DeleteTimerMutationVariables, Project, refetchAllTimerQuery, Timer as TimerEntry } from "../../graphql";
import { Timer as TimerSrv } from "../../lib/timer";
import theme from '../../theme';
import BarChart from "./barchart";
import { FilterData } from './filter';

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
  addOpen: boolean;
  editOpen: boolean;
  filterProject?: Project;
  currentTimer: TimerEntry;
  isOpen: { [id: string]: boolean };
}
interface IProps {
  history?: any;
  filter: FilterData;
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
  /**
   *
   *
   */
  constructor(props: PROPS_WITH_STYLES, state: IState) {
    super(props, state);
    this.state = {
      sum: 0,
      addOpen: false,
      editOpen: false,
      currentTimer: {},
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
  }
  componentWillUnmount() {
    //clearInterval(this.interval!);
  }

  render() {
    const { isOpen, addOpen, editOpen } = this.state;
    const { timefilter, filterProject, filterTimerStart, filterTimerEnd, filterIsUnbilled, filterIsBilled } = this.props.filter;
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

  deleteTimer = (entity: TimerEntry, deleteTimer: MutationFunction<DeleteTimerMutation, DeleteTimerMutationVariables>) => {
    deleteTimer({
      refetchQueries: [
        refetchAllTimerQuery({ d: { dayrange: this.props.filter.timefilter } }),
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

export default withStyles(styles as any)(Summary as any);
