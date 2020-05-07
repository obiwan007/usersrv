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

export type AllTimerFilter = {
  dayrange?: Maybe<Scalars['String']>;
};

export type Client = {
   __typename?: 'Client';
  id?: Maybe<Scalars['String']>;
  description?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  address?: Maybe<Scalars['String']>;
};

export type ClientInput = {
  id?: Maybe<Scalars['String']>;
  description?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  address?: Maybe<Scalars['String']>;
};

export type CreateTimerInput = {
  description?: Maybe<Scalars['String']>;
  project?: Maybe<Scalars['String']>;
};

export type Mutation = {
   __typename?: 'Mutation';
  createUser?: Maybe<Scalars['String']>;
  createTimer?: Maybe<Timer>;
  startTimer?: Maybe<Timer>;
  stopTimer?: Maybe<Timer>;
  deleteTimer?: Maybe<Timer>;
  updateTimer?: Maybe<Timer>;
  createProject?: Maybe<Project>;
  updateProject?: Maybe<Project>;
  deleteProject?: Maybe<Project>;
  createClient?: Maybe<Client>;
  updateClient?: Maybe<Client>;
  deleteClient?: Maybe<Client>;
};


export type MutationCreateUserArgs = {
  user: UserInput;
};


export type MutationCreateTimerArgs = {
  t: CreateTimerInput;
};


export type MutationStartTimerArgs = {
  timerId: Scalars['String'];
};


export type MutationStopTimerArgs = {
  timerId: Scalars['String'];
};


export type MutationDeleteTimerArgs = {
  timerId: Scalars['String'];
};


export type MutationUpdateTimerArgs = {
  t: TimerInput;
};


export type MutationCreateProjectArgs = {
  p: ProjectInput;
};


export type MutationUpdateProjectArgs = {
  p: ProjectInput;
};


export type MutationDeleteProjectArgs = {
  projectId: Scalars['String'];
};


export type MutationCreateClientArgs = {
  c: ClientInput;
};


export type MutationUpdateClientArgs = {
  c: ClientInput;
};


export type MutationDeleteClientArgs = {
  clientId: Scalars['String'];
};

export type Project = {
   __typename?: 'Project';
  id?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  description?: Maybe<Scalars['String']>;
  team?: Maybe<Scalars['String']>;
  client?: Maybe<Client>;
  status?: Maybe<Scalars['String']>;
};

export type ProjectInput = {
  id?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  description?: Maybe<Scalars['String']>;
  team?: Maybe<Scalars['String']>;
  client?: Maybe<Scalars['String']>;
  status?: Maybe<Scalars['String']>;
};

export type Query = {
   __typename?: 'Query';
  hello: Scalars['String'];
  User?: Maybe<User>;
  allUsers?: Maybe<Array<Maybe<User>>>;
  allTimer?: Maybe<Array<Maybe<Timer>>>;
  runningTimer?: Maybe<Timer>;
  getTimer?: Maybe<Timer>;
  getProject?: Maybe<Project>;
  allProjects?: Maybe<Array<Maybe<Project>>>;
  getClient?: Maybe<Client>;
  allClients?: Maybe<Array<Maybe<Client>>>;
};


export type QueryHelloArgs = {
  name: Scalars['String'];
};


export type QueryUserArgs = {
  id?: Maybe<Scalars['String']>;
};


export type QueryAllTimerArgs = {
  filter: AllTimerFilter;
};


export type QueryGetTimerArgs = {
  id?: Maybe<Scalars['String']>;
};


export type QueryGetProjectArgs = {
  id?: Maybe<Scalars['String']>;
};


export type QueryGetClientArgs = {
  id?: Maybe<Scalars['String']>;
};

