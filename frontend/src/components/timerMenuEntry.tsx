// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { createStyles, List, ListItem, ListItemIcon, ListItemText, Theme, WithStyles, withStyles } from "@material-ui/core";
import { Timer as TimerIcon } from "@material-ui/icons";
import React from "react";
import { AllTimerComponent, Timer as TimerEntry } from '../graphql';
import { Timer as TimerSrv } from "../lib/timer";

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
    onClick: () => any;
    isActiveRoute: boolean;
}
export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class TimerMenuEntry extends React.PureComponent<IProps, IState> {
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
        if (!this.interval) {
            this.interval = setInterval(() => {
                this.checkTimer();
            }, 500);
        }
    }
    // componentWillReceiveProps(nextProps: IProps) {
    //   if (nextProps.currentTimer) {
    //     if (!this.interval) {
    //       this.interval = setInterval(() => {
    //         this.checkTimer();
    //       }, 500);
    //     }
    //   } else {
    //     if (this.interval) {
    //       clearInterval(this.interval!);
    //       this.interval = null;
    //     }
    //   }
    // }

    componentWillUnmount() {
        if (this.interval) {
            clearInterval(this.interval!);
            this.interval = null;
        }
    }

    render() {
        const { isActiveRoute } = this.props;
        let currentTimer: TimerEntry | undefined = undefined;
        return (
            <List>
                <AllTimerComponent variables={{ d: { dayrange: "0" } }}>
                    {({ data, loading, error }) => {
                        // Any errors? Say so!
                        currentTimer = undefined;
                        data?.allTimer?.forEach((d) => {
                            if (d) {
                                (d as any).projectId = d?.project?.id;
                                if (d.isRunning === true) {
                                    currentTimer = d;
                                }
                            }
                        });
                        const timerText =
                            currentTimer && currentTimer!.isRunning && !isActiveRoute
                                ? this.showElapsed(currentTimer)
                                : "Timer";
                        // console.log("currentTimer", currentTimer, data?.allTimer);
                        if (error) {
                            return (
                                <div>
                                    <h1>Error retrieving Timer list &mdash; {error.message}</h1>
                                    {/* <Button variant="contained" color="secondary" onClick={() => this.refreshClick()}>Refresh</Button> */}
                                </div>
                            );
                        }

                        // If the data is still loading, return with a basic
                        // message to alert the user

                        return (
                            <>
                                {[{ txt: timerText, link: "/timer" }].map((o, index) => (
                                    <ListItem
                                        selected={isActiveRoute}
                                        onClick={() => this.props.onClick()}
                                        button
                                        key={o.txt}
                                    >
                                        <ListItemIcon>
                                            <TimerIcon />
                                        </ListItemIcon>
                                        <ListItemText primary={o.txt} />
                                    </ListItem>
                                ))}
                            </>
                        );
                    }}
                </AllTimerComponent>
            </List>
        );
    }

    showElapsed(t: any): string {
        if (t) {
            const time1 = new Date(t.timerStart!);
            const time2 = new Date();
            return TimerSrv.hms((time2.getTime() - time1.getTime()) / 1000);
        }
        return "Timer";
    }
    checkTimer = () => {
        this.setState({
            elapsed: new Date().getTime(),
        });
    };


}

export default withStyles(styles as any)(TimerMenuEntry as any );
