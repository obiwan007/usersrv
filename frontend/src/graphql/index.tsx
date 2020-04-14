export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Mutation = {
  createUser?: Maybe<Scalars["String"]>;
};

export type MutationCreateUserArgs = {
  user: UserInput;
};

export type Query = {
  hello: Scalars["String"];
  User?: Maybe<User>;
  allUsers?: Maybe<Array<Maybe<User>>>;
};

export type QueryHelloArgs = {
  name: Scalars["String"];
};

export type QueryUserArgs = {
  id?: Maybe<Scalars["String"]>;
};

export type User = {
  id?: Maybe<Scalars["String"]>;
  name?: Maybe<Scalars["String"]>;
  email?: Maybe<Scalars["String"]>;
};

export type UserInput = {
  name?: Maybe<Scalars["String"]>;
  email?: Maybe<Scalars["String"]>;
  password?: Maybe<Scalars["String"]>;
};
export type GetUsersQueryVariables = {};

export type GetUsersQuery = { __typename?: "Query" } & {
  allUsers: Maybe<
    Array<Maybe<{ __typename?: "User" } & Pick<User, "name" | "email" | "id">>>
  >;
};

import gql from "graphql-tag";
import * as React from "react";
import * as ReactApollo from "react-apollo";
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;

export const GetUsersDocument = gql`
  query GetUsers {
    allUsers {
      name
      email
      id
    }
  }
`;

export const GetUsersComponent = (
  props: Omit<
    Omit<
      ReactApollo.QueryProps<GetUsersQuery, GetUsersQueryVariables>,
      "query"
    >,
    "variables"
  > & { variables?: GetUsersQueryVariables }
) => (
  <ReactApollo.Query<GetUsersQuery, GetUsersQueryVariables>
    query={GetUsersDocument}
    {...props}
  />
);

export type GetUsersProps<TChildProps = {}> = Partial<
  ReactApollo.DataProps<GetUsersQuery, GetUsersQueryVariables>
> &
  TChildProps;
export function withGetUsers<TProps, TChildProps = {}>(
  operationOptions?: ReactApollo.OperationOption<
    TProps,
    GetUsersQuery,
    GetUsersQueryVariables,
    GetUsersProps<TChildProps>
  >
) {
  return ReactApollo.withQuery<
    TProps,
    GetUsersQuery,
    GetUsersQueryVariables,
    GetUsersProps<TChildProps>
  >(GetUsersDocument, {
    alias: "withGetUsers",
    ...operationOptions
  });
}