export type Timer = {
   __typename?: 'Timer';
  id?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  description?: Maybe<Scalars['String']>;
  teammember?: Maybe<Scalars['String']>;
  client?: Maybe<Client>;
  project?: Maybe<Project>;
  tags?: Maybe<Scalars['String']>;
  elapsedSeconds?: Maybe<Scalars['Int']>;
  timerStart?: Maybe<Scalars['String']>;
  timerEnd?: Maybe<Scalars['String']>;
  isRunning?: Maybe<Scalars['Boolean']>;
  isBilled?: Maybe<Scalars['Boolean']>;
};

export type TimerInput = {
  id?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  description?: Maybe<Scalars['String']>;
  teammember?: Maybe<Scalars['String']>;
  client?: Maybe<Scalars['String']>;
  project?: Maybe<Scalars['String']>;
  tags?: Maybe<Scalars['String']>;
  elapsedSeconds?: Maybe<Scalars['Int']>;
  timerStart?: Maybe<Scalars['String']>;
  timerEnd?: Maybe<Scalars['String']>;
  isRunning?: Maybe<Scalars['Boolean']>;
  isBilled?: Maybe<Scalars['Boolean']>;
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

export type AllClientsQueryVariables = {};


export type AllClientsQuery = (
  { __typename?: 'Query' }
  & { allClients?: Maybe<Array<Maybe<(
    { __typename?: 'Client' }
    & Pick<Client, 'id' | 'name' | 'description' | 'address'>
  )>>> }
);

export type CreateClientMutationVariables = {
  d: ClientInput;
};


export type CreateClientMutation = (
  { __typename?: 'Mutation' }
  & { createClient?: Maybe<(
    { __typename?: 'Client' }
    & Pick<Client, 'id' | 'name' | 'description' | 'address'>
  )> }
);

export type UpdateClientMutationVariables = {
  d: ClientInput;
};


export type UpdateClientMutation = (
  { __typename?: 'Mutation' }
  & { updateClient?: Maybe<(
    { __typename?: 'Client' }
    & Pick<Client, 'id' | 'name' | 'description' | 'address'>
  )> }
);

export type DeleteClientMutationVariables = {
  d: Scalars['String'];
};


export type DeleteClientMutation = (
  { __typename?: 'Mutation' }
  & { deleteClient?: Maybe<(
    { __typename?: 'Client' }
    & Pick<Client, 'id'>
  )> }
);

export type GetUsersQueryVariables = {};


export type GetUsersQuery = (
  { __typename?: 'Query' }
  & { allUsers?: Maybe<Array<Maybe<(
    { __typename?: 'User' }
    & Pick<User, 'name' | 'email' | 'id'>
  )>>> }
);

export type AllProjectsQueryVariables = {};


export type AllProjectsQuery = (
  { __typename?: 'Query' }
  & { allProjects?: Maybe<Array<Maybe<(
    { __typename?: 'Project' }
    & Pick<Project, 'name' | 'description' | 'id' | 'team' | 'status'>
    & { client?: Maybe<(
      { __typename?: 'Client' }
      & Pick<Client, 'id' | 'name' | 'description'>
    )> }
  )>>> }
);

export type CreateProjectMutationVariables = {
  d: ProjectInput;
};


export type CreateProjectMutation = (
  { __typename?: 'Mutation' }
  & { createProject?: Maybe<(
    { __typename?: 'Project' }
    & Pick<Project, 'name' | 'description' | 'id' | 'team' | 'status'>
    & { client?: Maybe<(
      { __typename?: 'Client' }
      & Pick<Client, 'id' | 'name' | 'description'>
    )> }
  )> }
);

export type UpdateProjectMutationVariables = {
  d: ProjectInput;
};


export type UpdateProjectMutation = (
  { __typename?: 'Mutation' }
  & { updateProject?: Maybe<(
    { __typename?: 'Project' }
    & Pick<Project, 'name' | 'description' | 'id' | 'team' | 'status'>
    & { client?: Maybe<(
      { __typename?: 'Client' }
      & Pick<Client, 'id' | 'name' | 'description'>
    )> }
  )> }
);

export type DeleteProjectMutationVariables = {
  d: Scalars['String'];
};


export type DeleteProjectMutation = (
  { __typename?: 'Mutation' }
  & { deleteProject?: Maybe<(
    { __typename?: 'Project' }
    & Pick<Project, 'id'>
  )> }
);

export type AllTimerQueryVariables = {};


export type AllTimerQuery = (
  { __typename?: 'Query' }
  & { allTimer?: Maybe<Array<Maybe<(
    { __typename?: 'Timer' }
    & Pick<Timer, 'description' | 'id' | 'timerStart' | 'timerEnd' | 'isRunning' | 'isBilled' | 'tags' | 'elapsedSeconds'>
    & { project?: Maybe<(
      { __typename?: 'Project' }
      & Pick<Project, 'id' | 'name'>
    )> }
  )>>> }
);

export type CreateTimerMutationVariables = {
  d: CreateTimerInput;
};


export type CreateTimerMutation = (
  { __typename?: 'Mutation' }
  & { createTimer?: Maybe<(
    { __typename?: 'Timer' }
    & Pick<Timer, 'description' | 'id' | 'timerStart' | 'timerEnd' | 'isRunning' | 'isBilled' | 'tags' | 'elapsedSeconds'>
    & { project?: Maybe<(
      { __typename?: 'Project' }
      & Pick<Project, 'id' | 'name'>
    )> }
  )> }
);

export type UpdateTimerMutationVariables = {
  d: TimerInput;
};


export type UpdateTimerMutation = (
  { __typename?: 'Mutation' }
  & { updateTimer?: Maybe<(
    { __typename?: 'Timer' }
    & Pick<Timer, 'description' | 'id' | 'timerStart' | 'timerEnd' | 'isRunning' | 'isBilled' | 'tags' | 'elapsedSeconds'>
    & { project?: Maybe<(
      { __typename?: 'Project' }
      & Pick<Project, 'id' | 'name'>
    )> }
  )> }
);

export type DeleteTimerMutationVariables = {
  timerId: Scalars['String'];
};


export type DeleteTimerMutation = (
  { __typename?: 'Mutation' }
  & { deleteTimer?: Maybe<(
    { __typename?: 'Timer' }
    & Pick<Timer, 'id'>
  )> }
);

export type StartTimerMutationVariables = {
  timerId: Scalars['String'];
};


export type StartTimerMutation = (
  { __typename?: 'Mutation' }
  & { startTimer?: Maybe<(
    { __typename?: 'Timer' }
    & Pick<Timer, 'id' | 'description' | 'timerStart' | 'timerEnd' | 'isRunning' | 'isBilled' | 'tags' | 'elapsedSeconds'>
    & { project?: Maybe<(
      { __typename?: 'Project' }
      & Pick<Project, 'id' | 'name'>
    )> }
  )> }
);

export type StopTimerMutationVariables = {
  timerId: Scalars['String'];
};


export type StopTimerMutation = (
  { __typename?: 'Mutation' }
  & { stopTimer?: Maybe<(
    { __typename?: 'Timer' }
    & Pick<Timer, 'id' | 'description' | 'timerStart' | 'timerEnd' | 'isRunning' | 'isBilled' | 'tags' | 'elapsedSeconds'>
    & { project?: Maybe<(
      { __typename?: 'Project' }
      & Pick<Project, 'id' | 'name'>
    )> }
  )> }
);


export const AllClientsDocument = gql`
    query AllClients {
  allClients {
    id
    name
    description
    address
  }
}
    `;
export type AllClientsComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<AllClientsQuery, AllClientsQueryVariables>, 'query'>;

    export const AllClientsComponent = (props: AllClientsComponentProps) => (
      <ApolloReactComponents.Query<AllClientsQuery, AllClientsQueryVariables> query={AllClientsDocument} {...props} />
    );
    
export type AllClientsProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<AllClientsQuery, AllClientsQueryVariables>
    } & TChildProps;
