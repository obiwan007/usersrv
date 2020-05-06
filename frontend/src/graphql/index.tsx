import * as ApolloReactCommon from "@apollo/react-common";
import * as ApolloReactComponents from "@apollo/react-components";
import * as ApolloReactHoc from "@apollo/react-hoc";
import gql from "graphql-tag";
import * as React from "react";
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
  dayrange?: Maybe<Scalars["String"]>;
};

export type Client = {
  __typename?: "Client";
  id?: Maybe<Scalars["String"]>;
  description?: Maybe<Scalars["String"]>;
  name?: Maybe<Scalars["String"]>;
  address?: Maybe<Scalars["String"]>;
};

export type ClientInput = {
  id?: Maybe<Scalars["String"]>;
  description?: Maybe<Scalars["String"]>;
  name?: Maybe<Scalars["String"]>;
  address?: Maybe<Scalars["String"]>;
};

export type CreateTimerInput = {
  description?: Maybe<Scalars["String"]>;
  project?: Maybe<Scalars["String"]>;
};

export type Mutation = {
  __typename?: "Mutation";
  createUser?: Maybe<Scalars["String"]>;
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
  timerId: Scalars["String"];
};

export type MutationStopTimerArgs = {
  timerId: Scalars["String"];
};

export type MutationDeleteTimerArgs = {
  timerId: Scalars["String"];
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
  projectId: Scalars["String"];
};

export type MutationCreateClientArgs = {
  c: ClientInput;
};

export type MutationUpdateClientArgs = {
  c: ClientInput;
};

export type MutationDeleteClientArgs = {
  clientId: Scalars["String"];
};

export type Project = {
  __typename?: "Project";
  id?: Maybe<Scalars["String"]>;
  name?: Maybe<Scalars["String"]>;
  description?: Maybe<Scalars["String"]>;
  team?: Maybe<Scalars["String"]>;
  client?: Maybe<Client>;
  status?: Maybe<Scalars["String"]>;
};

export type ProjectInput = {
  id?: Maybe<Scalars["String"]>;
  name?: Maybe<Scalars["String"]>;
  description?: Maybe<Scalars["String"]>;
  team?: Maybe<Scalars["String"]>;
  client?: Maybe<Scalars["String"]>;
  status?: Maybe<Scalars["String"]>;
};

export type Query = {
  __typename?: "Query";
  hello: Scalars["String"];
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
  name: Scalars["String"];
};

export type QueryUserArgs = {
  id?: Maybe<Scalars["String"]>;
};

export type QueryAllTimerArgs = {
  filter: AllTimerFilter;
};

export type QueryGetTimerArgs = {
  id?: Maybe<Scalars["String"]>;
};

export type QueryGetProjectArgs = {
  id?: Maybe<Scalars["String"]>;
};

export type QueryGetClientArgs = {
  id?: Maybe<Scalars["String"]>;
};

export type Timer = {
  __typename?: "Timer";
  id?: Maybe<Scalars["String"]>;
  name?: Maybe<Scalars["String"]>;
  description?: Maybe<Scalars["String"]>;
  teammember?: Maybe<Scalars["String"]>;
  client?: Maybe<Client>;
  project?: Maybe<Project>;
  tags?: Maybe<Scalars["String"]>;
  elapsedSeconds?: Maybe<Scalars["Int"]>;
  timerStart?: Maybe<Scalars["String"]>;
  timerEnd?: Maybe<Scalars["String"]>;
  isRunning?: Maybe<Scalars["Boolean"]>;
  isBilled?: Maybe<Scalars["Boolean"]>;
};

export type TimerInput = {
  id?: Maybe<Scalars["String"]>;
  name?: Maybe<Scalars["String"]>;
  description?: Maybe<Scalars["String"]>;
  teammember?: Maybe<Scalars["String"]>;
  client?: Maybe<Scalars["String"]>;
  project?: Maybe<Scalars["String"]>;
  tags?: Maybe<Scalars["String"]>;
  elapsedSeconds?: Maybe<Scalars["Int"]>;
  timerStart?: Maybe<Scalars["String"]>;
  timerEnd?: Maybe<Scalars["String"]>;
  isRunning?: Maybe<Scalars["Boolean"]>;
  isBilled?: Maybe<Scalars["Boolean"]>;
};

