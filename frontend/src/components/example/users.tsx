// ReactQL Hacker News GraphQL example

// ----------------------------------------------------------------------------
// IMPORTS

/* NPM */
/* Local */
// Query to get top stories from HackerNews
import { GetUsersComponent } from "@/graphql";
// Emotion styled component
import styled from "@emotion/styled";
import React from "react";
import { withRouter } from "react-router-dom";

// ----------------------------------------------------------------------------

// Unstyled Emotion parent block, to avoid repeating <style> tags
// on child elements -- see https://github.com/emotion-js/emotion/issues/1061
const List = styled.ul``;

// Style the list item so it overrides the default font
const Story = styled("li")`
  font-size: 16px;

  a:hover {
    /* shows an example of how we can use themes */
    color: orange;
  }
`;

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
              <List>
                {data!.allUsers!.map(story => (
                  <Story key={story!.id!}>
                    <a href={story!.email!} target="_blank">
                      {story!.name}
                    </a>
                    &nbsp;
                    <span>{story!.email}</span>
                  </Story>
                ))}
              </List>
            </>
          );
        }}
      </GetUsersComponent>
    );
  }
}

export default withRouter(Users as any);
