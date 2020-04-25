import gql from 'graphql-tag';
import * as React from 'react';
import * as ApolloReactCommon from '@apollo/react-common';
import * as ApolloReactComponents from '@apollo/react-components';
import * as ApolloReactHoc from '@apollo/react-hoc';
export type Maybe<T> = T | null;
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Mutation = {
   __typename?: 'Mutation';
  createUser?: Maybe<Scalars['String']>;
};


export type MutationCreateUserArgs = {
  user: UserInput;
};

export type Query = {
   __typename?: 'Query';
  hello: Scalars['String'];
  User?: Maybe<User>;
  allUsers?: Maybe<Array<Maybe<User>>>;
};


export type QueryHelloArgs = {
  name: Scalars['String'];
};


export type QueryUserArgs = {
  id?: Maybe<Scalars['String']>;
};

export type User = {
   __typename?: 'User';
  id?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  email?: Maybe<Scalars['String']>;
};

export type UserInput = {
  name?: Maybe<Scalars['String']>;
  email?: Maybe<Scalars['String']>;
  password?: Maybe<Scalars['String']>;
};

export type GetUsersQueryVariables = {};


export type GetUsersQuery = (
  { __typename?: 'Query' }
  & { allUsers?: Maybe<Array<Maybe<(
    { __typename?: 'User' }
    & Pick<User, 'name' | 'email' | 'id'>
  )>>> }
);


export const GetUsersDocument = gql`
    query GetUsers {
  allUsers {
    name
    email
    id
  }
}
    `;
export type GetUsersComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<GetUsersQuery, GetUsersQueryVariables>, 'query'>;

    export const GetUsersComponent = (props: GetUsersComponentProps) => (
      <ApolloReactComponents.Query<GetUsersQuery, GetUsersQueryVariables> query={GetUsersDocument} {...props} />
    );
    
export type GetUsersProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<GetUsersQuery, GetUsersQueryVariables>
    } & TChildProps;
export function withGetUsers<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  GetUsersQuery,
  GetUsersQueryVariables,
  GetUsersProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, GetUsersQuery, GetUsersQueryVariables, GetUsersProps<TChildProps, TDataName>>(GetUsersDocument, {
      alias: 'getUsers',
      ...operationOptions
    });
};
export type GetUsersQueryResult = ApolloReactCommon.QueryResult<GetUsersQuery, GetUsersQueryVariables>;