// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { createStyles, InputLabel, MenuItem, Select, Theme, WithStyles, withStyles } from "@material-ui/core";
import React from "react";
import { AllProjectsComponent, Project } from "../../graphql";

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


}
interface IProps {
  project: Project;
  onChanged: (newProject: Project | null) => any;

}
export type PROPS_WITH_STYLES = IProps & WithStyles<typeof styles>;
export class ProjectSelect extends React.PureComponent<PROPS_WITH_STYLES, IState> {
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
  }
  componentWillUnmount() {
    //clearInterval(this.interval!);
  }

  componentDidUpdate(nextProps: IProps) {
    // console.log('Didl Update', nextProps.timer, this.props.timer);
    // if (nextProps.timer !== this.props.timer) {
    //   console.log('Change detected', nextProps.timer)
    //   this.setState({ currentTimer: Object.assign({}, this.props.timer), addOpen: this.props.addOpen, editOpen: this.props.editOpen })
    // }
  }

  render() {
    return (
      <>
        {this.renderDialog()}
      </>
    );
  }
  renderDialog = () => {
    const { project, classes } = this.props;
    return (
      <AllProjectsComponent>
        {({ data, loading, error }) => {
          const allProjects = data;
          return (
            <>
              <InputLabel id="demo-simple-select-helper-label">Project</InputLabel>
              <Select
                // className={classes.selectEmpty}
                label="Project"
                value={
                  project && !loading && 
                    (allProjects?.allProjects as any[])
                      .length > 0
                    ? project?.id
                    : ""
                }
                onChange={(event) => {
                  console.log(
                    "Clientselection:",
                    event.target
                  );

                  const selected = allProjects?.allProjects?.find(c => c?.id === event.target.value as string);
                  this.props.onChanged(selected ? selected : null);
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
            </>
          )
        }}
      </AllProjectsComponent >
    );
  }

}

export default withStyles(styles as any)((ProjectSelect as any));