export type User = {
  __typename?: "User";
  id?: Maybe<Scalars["String"]>;
  name?: Maybe<Scalars["String"]>;
  email?: Maybe<Scalars["String"]>;
};

export type UserInput = {
  name?: Maybe<Scalars["String"]>;
  email?: Maybe<Scalars["String"]>;
  password?: Maybe<Scalars["String"]>;
};

export type AllClientsQueryVariables = {};

export type AllClientsQuery = { __typename?: "Query" } & {
  allClients?: Maybe<
    Array<
      Maybe<
        { __typename?: "Client" } & Pick<
          Client,
          "id" | "name" | "description" | "address"
        >
      >
    >
  >;
};

export type CreateClientMutationVariables = {
  d: ClientInput;
};

export type CreateClientMutation = { __typename?: "Mutation" } & {
  createClient?: Maybe<
    { __typename?: "Client" } & Pick<
      Client,
      "id" | "name" | "description" | "address"
    >
  >;
};

export type UpdateClientMutationVariables = {
  d: ClientInput;
};

export type UpdateClientMutation = { __typename?: "Mutation" } & {
  updateClient?: Maybe<
    { __typename?: "Client" } & Pick<
      Client,
      "id" | "name" | "description" | "address"
    >
  >;
};

export type DeleteClientMutationVariables = {
  d: Scalars["String"];
};

export type DeleteClientMutation = { __typename?: "Mutation" } & {
  deleteClient?: Maybe<{ __typename?: "Client" } & Pick<Client, "id">>;
};

export type GetUsersQueryVariables = {};

export type GetUsersQuery = { __typename?: "Query" } & {
  allUsers?: Maybe<
    Array<Maybe<{ __typename?: "User" } & Pick<User, "name" | "email" | "id">>>
  >;
};

export type AllProjectsQueryVariables = {};

export type AllProjectsQuery = { __typename?: "Query" } & {
  allProjects?: Maybe<
    Array<
      Maybe<
        { __typename?: "Project" } & Pick<
          Project,
          "name" | "description" | "id" | "team" | "status"
        > & {
            client?: Maybe<
              { __typename?: "Client" } & Pick<
                Client,
                "id" | "name" | "description"
              >
            >;
          }
      >
    >
  >;
};

export type CreateProjectMutationVariables = {
  d: ProjectInput;
};

export type CreateProjectMutation = { __typename?: "Mutation" } & {
  createProject?: Maybe<
    { __typename?: "Project" } & Pick<
      Project,
      "name" | "description" | "id" | "team" | "status"
    > & {
        client?: Maybe<
          { __typename?: "Client" } & Pick<
            Client,
            "id" | "name" | "description"
          >
        >;
      }
  >;
};

export type UpdateProjectMutationVariables = {
  d: ProjectInput;
};

export type UpdateProjectMutation = { __typename?: "Mutation" } & {
  updateProject?: Maybe<
    { __typename?: "Project" } & Pick<
      Project,
      "name" | "description" | "id" | "team" | "status"
    > & {
        client?: Maybe<
          { __typename?: "Client" } & Pick<
            Client,
            "id" | "name" | "description"
          >
        >;
      }
  >;
};

export type DeleteProjectMutationVariables = {
  d: Scalars["String"];
};

export type DeleteProjectMutation = { __typename?: "Mutation" } & {
  deleteProject?: Maybe<{ __typename?: "Project" } & Pick<Project, "id">>;
};

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
export type AllClientsComponentProps = Omit<
  ApolloReactComponents.QueryComponentOptions<
    AllClientsQuery,
    AllClientsQueryVariables
  >,
  "query"
