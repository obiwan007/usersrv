// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import DateFnsUtils from '@date-io/date-fns';
import { createStyles, FormControl, FormControlLabel, Grid, InputLabel, MenuItem, Select, Switch, Theme, WithStyles, withStyles } from "@material-ui/core";
import { KeyboardDatePicker, MuiPickersUtilsProvider } from '@material-ui/pickers';
import * as Moment from 'moment';
import { extendMoment } from 'moment-range';
import React from "react";
import { Project } from "../../graphql";
import { Timer as TimerSrv } from "../../lib/timer";
import theme from '../../theme';
import ProjectSelect from "../projects/projectSelect";

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

export class FilterData {
  filterTimerEnd: any;
  filterTimerStart: any;
  filterIsBilled: boolean = false;
  filterIsUnbilled: boolean = true;
  filterProject?: Project;
  timefilter: string = '';

}

interface IState {
}
interface IProps {
  history?: any;
  filter: FilterData;
  onUpdate: (filter: FilterData) => any;
}

export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class Filter extends React.PureComponent<PROPS_WITH_STYLES, IState> {


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
    };
  }



  componentDidMount() {
    // this.setState({
    //   filterProject: {},
    // });
    // this.interval = setInterval(() => {
    //   this.checkTimer();
    // }, 500);
    this.setFilterTimerange(this.props.filter, "7")
  }
  componentWillUnmount() {
    //clearInterval(this.interval!);
  }

  render() {

    const { onUpdate, classes } = this.props;

    const filter = this.props.filter;

    const { timefilter, filterProject, filterTimerStart, filterTimerEnd, filterIsUnbilled, filterIsBilled } = filter;


    return (
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
                  this.setFilterTimerange(filter, event.target.value! as string)

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
                onChanged={(p: Project) => {
                  filter.filterProject = p;
                  onUpdate(filter);
                }}
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

                  filter.filterTimerStart = filterTimerStart;
                  onUpdate(filter);

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

                  filter.filterTimerEnd = filterTimerEnd;
                  onUpdate(filter);

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
                    checked={filterIsBilled}
                    onChange={(event) => {
                      filter.filterIsBilled = event.target.checked;
                      onUpdate(filter);
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
                    checked={filterIsUnbilled}
                    onChange={(event) => {
                      filter.filterIsUnbilled = event.target.checked;
                      onUpdate(filter);
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
    );
  }


  setFilterTimerange(filter: FilterData, days: string) {
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

    filter.timefilter = days;
    filter.filterTimerStart = t2;
    filter.filterTimerEnd = t1;
    console.log("Filter", JSON.stringify(filter,null,2));
    this.props.onUpdate(filter)
  }

}

export default withStyles(styles as any)(Filter as any);