export function withAllClients<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  AllClientsQuery,
  AllClientsQueryVariables,
  AllClientsProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, AllClientsQuery, AllClientsQueryVariables, AllClientsProps<TChildProps, TDataName>>(AllClientsDocument, {
      alias: 'allClients',
      ...operationOptions
    });
};
export type AllClientsQueryResult = ApolloReactCommon.QueryResult<AllClientsQuery, AllClientsQueryVariables>;
export function refetchAllClientsQuery(variables?: AllClientsQueryVariables) {
      return { query: AllClientsDocument, variables: variables }
    }
export const CreateClientDocument = gql`
    mutation createClient($d: ClientInput!) {
  createClient(c: $d) {
    id
    name
    description
    address
  }
}
    `;
export type CreateClientMutationFn = ApolloReactCommon.MutationFunction<CreateClientMutation, CreateClientMutationVariables>;
export type CreateClientComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<CreateClientMutation, CreateClientMutationVariables>, 'mutation'>;

    export const CreateClientComponent = (props: CreateClientComponentProps) => (
      <ApolloReactComponents.Mutation<CreateClientMutation, CreateClientMutationVariables> mutation={CreateClientDocument} {...props} />
    );
    
export type CreateClientProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<CreateClientMutation, CreateClientMutationVariables>
    } & TChildProps;
