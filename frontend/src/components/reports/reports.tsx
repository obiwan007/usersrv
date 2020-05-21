// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from "@apollo/react-common";
import { Box, Button, createStyles, FormControl, Hidden, IconButton, InputLabel, ListItemIcon, ListItemSecondaryAction, ListItemText, MenuItem, Select, Theme, WithStyles, withStyles } from "@material-ui/core";
import ListItem from "@material-ui/core/ListItem";
import { Delete, Timer as TimerIcon } from "@material-ui/icons";
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
      paddingRight: spacing(1),
      paddingLeft: spacing(1),
      minWidth: 120,
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
  });

interface IState {
  sum: number;
  timefilter: string;
  addOpen: boolean;
  editOpen: boolean;
  filterProject: Project;
  currentTimer: TimerEntry;
}
interface IProps {
  history?: any;
}
export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class Reports extends React.PureComponent<PROPS_WITH_STYLES, IState> {

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
      timefilter: "1",
      sum: 0,
      addOpen: false,
      editOpen: false,
      filterProject: {},
      currentTimer: {}
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
    const { filterProject, timefilter, addOpen, editOpen } = this.state;
    const { classes } = this.props;
    let allTimer: any = [];

    console.log('FilterProject:', filterProject);
    return (
      <div>

        <DeleteTimerComponent>
          {(deleteTimer, { data }) => {
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
                  allTimer = data;
                  const count = allTimer?.allTimer?.length;

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
                        </Box>


                        <Box>
                          <FormControl
                            style={{ width: 250 }}
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

  // recalcSummary( TimeEntry[]): number {
  //   let sum = 0;
  //   list.forEach((l) => (sum += l.elapsedSeconds));
  //   return sum;
  // }
}

export default withStyles(styles as any)(withRouter(Reports as any));
