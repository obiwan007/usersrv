// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import {
  createStyles,
  Theme,
  Typography,
  WithStyles,
  withStyles,
} from "@material-ui/core";
import React from "react";
import { withRouter } from "react-router-dom";
import { Timer as TimerSrv } from "../../lib/timer";

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
    topButtons: {
      marginLeft: spacing(2),
      paddingRight: spacing(1),
    },
    selectEmpty: {
      // marginTop: 13,
    },
  });

interface IState {
  elapsed: number;
}
interface IProps {
  history?: any;
  currentTimer: any;
}
export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class RunningClock extends React.PureComponent<IProps, IState> {
  interval?: any;

  /**
   *
   *
   */
  constructor(props: IProps, state: IState) {
    super(props, state);
    this.state = {
      elapsed: 0,
    };
  }

  componentDidMount() {
    this.componentWillUpdate();
  }
  componentWillReceiveProps(nextProps: IProps) {
    if (nextProps.currentTimer) {
      if (!this.interval) {
        this.interval = setInterval(() => {
          this.checkTimer();
        }, 500);
      }
    } else {
      if (this.interval) {
        clearInterval(this.interval!);
        this.interval = null;
      }
    }
  }
  componentWillUpdate() {}
  componentWillUnmount() {
    if (this.interval) {
      clearInterval(this.interval!);
      this.interval = null;
    }
  }

  render() {
    const { currentTimer } = this.props;

    return (
      <Typography variant="body2" color="textSecondary" align="center">
        {this.showElapsed(currentTimer)}
      </Typography>
    );
  }

  showElapsed(t: any): string {
    if (t) {
      const time1 = new Date(t.timerStart!);
      const time2 = new Date();
      return TimerSrv.hms((time2.getTime() - time1.getTime()) / 1000);
    }
    return "00:00:00";
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

export default withStyles(styles as any)(withRouter(RunningClock as any));
