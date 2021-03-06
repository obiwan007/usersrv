// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { Box, createStyles, Grid, IconButton, Tab, Tabs, Theme, WithStyles, withStyles } from "@material-ui/core";
import { GetApp as ExportIcon } from '@material-ui/icons';
import moment from "moment";
import React from "react";
import { withRouter } from "react-router-dom";
import { BehaviorSubject } from "rxjs";
import { Project, Timer as TimerEntry } from "../../graphql";
import Filter, { FilterData } from "./filter";
import ReportDetails from "./reportDetails";
import Summary from "./summary";

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
  selectedTab: number;
  filter: FilterData;
}
interface IProps {
  history?: any;
}

type TimerMoment = TimerEntry & {
  t1: moment.Moment,
  t2: moment.Moment,
}

interface TabPanelProps {
  children?: React.ReactNode;
  index: any;
  value: any;
}

function TabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box p={3}>
          <div>{children}</div>
        </Box>
      )}
    </div>
  );
}

export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class Reports extends React.PureComponent<PROPS_WITH_STYLES, IState> {

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

  exportClicked = new BehaviorSubject(false);
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
      selectedTab: 0,
      filter: new FilterData(),
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
    const { filter, selectedTab, filterProject, filterTimerStart, filterTimerEnd, timefilter, addOpen, editOpen } = this.state;
    const { classes } = this.props;
    let allTimer: TimerMoment[] | null | undefined = [];
    console.log("Filter", filter);
    return (
      <div>
        <div style={{ paddingLeft: 5, paddingRight: 5, width: "100%" }}>
          <Filter filter={filter} onUpdate={(filter: FilterData) => {
            console.log("Onpdate", JSON.stringify(filter, null, 2));
            this.setState({ filter: Object.assign({}, filter) })
          }}>
          </Filter>
        </div>
        <Grid container>
          <Grid item sm={11}>
            <Tabs value={selectedTab} onChange={this.handleChange} aria-label="simple tabs example">
              <Tab label="Summary" />
              <Tab label="Details" />
              <Tab label="Item Three" />
            </Tabs>
          </Grid>
          <Grid item sm={1}>
            {
              selectedTab === 1 &&
              <IconButton onClick={() => this.exportClicked.next(true)} color="primary" aria-label="add to shopping cart">
                <ExportIcon />
              </IconButton>
            }

          </Grid>
        </Grid>
        <TabPanel value={selectedTab} index={0}>
          <Summary filter={filter}></Summary>
        </TabPanel>
        <TabPanel value={selectedTab} index={1}>
          <ReportDetails export={this.exportClicked.asObservable()} filter={filter}></ReportDetails>
        </TabPanel>
        <TabPanel value={selectedTab} index={2}>
          Item Three
      </TabPanel>
      </div >
    );

  }



  handleChange = (event: React.ChangeEvent<{}>, selectedTab: number) => {
    // console.log('Selected', selectedTab)
    this.setState({ selectedTab })
  };

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