export function withCreateClient<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  CreateClientMutation,
  CreateClientMutationVariables,
  CreateClientProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, CreateClientMutation, CreateClientMutationVariables, CreateClientProps<TChildProps, TDataName>>(CreateClientDocument, {
      alias: 'createClient',
      ...operationOptions
    });
};
export type CreateClientMutationResult = ApolloReactCommon.MutationResult<CreateClientMutation>;
export type CreateClientMutationOptions = ApolloReactCommon.BaseMutationOptions<CreateClientMutation, CreateClientMutationVariables>;
export const UpdateClientDocument = gql`
    mutation updateClient($d: ClientInput!) {
  updateClient(c: $d) {
    id
    name
    description
    address
  }
}
    `;
export type UpdateClientMutationFn = ApolloReactCommon.MutationFunction<UpdateClientMutation, UpdateClientMutationVariables>;
export type UpdateClientComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<UpdateClientMutation, UpdateClientMutationVariables>, 'mutation'>;

    export const UpdateClientComponent = (props: UpdateClientComponentProps) => (
      <ApolloReactComponents.Mutation<UpdateClientMutation, UpdateClientMutationVariables> mutation={UpdateClientDocument} {...props} />
    );
    
export type UpdateClientProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<UpdateClientMutation, UpdateClientMutationVariables>
    } & TChildProps;
export function withUpdateClient<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  UpdateClientMutation,
  UpdateClientMutationVariables,
  UpdateClientProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, UpdateClientMutation, UpdateClientMutationVariables, UpdateClientProps<TChildProps, TDataName>>(UpdateClientDocument, {
      alias: 'updateClient',
      ...operationOptions
    });
};
export type UpdateClientMutationResult = ApolloReactCommon.MutationResult<UpdateClientMutation>;
export type UpdateClientMutationOptions = ApolloReactCommon.BaseMutationOptions<UpdateClientMutation, UpdateClientMutationVariables>;
export const DeleteClientDocument = gql`
    mutation deleteClient($d: String!) {
  deleteClient(clientId: $d) {
    id
  }
}
    `;
export type DeleteClientMutationFn = ApolloReactCommon.MutationFunction<DeleteClientMutation, DeleteClientMutationVariables>;
export type DeleteClientComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<DeleteClientMutation, DeleteClientMutationVariables>, 'mutation'>;

    export const DeleteClientComponent = (props: DeleteClientComponentProps) => (
      <ApolloReactComponents.Mutation<DeleteClientMutation, DeleteClientMutationVariables> mutation={DeleteClientDocument} {...props} />
    );
    