>;

export const AllClientsComponent = (props: AllClientsComponentProps) => (
  <ApolloReactComponents.Query<AllClientsQuery, AllClientsQueryVariables>
    query={AllClientsDocument}
    {...props}
  />
);

export type AllClientsProps<
  TChildProps = {},
  TDataName extends string = "data"
> = {
  [key in TDataName]: ApolloReactHoc.DataValue<
    AllClientsQuery,
    AllClientsQueryVariables
  >;
} &
  TChildProps;
export function withAllClients<
  TProps,
  TChildProps = {},
  TDataName extends string = "data"
>(
  operationOptions?: ApolloReactHoc.OperationOption<
    TProps,
    AllClientsQuery,
    AllClientsQueryVariables,
    AllClientsProps<TChildProps, TDataName>
  >
) {
  return ApolloReactHoc.withQuery<
    TProps,
    AllClientsQuery,
    AllClientsQueryVariables,
    AllClientsProps<TChildProps, TDataName>
  >(AllClientsDocument, {
    alias: "allClients",
    ...operationOptions,
  });
}
export type AllClientsQueryResult = ApolloReactCommon.QueryResult<
  AllClientsQuery,
  AllClientsQueryVariables
>;
export function refetchAllClientsQuery(variables?: AllClientsQueryVariables) {
  return { query: AllClientsDocument, variables: variables };
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
export type CreateClientMutationFn = ApolloReactCommon.MutationFunction<
  CreateClientMutation,
  CreateClientMutationVariables
>;
export type CreateClientComponentProps = Omit<
  ApolloReactComponents.MutationComponentOptions<
    CreateClientMutation,
    CreateClientMutationVariables
  >,
  "mutation"
>;

export const CreateClientComponent = (props: CreateClientComponentProps) => (
  <ApolloReactComponents.Mutation<
    CreateClientMutation,
    CreateClientMutationVariables
  >
    mutation={CreateClientDocument}
    {...props}
  />
);

export type CreateClientProps<
  TChildProps = {},
  TDataName extends string = "mutate"
> = {
  [key in TDataName]: ApolloReactCommon.MutationFunction<
    CreateClientMutation,
    CreateClientMutationVariables
  >;
} &
  TChildProps;
export function withCreateClient<
  TProps,
  TChildProps = {},
  TDataName extends string = "mutate"
>(
  operationOptions?: ApolloReactHoc.OperationOption<
    TProps,
    CreateClientMutation,
    CreateClientMutationVariables,
    CreateClientProps<TChildProps, TDataName>
  >
) {
  return ApolloReactHoc.withMutation<
    TProps,
    CreateClientMutation,
    CreateClientMutationVariables,
    CreateClientProps<TChildProps, TDataName>
  >(CreateClientDocument, {
    alias: "createClient",
    ...operationOptions,
  });
}
export type CreateClientMutationResult = ApolloReactCommon.MutationResult<
  CreateClientMutation
>;
export type CreateClientMutationOptions = ApolloReactCommon.BaseMutationOptions<
  CreateClientMutation,
  CreateClientMutationVariables
>;
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
export type UpdateClientMutationFn = ApolloReactCommon.MutationFunction<
  UpdateClientMutation,
  UpdateClientMutationVariables
>;
export type UpdateClientComponentProps = Omit<
  ApolloReactComponents.MutationComponentOptions<
    UpdateClientMutation,
    UpdateClientMutationVariables
  >,
  "mutation"
>;

export const UpdateClientComponent = (props: UpdateClientComponentProps) => (
  <ApolloReactComponents.Mutation<
    UpdateClientMutation,
    UpdateClientMutationVariables
  >
    mutation={UpdateClientDocument}
    {...props}
  />
);

export type UpdateClientProps<
  TChildProps = {},
  TDataName extends string = "mutate"
> = {
  [key in TDataName]: ApolloReactCommon.MutationFunction<
    UpdateClientMutation,
    UpdateClientMutationVariables
  >;
} &
  TChildProps;
export function withUpdateClient<
  TProps,
  TChildProps = {},
  TDataName extends string = "mutate"
>(
  operationOptions?: ApolloReactHoc.OperationOption<
    TProps,
    UpdateClientMutation,
    UpdateClientMutationVariables,
    UpdateClientProps<TChildProps, TDataName>
  >
) {
  return ApolloReactHoc.withMutation<
    TProps,
    UpdateClientMutation,
    UpdateClientMutationVariables,
    UpdateClientProps<TChildProps, TDataName>
  >(UpdateClientDocument, {
    alias: "updateClient",
    ...operationOptions,
  });
}
export type UpdateClientMutationResult = ApolloReactCommon.MutationResult<
  UpdateClientMutation
>;
export type UpdateClientMutationOptions = ApolloReactCommon.BaseMutationOptions<
  UpdateClientMutation,
  UpdateClientMutationVariables
>;
export const DeleteClientDocument = gql`
  mutation deleteClient($d: String!) {
    deleteClient(clientId: $d) {
      id
    }
  }
`;
export type DeleteClientMutationFn = ApolloReactCommon.MutationFunction<
  DeleteClientMutation,
  DeleteClientMutationVariables
>;
export type DeleteClientComponentProps = Omit<
  ApolloReactComponents.MutationComponentOptions<
    DeleteClientMutation,
    DeleteClientMutationVariables
  >,
  "mutation"
>;

export const DeleteClientComponent = (props: DeleteClientComponentProps) => (
  <ApolloReactComponents.Mutation<
    DeleteClientMutation,
    DeleteClientMutationVariables
  >
    mutation={DeleteClientDocument}
    {...props}
  />
);

export type DeleteClientProps<
  TChildProps = {},
  TDataName extends string = "mutate"
> = {
  [key in TDataName]: ApolloReactCommon.MutationFunction<
    DeleteClientMutation,
    DeleteClientMutationVariables
  >;
} &
  TChildProps;
export function withDeleteClient<
  TProps,
  TChildProps = {},
  TDataName extends string = "mutate"
>(
  operationOptions?: ApolloReactHoc.OperationOption<
    TProps,
    DeleteClientMutation,
    DeleteClientMutationVariables,
    DeleteClientProps<TChildProps, TDataName>
  >
) {
  return ApolloReactHoc.withMutation<
    TProps,
    DeleteClientMutation,
    DeleteClientMutationVariables,
    DeleteClientProps<TChildProps, TDataName>
  >(DeleteClientDocument, {
    alias: "deleteClient",
    ...operationOptions,
  });
}
export type DeleteClientMutationResult = ApolloReactCommon.MutationResult<
  DeleteClientMutation
>;
export type DeleteClientMutationOptions = ApolloReactCommon.BaseMutationOptions<
  DeleteClientMutation,
  DeleteClientMutationVariables
>;
export const GetUsersDocument = gql`
  query GetUsers {
    allUsers {
      name
      email
      id
    }
  }
`;
export type GetUsersComponentProps = Omit<
  ApolloReactComponents.QueryComponentOptions<
    GetUsersQuery,
    GetUsersQueryVariables
  >,
  "query"
>;

export const GetUsersComponent = (props: GetUsersComponentProps) => (
  <ApolloReactComponents.Query<GetUsersQuery, GetUsersQueryVariables>
    query={GetUsersDocument}
    {...props}
  />
);

export type GetUsersProps<
  TChildProps = {},
  TDataName extends string = "data"
> = {
  [key in TDataName]: ApolloReactHoc.DataValue<
    GetUsersQuery,
    GetUsersQueryVariables
  >;
} &
  TChildProps;
export function withGetUsers<
  TProps,
  TChildProps = {},
  TDataName extends string = "data"
>(
  operationOptions?: ApolloReactHoc.OperationOption<
    TProps,
    GetUsersQuery,
    GetUsersQueryVariables,
    GetUsersProps<TChildProps, TDataName>
  >
) {
  return ApolloReactHoc.withQuery<
    TProps,
    GetUsersQuery,
    GetUsersQueryVariables,
    GetUsersProps<TChildProps, TDataName>
  >(GetUsersDocument, {
    alias: "getUsers",
    ...operationOptions,
  });
}
export type GetUsersQueryResult = ApolloReactCommon.QueryResult<
  GetUsersQuery,
  GetUsersQueryVariables
>;
export function refetchGetUsersQuery(variables?: GetUsersQueryVariables) {
  return { query: GetUsersDocument, variables: variables };
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
export type AllProjectsComponentProps = Omit<
  ApolloReactComponents.QueryComponentOptions<
    AllProjectsQuery,
    AllProjectsQueryVariables
  >,
  "query"
>;

export const AllProjectsComponent = (props: AllProjectsComponentProps) => (
  <ApolloReactComponents.Query<AllProjectsQuery, AllProjectsQueryVariables>
    query={AllProjectsDocument}
    {...props}
  />
);

export type AllProjectsProps<
  TChildProps = {},
  TDataName extends string = "data"
> = {
  [key in TDataName]: ApolloReactHoc.DataValue<
    AllProjectsQuery,
    AllProjectsQueryVariables
  >;
} &
  TChildProps;
export function withAllProjects<
  TProps,
  TChildProps = {},
  TDataName extends string = "data"
>(
  operationOptions?: ApolloReactHoc.OperationOption<
    TProps,
    AllProjectsQuery,
    AllProjectsQueryVariables,
    AllProjectsProps<TChildProps, TDataName>
  >
) {
  return ApolloReactHoc.withQuery<
    TProps,
    AllProjectsQuery,
    AllProjectsQueryVariables,
    AllProjectsProps<TChildProps, TDataName>
  >(AllProjectsDocument, {
    alias: "allProjects",
    ...operationOptions,
  });
}
export type AllProjectsQueryResult = ApolloReactCommon.QueryResult<
  AllProjectsQuery,
  AllProjectsQueryVariables
>;
export function refetchAllProjectsQuery(variables?: AllProjectsQueryVariables) {
  return { query: AllProjectsDocument, variables: variables };
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
export type CreateProjectMutationFn = ApolloReactCommon.MutationFunction<
  CreateProjectMutation,
  CreateProjectMutationVariables
>;
export type CreateProjectComponentProps = Omit<
  ApolloReactComponents.MutationComponentOptions<
    CreateProjectMutation,
    CreateProjectMutationVariables
  >,
  "mutation"
>;

export const CreateProjectComponent = (props: CreateProjectComponentProps) => (
  <ApolloReactComponents.Mutation<
    CreateProjectMutation,
    CreateProjectMutationVariables
  >
    mutation={CreateProjectDocument}
    {...props}
  />
);

export type CreateProjectProps<
  TChildProps = {},
  TDataName extends string = "mutate"
> = {
  [key in TDataName]: ApolloReactCommon.MutationFunction<
    CreateProjectMutation,
    CreateProjectMutationVariables
  >;
} &
  TChildProps;
export function withCreateProject<
  TProps,
  TChildProps = {},
  TDataName extends string = "mutate"
>(
  operationOptions?: ApolloReactHoc.OperationOption<
    TProps,
    CreateProjectMutation,
    CreateProjectMutationVariables,
    CreateProjectProps<TChildProps, TDataName>
  >
) {
  return ApolloReactHoc.withMutation<
    TProps,
    CreateProjectMutation,
    CreateProjectMutationVariables,
    CreateProjectProps<TChildProps, TDataName>
  >(CreateProjectDocument, {
    alias: "createProject",
    ...operationOptions,
  });
}
export type CreateProjectMutationResult = ApolloReactCommon.MutationResult<
  CreateProjectMutation
>;
export type CreateProjectMutationOptions = ApolloReactCommon.BaseMutationOptions<
  CreateProjectMutation,
  CreateProjectMutationVariables
>;
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
export type UpdateProjectMutationFn = ApolloReactCommon.MutationFunction<
  UpdateProjectMutation,
  UpdateProjectMutationVariables
>;
export type UpdateProjectComponentProps = Omit<
  ApolloReactComponents.MutationComponentOptions<
    UpdateProjectMutation,
    UpdateProjectMutationVariables
  >,
  "mutation"
>;

export const UpdateProjectComponent = (props: UpdateProjectComponentProps) => (
  <ApolloReactComponents.Mutation<
    UpdateProjectMutation,
    UpdateProjectMutationVariables
  >
    mutation={UpdateProjectDocument}
    {...props}
  />
);

export type UpdateProjectProps<
  TChildProps = {},
  TDataName extends string = "mutate"
> = {
  [key in TDataName]: ApolloReactCommon.MutationFunction<
    UpdateProjectMutation,
    UpdateProjectMutationVariables
  >;
} &
  TChildProps;
export function withUpdateProject<
  TProps,
  TChildProps = {},
  TDataName extends string = "mutate"
>(
  operationOptions?: ApolloReactHoc.OperationOption<
    TProps,
    UpdateProjectMutation,
    UpdateProjectMutationVariables,
    UpdateProjectProps<TChildProps, TDataName>
  >
) {
  return ApolloReactHoc.withMutation<
    TProps,
    UpdateProjectMutation,
    UpdateProjectMutationVariables,
    UpdateProjectProps<TChildProps, TDataName>
  >(UpdateProjectDocument, {
    alias: "updateProject",
    ...operationOptions,
  });
}
export type UpdateProjectMutationResult = ApolloReactCommon.MutationResult<
  UpdateProjectMutation
>;
export type UpdateProjectMutationOptions = ApolloReactCommon.BaseMutationOptions<
  UpdateProjectMutation,
  UpdateProjectMutationVariables
>;
export const DeleteProjectDocument = gql`
  mutation deleteProject($d: String!) {
    deleteProject(projectId: $d) {
      id
    }
  }
`;
export type DeleteProjectMutationFn = ApolloReactCommon.MutationFunction<
  DeleteProjectMutation,
  DeleteProjectMutationVariables
>;
export type DeleteProjectComponentProps = Omit<
  ApolloReactComponents.MutationComponentOptions<
    DeleteProjectMutation,
    DeleteProjectMutationVariables
  >,
  "mutation"
>;

export const DeleteProjectComponent = (props: DeleteProjectComponentProps) => (
  <ApolloReactComponents.Mutation<
    DeleteProjectMutation,
    DeleteProjectMutationVariables
  >
    mutation={DeleteProjectDocument}
    {...props}
  />
);

export type DeleteProjectProps<
  TChildProps = {},
  TDataName extends string = "mutate"
> = {
  [key in TDataName]: ApolloReactCommon.MutationFunction<
    DeleteProjectMutation,
    DeleteProjectMutationVariables
  >;
} &
  TChildProps;
export function withDeleteProject<
  TProps,
  TChildProps = {},
  TDataName extends string = "mutate"
>(
  operationOptions?: ApolloReactHoc.OperationOption<
    TProps,
    DeleteProjectMutation,
    DeleteProjectMutationVariables,
    DeleteProjectProps<TChildProps, TDataName>
  >
) {
  return ApolloReactHoc.withMutation<
    TProps,
    DeleteProjectMutation,
    DeleteProjectMutationVariables,
    DeleteProjectProps<TChildProps, TDataName>
  >(DeleteProjectDocument, {
    alias: "deleteProject",
    ...operationOptions,
  });
}
export type DeleteProjectMutationResult = ApolloReactCommon.MutationResult<
  DeleteProjectMutation
>;
export type DeleteProjectMutationOptions = ApolloReactCommon.BaseMutationOptions<
  DeleteProjectMutation,
  DeleteProjectMutationVariables
>;
