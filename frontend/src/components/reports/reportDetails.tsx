// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from "@apollo/react-common";
import { Button, createStyles, Hidden, IconButton, ListItemIcon, ListItemSecondaryAction, ListItemText, Theme, WithStyles, withStyles } from "@material-ui/core";
import ListItem from "@material-ui/core/ListItem";
import { Delete, Timer as TimerIcon } from "@material-ui/icons";
import moment from "moment";
import React from "react";
import Autosizer from "react-virtualized-auto-sizer";
import { FixedSizeList, ListChildComponentProps } from 'react-window';
import { Observable, Subscription } from 'rxjs';
import { AllTimerComponent, DeleteTimerComponent, DeleteTimerMutation, DeleteTimerMutationVariables, refetchAllTimerQuery, Timer as TimerEntry } from "../../graphql";
import { Timer as TimerSrv } from "../../lib/timer";
import TimerEdit from "../timer/timerEdit";
import { FilterData } from './filter';
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
  addOpen: boolean;
  editOpen: boolean;
  currentTimer: TimerEntry;

}
interface IProps {
  history?: any;
  filter: FilterData;
  export: Observable<boolean>;
}

type TimerMoment = TimerEntry & {
  t1: moment.Moment,
  t2: moment.Moment,
}

export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class ReportDetails extends React.PureComponent<PROPS_WITH_STYLES, IState> {

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
  subExport?: Subscription;



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
    };
  }

  componentDidMount() {
    // this.setState({
    //   filterProject: {},
    // });
    // this.interval = setInterval(() => {
    //   this.checkTimer();
    // }, 500);

    this.subExport = this.props.export.subscribe((pressed) => {
      console.log("Export clicked", pressed);
    });

  }
  componentWillUnmount() {
    //clearInterval(this.interval!);
    this.subExport!.unsubscribe();
  }

  render() {
    const { addOpen, editOpen } = this.state;
    const { classes } = this.props;

    const { timefilter, filterProject, filterTimerStart, filterTimerEnd, filterIsUnbilled, filterIsBilled } = this.props.filter;

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

                  if (filterProject?.id && !loading && allTimer) {
                    allTimer = allTimer.filter((a: TimerMoment) => a?.project?.id === filterProject.id);
                  }
                  if (filterTimerStart && filterTimerEnd) {
                    allTimer = allTimer?.filter((a: TimerMoment) => a.t1.isBefore(filterTimerEnd) && a.t1.isAfter(filterTimerStart));
                  }
                  if (filterIsBilled === false || filterIsUnbilled === false) {
                    allTimer = allTimer?.filter((a: TimerMoment) =>
                      (filterIsBilled && (a.isBilled === true))
                      || (filterIsUnbilled && (a.isBilled === false))

                    );
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
                      <div
                        className={classes.list}
                        style={{
                          height: "calc(100vh - 260px)",
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

export default withStyles(styles as any)(ReportDetails as any);