export type DeleteClientProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<DeleteClientMutation, DeleteClientMutationVariables>
    } & TChildProps;
export function withDeleteClient<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  DeleteClientMutation,
  DeleteClientMutationVariables,
  DeleteClientProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, DeleteClientMutation, DeleteClientMutationVariables, DeleteClientProps<TChildProps, TDataName>>(DeleteClientDocument, {
      alias: 'deleteClient',
      ...operationOptions
    });
};
export type DeleteClientMutationResult = ApolloReactCommon.MutationResult<DeleteClientMutation>;
export type DeleteClientMutationOptions = ApolloReactCommon.BaseMutationOptions<DeleteClientMutation, DeleteClientMutationVariables>;
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
export function refetchGetUsersQuery(variables?: GetUsersQueryVariables) {
      return { query: GetUsersDocument, variables: variables }
    }
export const AllProjectsDocument = gql`
    query AllProjects {
  allProjects {
    name
    description
    id
    team
    status
    client {
      id
      name
      description
    }
  }
}
    `;
export type AllProjectsComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<AllProjectsQuery, AllProjectsQueryVariables>, 'query'>;

    export const AllProjectsComponent = (props: AllProjectsComponentProps) => (
      <ApolloReactComponents.Query<AllProjectsQuery, AllProjectsQueryVariables> query={AllProjectsDocument} {...props} />
    );
    
export type AllProjectsProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<AllProjectsQuery, AllProjectsQueryVariables>
    } & TChildProps;
export function withAllProjects<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  AllProjectsQuery,
  AllProjectsQueryVariables,
  AllProjectsProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, AllProjectsQuery, AllProjectsQueryVariables, AllProjectsProps<TChildProps, TDataName>>(AllProjectsDocument, {
      alias: 'allProjects',
      ...operationOptions
    });
};
export type AllProjectsQueryResult = ApolloReactCommon.QueryResult<AllProjectsQuery, AllProjectsQueryVariables>;
export function refetchAllProjectsQuery(variables?: AllProjectsQueryVariables) {
      return { query: AllProjectsDocument, variables: variables }
    }
export const CreateProjectDocument = gql`
    mutation createProject($d: ProjectInput!) {
  createProject(p: $d) {
    name
    description
    id
    team
    status
    client {
      id
      name
      description
    }
  }
}
    `;
export type CreateProjectMutationFn = ApolloReactCommon.MutationFunction<CreateProjectMutation, CreateProjectMutationVariables>;
export type CreateProjectComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<CreateProjectMutation, CreateProjectMutationVariables>, 'mutation'>;

    export const CreateProjectComponent = (props: CreateProjectComponentProps) => (
      <ApolloReactComponents.Mutation<CreateProjectMutation, CreateProjectMutationVariables> mutation={CreateProjectDocument} {...props} />
    );
    
export type CreateProjectProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<CreateProjectMutation, CreateProjectMutationVariables>
    } & TChildProps;
export function withCreateProject<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  CreateProjectMutation,
  CreateProjectMutationVariables,
  CreateProjectProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, CreateProjectMutation, CreateProjectMutationVariables, CreateProjectProps<TChildProps, TDataName>>(CreateProjectDocument, {
      alias: 'createProject',
      ...operationOptions
    });
};
export type CreateProjectMutationResult = ApolloReactCommon.MutationResult<CreateProjectMutation>;
export type CreateProjectMutationOptions = ApolloReactCommon.BaseMutationOptions<CreateProjectMutation, CreateProjectMutationVariables>;
export const UpdateProjectDocument = gql`
    mutation updateProject($d: ProjectInput!) {
  updateProject(p: $d) {
    name
    description
    id
    team
    status
    client {
      id
      name
      description
    }
  }
}
    `;
