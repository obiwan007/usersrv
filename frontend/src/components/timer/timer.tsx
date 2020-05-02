// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { IconButton, List, Typography } from "@material-ui/core";
import { PlayArrow, Stop } from "@material-ui/icons";
import MaterialTable from "material-table";
import React from "react";
import { withRouter } from "react-router-dom";
import client from "../../lib/client";
import project from "../../lib/project";
import timer, { TimeEntry } from "../../lib/timer";
// ----------------------------------------------------------------------------

interface IState {
  startTime: Date;
  endTime: Date;
  isRunning: boolean;
  elapsed: number;
  list: TimeEntry[];
  columns: any[];
}
interface IProps {
  history?: any;
}

export class Timer extends React.PureComponent<IProps, IState> {
  interval?: NodeJS.Timeout;

  /**
   *
   *
   */
  constructor(props: IProps, state: IState) {
    super(props, state);
    this.state = {
      isRunning: timer.getTimer().isRunning,
      startTime: timer.getTimer().timerStart,
      endTime: timer.getTimer().timerEnd,
      elapsed: timer.getTimer().elapsed(),
      list: timer.Entries(),
      columns: [],
    };
  }

  componentDidMount() {
    this.setState({ columns: this.getColumns() });
    this.interval = setInterval(() => {
      this.checkTimer();
    }, 500);
  }
  componentWillUnmount() {
    clearInterval(this.interval!);
  }
  render() {
    const { isRunning, list, columns } = this.state;
    const seconds = timer.elapsed();
    return (
      <div>
        <h3>Timer</h3>
        <IconButton
          onClick={() => this.startStopTimer()}
          edge="start"
          color="inherit"
          aria-label="menu"
        >
          {!isRunning && <PlayArrow />}
          {isRunning && <Stop />}
        </IconButton>
        <Typography variant="body2" color="textSecondary" align="center">
          {seconds}
        </Typography>

        <h3>Completed workunits</h3>
        {list.length}

        <List component="nav" aria-label="main mailbox folders">
          <MaterialTable
            options={{
              minBodyHeight: "calc(100vh - 360px)",
              maxBodyHeight: "calc(100vh - 360px)",
            }}
            title="Timetable"
            columns={columns}
            data={list?.map((u) => u) as any[]}
            editable={{
              isEditable: (rowData) => {
                console.log("Rorwdata", rowData);
                return true;
              }, // only name(a) rows would be editable
              isDeletable: (rowData) => true, // only name(a) rows would be deletable
              onRowUpdate: (newData, oldData) =>
                new Promise((resolve, reject) => {
                  setTimeout(() => {
                    {
                      const d = timer.update(newData);
                      console.log("New List", d);
                      this.setState({ list: d }, () => resolve());
                    }
                    resolve();
                  }, 1000);
                }),
              onRowDelete: (oldData) =>
                new Promise((resolve, reject) => {
                  setTimeout(() => {
                    {
                      const d = timer.del(oldData);
                      console.log("New List", d);
                      this.setState({ list: d }, () => resolve());
                    }
                    resolve();
                  }, 1000);
                }),
            }}
          />
          {/* {data!.allUsers!.map(data => (
    <ListItem button>
      <ListItemIcon>
        <Person />
      </ListItemIcon>
      <ListItemText primary={data!.name} secondary={data!.email} />
    </ListItem>
  ))} */}
        </List>
      </div>
    );
  }

  getColumns = () => {
    const columns = [
      { title: "Title", field: "title" },
      { title: "Description", field: "description" },
      {
        title: "Client",
        field: "client",
        editable: () => true,
        lookup: client.EntriesDict(),
      },
      {
        title: "Project",
        field: "project",
        editable: () => true,
        lookup: project.EntriesDict(),
      },
      { title: "Start", field: "tStart" },
      { title: "End", field: "tEnd" },
      { title: "Seconds", field: "elapsedSeconds" },
    ];
    return columns;
  };
  startStopTimer() {
    this.state.isRunning ? timer.endTimer() : timer.startTimer();
    this.setState({
      list: timer.Entries(),
    });
  }

  checkTimer = () => {
    this.setState({
      isRunning: timer.getTimer().isRunning,
      startTime: timer.getTimer().timerStart,
      endTime: timer.getTimer().timerEnd,
      elapsed: timer.getTimer().elapsed(),
    });
  };
}

export default withRouter(Timer as any);
