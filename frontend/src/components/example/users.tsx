// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import { Button, List, ListItem, ListItemIcon, ListItemText } from '@material-ui/core';
import { Person } from '@material-ui/icons';
import React from "react";
import { withRouter } from "react-router-dom";
import { GetUsersComponent } from "../../graphql";
import Security from "../../lib/security";
// ----------------------------------------------------------------------------

function ListItemLink(props: any) {
  return <ListItem button component="a" {...props} />;
}


interface IProps {
  history?: any;
}
// Say hello from GraphQL, along with a HackerNews feed fetched by GraphQL
export class Users extends React.PureComponent<IProps, any> {
  render() {
    return (
      <GetUsersComponent>
        {({ data, loading, error }) => {
          // Any errors? Say so!
          if (error) {
            return (
              <div>
                <h1>Error retrieving users list &mdash; {error.message}</h1>
                <Button variant="contained" color="primary" onClick={() => this.props.history.push("/login")}>
                  Login
                </Button>
                {/* <Button variant="contained" color="secondary" onClick={() => this.refreshClick()}>Refresh</Button> */}
              </div>
            );
          }

          // If the data is still loading, return with a basic
          // message to alert the user
          if (loading) {
            return <h1>Loading Users...</h1>;
          }

          // Otherwise, we have data to work with... map over it with a
          // bullet-point list
          return (
            <>
              <h3>Registered Users</h3>

              <List component="nav" aria-label="main mailbox folders">
                {data!.allUsers!.map(story => (
                  <ListItem button>
                    <ListItemIcon>
                      <Person />
                    </ListItemIcon>
                    <ListItemText primary={story!.name} secondary={story!.email} />
                  </ListItem>
                ))}
              </List>

              <Button variant="contained" color="primary" onClick={() => this.refreshClick()}>Refresh</Button>
            </>
          );
        }}
      </GetUsersComponent>
    );
  }

  async refreshClick() {
    try {
      const ok = await Security.refresh();
    } catch (ex) {
      console.log("Error in refreshing", ex);
    }
  }
}

export default withRouter(Users as any);