export type UpdateProjectMutationFn = ApolloReactCommon.MutationFunction<UpdateProjectMutation, UpdateProjectMutationVariables>;
export type UpdateProjectComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<UpdateProjectMutation, UpdateProjectMutationVariables>, 'mutation'>;

    export const UpdateProjectComponent = (props: UpdateProjectComponentProps) => (
      <ApolloReactComponents.Mutation<UpdateProjectMutation, UpdateProjectMutationVariables> mutation={UpdateProjectDocument} {...props} />
    );
    
export type UpdateProjectProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<UpdateProjectMutation, UpdateProjectMutationVariables>
    } & TChildProps;
export function withUpdateProject<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  UpdateProjectMutation,
  UpdateProjectMutationVariables,
  UpdateProjectProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, UpdateProjectMutation, UpdateProjectMutationVariables, UpdateProjectProps<TChildProps, TDataName>>(UpdateProjectDocument, {
      alias: 'updateProject',
      ...operationOptions
    });
};
export type UpdateProjectMutationResult = ApolloReactCommon.MutationResult<UpdateProjectMutation>;
export type UpdateProjectMutationOptions = ApolloReactCommon.BaseMutationOptions<UpdateProjectMutation, UpdateProjectMutationVariables>;
export const DeleteProjectDocument = gql`
    mutation deleteProject($d: String!) {
  deleteProject(projectId: $d) {
    id
  }
}
    `;
export type DeleteProjectMutationFn = ApolloReactCommon.MutationFunction<DeleteProjectMutation, DeleteProjectMutationVariables>;
export type DeleteProjectComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<DeleteProjectMutation, DeleteProjectMutationVariables>, 'mutation'>;

    export const DeleteProjectComponent = (props: DeleteProjectComponentProps) => (
      <ApolloReactComponents.Mutation<DeleteProjectMutation, DeleteProjectMutationVariables> mutation={DeleteProjectDocument} {...props} />
    );
    
export type DeleteProjectProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<DeleteProjectMutation, DeleteProjectMutationVariables>
    } & TChildProps;
export function withDeleteProject<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  DeleteProjectMutation,
  DeleteProjectMutationVariables,
  DeleteProjectProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, DeleteProjectMutation, DeleteProjectMutationVariables, DeleteProjectProps<TChildProps, TDataName>>(DeleteProjectDocument, {
      alias: 'deleteProject',
      ...operationOptions
    });
};
export type DeleteProjectMutationResult = ApolloReactCommon.MutationResult<DeleteProjectMutation>;
export type DeleteProjectMutationOptions = ApolloReactCommon.BaseMutationOptions<DeleteProjectMutation, DeleteProjectMutationVariables>;
export const AllTimerDocument = gql`
    query AllTimer {
  allTimer(filter: {dayrange: "1"}) {
    description
    id
    timerStart
    timerEnd
    isRunning
    isBilled
    tags
    elapsedSeconds
    project {
      id
      name
    }
  }
}
    `;
export type AllTimerComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<AllTimerQuery, AllTimerQueryVariables>, 'query'>;

    export const AllTimerComponent = (props: AllTimerComponentProps) => (
      <ApolloReactComponents.Query<AllTimerQuery, AllTimerQueryVariables> query={AllTimerDocument} {...props} />
    );
    
export type AllTimerProps<TChildProps = {}, TDataName extends string = 'data'> = {
      [key in TDataName]: ApolloReactHoc.DataValue<AllTimerQuery, AllTimerQueryVariables>
    } & TChildProps;
export function withAllTimer<TProps, TChildProps = {}, TDataName extends string = 'data'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  AllTimerQuery,
  AllTimerQueryVariables,
  AllTimerProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withQuery<TProps, AllTimerQuery, AllTimerQueryVariables, AllTimerProps<TChildProps, TDataName>>(AllTimerDocument, {
      alias: 'allTimer',
      ...operationOptions
    });
};
export type AllTimerQueryResult = ApolloReactCommon.QueryResult<AllTimerQuery, AllTimerQueryVariables>;
export function refetchAllTimerQuery(variables?: AllTimerQueryVariables) {
      return { query: AllTimerDocument, variables: variables }
    }
