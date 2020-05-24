// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { createStyles, Theme, WithStyles, withStyles } from "@material-ui/core";
import * as _ from "lodash";
import React from "react";
import { Bar, defaults } from 'react-chartjs-2';

_.merge(defaults, {
  bar: {
    scales: {
      yAxes: [{
        gridLines: {
          color: 'grey' // makes grid lines from y axis red
        }
      }],
      xAxes: [{
        gridLines: {
          color: 'grey' // makes grid lines from y axis red
        }
      }]
    },
  },
  global: {
    defaultColor: 'white',
    defaultFontColor: "white",
    line: {
      borderColor: 'white',
    },

  }
});
console.log("CHART", defaults);
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
  data: any;
  options: any;
}
export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class BarChart extends React.PureComponent<IProps, IState> {
  interval?: any;
  data = {
    labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
    datasets: [
      {
        label: 'My First dataset',
        // backgroundColor: 'rgba(255,99,132,0.2)',
        // borderColor: 'rgba(255,99,132,1)',
        // borderWidth: 1,
        // hoverBackgroundColor: 'rgba(255,99,132,0.4)',
        // hoverBorderColor: 'rgba(255,99,132,1)',
        data: [65, 59, 80, 81, 56, 55, 40]
      },
      {
        label: 'My Second dataset',
        // backgroundColor: 'rgba(255,255,132,0.2)',
        // // borderColor: 'rgba(255,99,132,1)',
        // borderWidth: 1,
        // hoverBackgroundColor: 'rgba(255,99,132,1)',
        // hoverBorderColor: 'rgba(255,99,132,1)',
        data: [165, 39, 40, 41, 36, 85, 80]
      }
    ]
  };
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

  }
  componentDidUpdate(nextProps: IProps) {
    // console.log('Didl Update', nextProps.currentTimer, this.props.currentTimer );

  }
  componentWillUnmount() {

  }

  render() {
    const { data, options } = this.props;
    return (
      <div style={{ position: "relative", margin: "auto", width: "100%", height: '100%' }}>
        {/* <h2>Bar Example (custom size)</h2> */}
        <Bar
          data={data}
          // width={50}
          // height={150}
          options={{...options,
            maintainAspectRatio: false
          }}
        />
      </div>
    );
  }




  // recalcSummary( TimeEntry[]): number {
  //   let sum = 0;
  //   list.forEach((l) => (sum += l.elapsedSeconds));
  //   return sum;
  // }
}

export default withStyles(styles as any)(BarChart as any);
