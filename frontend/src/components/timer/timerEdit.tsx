// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { MutationFunction } from '@apollo/react-common';
import DateFnsUtils from '@date-io/date-fns';
import { Button, Checkbox, createStyles, Dialog, DialogActions, DialogContent, DialogTitle, FormControl, FormControlLabel, Grid, InputLabel, MenuItem, Select, TextField, Theme, WithStyles, withStyles } from "@material-ui/core";
import { KeyboardDatePicker, KeyboardTimePicker, MuiPickersUtilsProvider } from '@material-ui/pickers';
import * as _ from "lodash";
import React from "react";
import { AllProjectsComponent, CreateTimerComponent, CreateTimerMutation, CreateTimerMutationVariables, refetchAllTimerQuery, Timer as TimerEntry, TimerInput, UpdateTimerComponent, UpdateTimerMutation, UpdateTimerMutationVariables } from "../../graphql";
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
  description: string;
  currentTimer: TimerEntry;
  addOpen: boolean;
  editOpen: boolean;

}
interface IProps {
  history?: any;
  timefilter: string;
  onClose: () => any;
  timer: TimerEntry;
  addOpen: boolean;
  editOpen: boolean;

}
export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class TimerEdit extends React.PureComponent<PROPS_WITH_STYLES, IState> {
  /**
   *
   *
   */
  constructor(props: PROPS_WITH_STYLES, state: IState) {
    super(props, state);
    this.state = {
      currentTimer: {},
      description: "",
      addOpen: false,
      editOpen: false,
    };
  }

  componentDidMount() {
  }
  componentWillUnmount() {
    //clearInterval(this.interval!);
  }

  componentDidUpdate(nextProps: IProps) {
    if (nextProps.timer !== this.props.timer) {
      this.setState({ currentTimer: Object.assign({}, this.props.timer), addOpen: this.props.addOpen, editOpen: this.props.editOpen })
    }
  }

  render() {
    return (
      <>
        {this.renderDialog()}
      </>
    );
  }
  renderDialog = () => {
    const { addOpen, editOpen, currentTimer } = this.state;
    const { classes } = this.props;
    return (
      <AllProjectsComponent>
        {({ data, loading, error }) => {
          const allProjects = data;
          return (
            <UpdateTimerComponent>
              {(updateTimer, { data }) => {
                return (
                  <CreateTimerComponent>
                    {(createTimer, { data }) => {
                      return (

                        <Dialog
                          maxWidth="lg"
                          open={addOpen || editOpen}
                          onClose={() => this.closeDialog()}
                        >
                          <DialogTitle id="simple-dialog-title">
                            {editOpen ? "Edit Timer Entry" : "Add new Timer Entry"}
                          </DialogTitle>
                          <DialogContent>
                            {/* <DialogContentText>
                          To subscribe to this website, please enter
                          your email address here. We will send
                          updates occasionally.
                    </DialogContentText> */}
                            <form noValidate autoComplete="off">
                              {/* <FormControl className={classes.formControl}>
                                <TextField
                                  required
                                  autoFocus
                                  defaultValue={currentTimer.name}
                                  onChange={(data) => {
                                    currentTimer.name = data.target.value!;
                                    console.log(currentTimer);
                                    this.setState({
                                      currentTimer: currentTimer,
                                    });
                                  }}
                                  margin="dense"
                                  id="name"
                                  label="Title of Project"
                                  type="text"
                                  fullWidth
                                />
                              </FormControl> */}
                              <FormControl className={classes.formControl}>
                                <TextField
                                  margin="dense"
                                  id="description"
                                  label="Description of Workunit"
                                  type="text"
                                  defaultValue={currentTimer.description}
                                  onChange={(data) => {
                                    currentTimer.description = data.target.value!;
                                    console.log(currentTimer);
                                    this.setState({
                                      currentTimer: currentTimer,
                                    });
                                  }}
                                  fullWidth
                                />
                              </FormControl>
                              <FormControl className={classes.formControl}>
                                <InputLabel id="demo-simple-select-helper-label">Project</InputLabel>
                                <Select
                                  // className={classes.selectEmpty}
                                  label="Project"
                                  value={
                                    allProjects &&
                                      (allProjects.allProjects as any[])
                                        .length > 0
                                      ? currentTimer?.project?.id
                                      : ""
                                  }
                                  onChange={(event) => {
                                    console.log(
                                      "Clientselection:",
                                      event.target
                                    );
                                    if (currentTimer) {
                                      console.log('Selected', event.target.value as string);
                                      // </form>currentProject.client = event.target.value as string;
                                      const selected = allProjects?.allProjects?.find(c => c?.id === event.target.value as string);
                                      currentTimer.project = selected;
                                      this.setState({ currentTimer: currentTimer });
                                    }
                                  }}
                                >
                                  <MenuItem
                                    aria-label="None"
                                    value=""
                                  />
                                  {allProjects &&
                                    allProjects.allProjects?.map((e: any) => (
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


                              <FormControl className={classes.formControl}>
                                <MuiPickersUtilsProvider utils={DateFnsUtils}>
                                  <Grid container justify="space-around">
                                    <Grid item>
                                      <KeyboardDatePicker
                                        margin="normal"
                                        id="date-picker-dialog"
                                        label="Date Start"
                                        format="MM/dd/yyyy"
                                        value={currentTimer.timerStart}
                                        onChange={(date) => {
                                          currentTimer.timerStart = date?.toISOString()!;
                                          console.log(currentTimer);
                                          this.setState({
                                            currentTimer: currentTimer,
                                          });
                                        }
                                        }
                                        KeyboardButtonProps={{
                                          'aria-label': 'change date',
                                        }}
                                      />
                                    </Grid>
                                    <Grid item>
                                      <KeyboardTimePicker
                                        margin="normal"
                                        id="time-picker"
                                        label="Time Start"
                                        value={currentTimer.timerStart}
                                        onChange={(date) => {
                                          currentTimer.timerStart = date?.toISOString()!;
                                          console.log(currentTimer);
                                          this.setState({
                                            currentTimer: currentTimer,
                                          });
                                        }
                                        }
                                        KeyboardButtonProps={{
                                          'aria-label': 'change time',
                                        }}
                                      />
                                    </Grid>

                                    {!currentTimer.isRunning &&
                                      <>
                                        <Grid item>
                                          <KeyboardDatePicker
                                            margin="normal"
                                            id="date-picker-dialog"
                                            label="End"
                                            format="MM/dd/yyyy"
                                            value={currentTimer.timerEnd}
                                            onChange={(date) => {
                                              currentTimer.timerEnd = date?.toISOString()!;
                                              console.log(currentTimer);
                                              this.setState({
                                                currentTimer: currentTimer,
                                              });
                                            }
                                            }
                                            KeyboardButtonProps={{
                                              'aria-label': 'change date',
                                            }}
                                          />
                                        </Grid>
                                        <Grid item>
                                          <KeyboardTimePicker
                                            margin="normal"
                                            id="time-picker"
                                            label="Time End"
                                            value={currentTimer.timerEnd}
                                            onChange={(date) => {
                                              currentTimer.timerEnd = date?.toISOString()!;
                                              console.log(currentTimer);
                                              this.setState({
                                                currentTimer: currentTimer,
                                              });
                                            }
                                            }
                                            KeyboardButtonProps={{
                                              'aria-label': 'change time',
                                            }}
                                          />
                                        </Grid>
                                      </>
                                    }


                                  </Grid>
                                </MuiPickersUtilsProvider>
                              </FormControl>
                              <FormControl className={classes.formControl}>
                                <FormControlLabel
                                  control={<Checkbox checked={currentTimer!.isBilled!} onChange={(data) => {
                                    currentTimer.isBilled = data.target.checked;
                                    console.log(currentTimer);
                                    this.setState({
                                      currentTimer: _.cloneDeep(currentTimer),
                                    });
                                  }} name="checkedA" />}
                                  label="Is Billed"
                                />
                              </FormControl>
                            </form>
                          </DialogContent>
                          <DialogActions>
                            <Button
                              onClick={() => this.handleClose(false)}
                              color="primary"
                            >
                              Cancel
                    </Button>
                            <Button
                              onClick={() => this.handleClose(true, updateTimer, createTimer)}
                              color="primary"
                            >
                              {addOpen ? "Create" : "Update"}
                            </Button>
                          </DialogActions>

                        </Dialog>
                      )
                    }}
                  </CreateTimerComponent>)
              }}
            </UpdateTimerComponent>)
        }
        }
      </AllProjectsComponent >
    );
  }

  closeDialog = () => {
    this.setState({ addOpen: false, editOpen: false, currentTimer: {} })
    this.props.onClose();
  }


  handleClose = async (state: any,
    updateClient?: MutationFunction<UpdateTimerMutation, UpdateTimerMutationVariables>,
    createClient?: MutationFunction<CreateTimerMutation, CreateTimerMutationVariables>,
  ) => {
    console.log(state);
    // if (state) {
    //   project.addProject(this.state.currentProject);
    const newData = this.state.currentTimer!;
    const data: TimerInput = {
      id: newData.id,
      name: newData.name,
      description: newData.description,
      project: newData.project?.id,
      timerStart: newData.timerStart,
      isBilled: newData.isBilled,
    };
    if (!newData.isRunning) {
      data.timerEnd = newData.timerEnd;
    }
    console.log('Newdata', data)
    if (this.state.editOpen && updateClient) {
      await updateClient({
        refetchQueries: [
          refetchAllTimerQuery({ d: { dayrange: this.props.timefilter } }),
          refetchAllTimerQuery({ d: { dayrange: "0" } }),
        ],
        fetchPolicy: "no-cache",
        variables: {
          d: data
        },
      }).catch((ex) => {
        console.log(
          "Error in mutation",
          ex
        );
      });
    }
    else if (this.state.addOpen && createClient) {
      await createClient({
        refetchQueries: [
          refetchAllTimerQuery({ d: { dayrange: this.props.timefilter } }),
          refetchAllTimerQuery({ d: { dayrange: "0" } }),
        ],
        fetchPolicy: "no-cache",
        variables: {
          d: data
        },
      }).catch((ex) => {
        console.log(
          "Error in mutation",
          ex
        );
      });
    }
    this.closeDialog();
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

  handleDescriptionField = (event: any) => {
    this.state.currentTimer.description = event.target.value!;
    this.setState({ currentTimer: this.state.currentTimer });
    event.preventDefault();
  };


}

export default withStyles(styles as any)((TimerEdit as any));