export const CreateTimerDocument = gql`
    mutation createTimer($d: CreateTimerInput!) {
  createTimer(t: $d) {
    description
    id
    timerStart
    timerEnd
    isRunning
    isBilled
    tags
    elapsedSeconds
    project {
      id
      name
    }
  }
}
    `;
export type CreateTimerMutationFn = ApolloReactCommon.MutationFunction<CreateTimerMutation, CreateTimerMutationVariables>;
export type CreateTimerComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<CreateTimerMutation, CreateTimerMutationVariables>, 'mutation'>;

    export const CreateTimerComponent = (props: CreateTimerComponentProps) => (
      <ApolloReactComponents.Mutation<CreateTimerMutation, CreateTimerMutationVariables> mutation={CreateTimerDocument} {...props} />
    );
    
export type CreateTimerProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<CreateTimerMutation, CreateTimerMutationVariables>
    } & TChildProps;
export function withCreateTimer<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  CreateTimerMutation,
  CreateTimerMutationVariables,
  CreateTimerProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, CreateTimerMutation, CreateTimerMutationVariables, CreateTimerProps<TChildProps, TDataName>>(CreateTimerDocument, {
      alias: 'createTimer',
      ...operationOptions
    });
};
export type CreateTimerMutationResult = ApolloReactCommon.MutationResult<CreateTimerMutation>;
export type CreateTimerMutationOptions = ApolloReactCommon.BaseMutationOptions<CreateTimerMutation, CreateTimerMutationVariables>;
export const UpdateTimerDocument = gql`
    mutation updateTimer($d: TimerInput!) {
  updateTimer(t: $d) {
    description
    id
    timerStart
    timerEnd
    isRunning
    isBilled
    tags
    elapsedSeconds
    project {
      id
      name
    }
  }
}
    `;
export type UpdateTimerMutationFn = ApolloReactCommon.MutationFunction<UpdateTimerMutation, UpdateTimerMutationVariables>;
export type UpdateTimerComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<UpdateTimerMutation, UpdateTimerMutationVariables>, 'mutation'>;

    export const UpdateTimerComponent = (props: UpdateTimerComponentProps) => (
      <ApolloReactComponents.Mutation<UpdateTimerMutation, UpdateTimerMutationVariables> mutation={UpdateTimerDocument} {...props} />
    );
    
export type UpdateTimerProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<UpdateTimerMutation, UpdateTimerMutationVariables>
    } & TChildProps;
export function withUpdateTimer<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  UpdateTimerMutation,
  UpdateTimerMutationVariables,
  UpdateTimerProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, UpdateTimerMutation, UpdateTimerMutationVariables, UpdateTimerProps<TChildProps, TDataName>>(UpdateTimerDocument, {
      alias: 'updateTimer',
      ...operationOptions
    });
};
export type UpdateTimerMutationResult = ApolloReactCommon.MutationResult<UpdateTimerMutation>;
export type UpdateTimerMutationOptions = ApolloReactCommon.BaseMutationOptions<UpdateTimerMutation, UpdateTimerMutationVariables>;
export const DeleteTimerDocument = gql`
    mutation deleteTimer($timerId: String!) {
  deleteTimer(timerId: $timerId) {
    id
  }
}
    `;
export type DeleteTimerMutationFn = ApolloReactCommon.MutationFunction<DeleteTimerMutation, DeleteTimerMutationVariables>;
export type DeleteTimerComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<DeleteTimerMutation, DeleteTimerMutationVariables>, 'mutation'>;

    export const DeleteTimerComponent = (props: DeleteTimerComponentProps) => (
      <ApolloReactComponents.Mutation<DeleteTimerMutation, DeleteTimerMutationVariables> mutation={DeleteTimerDocument} {...props} />
    );
    
