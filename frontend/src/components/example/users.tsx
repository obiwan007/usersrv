// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
// Emotion styled component
import Button from '@material-ui/core/Button';
import React from "react";
import { withRouter } from "react-router-dom";
import { GetUsersComponent } from "../../graphql";
import Security from "../../lib/security";

// ----------------------------------------------------------------------------


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
                <button onClick={() => this.props.history.push("/login")}>
                  Login
                </button>
                <button onClick={() => this.refreshClick()}>Refresh</button>
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
              <ul>
                {data!.allUsers!.map(story => (
                  <li key={story!.id!}>
                    <a href={story!.email!} target="_blank">
                      {story!.name}
                    </a>
                    &nbsp;
                    <span>{story!.email}</span>
                  </li>
                ))}
              </ul>

              <Button color="primary" onClick={() => this.refreshClick()}>Refresh</Button>
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