export type DeleteTimerProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<DeleteTimerMutation, DeleteTimerMutationVariables>
    } & TChildProps;
export function withDeleteTimer<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  DeleteTimerMutation,
  DeleteTimerMutationVariables,
  DeleteTimerProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, DeleteTimerMutation, DeleteTimerMutationVariables, DeleteTimerProps<TChildProps, TDataName>>(DeleteTimerDocument, {
      alias: 'deleteTimer',
      ...operationOptions
    });
};
export type DeleteTimerMutationResult = ApolloReactCommon.MutationResult<DeleteTimerMutation>;
export type DeleteTimerMutationOptions = ApolloReactCommon.BaseMutationOptions<DeleteTimerMutation, DeleteTimerMutationVariables>;
export const StartTimerDocument = gql`
    mutation startTimer($timerId: String!) {
  startTimer(timerId: $timerId) {
    id
    description
    timerStart
    timerEnd
    isRunning
    isBilled
    tags
    elapsedSeconds
    project {
      id
      name
    }
  }
}
    `;
export type StartTimerMutationFn = ApolloReactCommon.MutationFunction<StartTimerMutation, StartTimerMutationVariables>;
export type StartTimerComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<StartTimerMutation, StartTimerMutationVariables>, 'mutation'>;

    export const StartTimerComponent = (props: StartTimerComponentProps) => (
      <ApolloReactComponents.Mutation<StartTimerMutation, StartTimerMutationVariables> mutation={StartTimerDocument} {...props} />
    );
    
export type StartTimerProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<StartTimerMutation, StartTimerMutationVariables>
    } & TChildProps;
export function withStartTimer<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  StartTimerMutation,
  StartTimerMutationVariables,
  StartTimerProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, StartTimerMutation, StartTimerMutationVariables, StartTimerProps<TChildProps, TDataName>>(StartTimerDocument, {
      alias: 'startTimer',
      ...operationOptions
    });
};
export type StartTimerMutationResult = ApolloReactCommon.MutationResult<StartTimerMutation>;
export type StartTimerMutationOptions = ApolloReactCommon.BaseMutationOptions<StartTimerMutation, StartTimerMutationVariables>;
export const StopTimerDocument = gql`
    mutation stopTimer($timerId: String!) {
  stopTimer(timerId: $timerId) {
    id
    description
    timerStart
    timerEnd
    isRunning
    isBilled
    tags
    elapsedSeconds
    project {
      id
      name
    }
  }
}
    `;
export type StopTimerMutationFn = ApolloReactCommon.MutationFunction<StopTimerMutation, StopTimerMutationVariables>;
export type StopTimerComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<StopTimerMutation, StopTimerMutationVariables>, 'mutation'>;

    export const StopTimerComponent = (props: StopTimerComponentProps) => (
      <ApolloReactComponents.Mutation<StopTimerMutation, StopTimerMutationVariables> mutation={StopTimerDocument} {...props} />
    );
    
export type StopTimerProps<TChildProps = {}, TDataName extends string = 'mutate'> = {
      [key in TDataName]: ApolloReactCommon.MutationFunction<StopTimerMutation, StopTimerMutationVariables>
    } & TChildProps;
export function withStopTimer<TProps, TChildProps = {}, TDataName extends string = 'mutate'>(operationOptions?: ApolloReactHoc.OperationOption<
  TProps,
  StopTimerMutation,
  StopTimerMutationVariables,
  StopTimerProps<TChildProps, TDataName>>) {
    return ApolloReactHoc.withMutation<TProps, StopTimerMutation, StopTimerMutationVariables, StopTimerProps<TChildProps, TDataName>>(StopTimerDocument, {
      alias: 'stopTimer',
      ...operationOptions
    });
};
export type StopTimerMutationResult = ApolloReactCommon.MutationResult<StopTimerMutation>;
export type StopTimerMutationOptions = ApolloReactCommon.BaseMutationOptions<StopTimerMutation, StopTimerMutationVariables>;