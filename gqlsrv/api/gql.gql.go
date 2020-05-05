package gql

type AllTimerFilter struct {
	Dayrange *string
}

type Client struct {
	ID          string
	Description string
	Name        string
	Address     string
}

type ClientResolver struct {
	R *Client
}

func (r ClientResolver) ID() *string {
	return &r.R.ID
}

func (r ClientResolver) Description() *string {
	return &r.R.Description
}

func (r ClientResolver) Name() *string {
	return &r.R.Name
}

func (r ClientResolver) Address() *string {
	return &r.R.Address
}

type ClientInput struct {
	ID          *string
	Description *string
	Name        *string
	Address     *string
}

type CreateTimerInput struct {
	Description *string
	Project     *string
}

type Project struct {
	Status      string
	ID          string
	Name        string
	Description string
	Team        string
	Client      *Client
}

func (r ProjectResolver) ID() *string {
	return &r.R.ID
}

func (r ProjectResolver) Name() *string {
	return &r.R.Name
}

func (r ProjectResolver) Description() *string {
	return &r.R.Description
}

func (r ProjectResolver) Team() *string {
	return &r.R.Team
}

func (r ProjectResolver) Status() *string {
	return &r.R.Status
}

type ProjectInput struct {
	ID          *string
	Name        *string
	Description *string
	Team        *string
	Client      *string
	Status      *string
}

type Timer struct {
	Client         *Client
	Tags           string
	TimerStart     string
	TimerEnd       string
	IsBilled       bool
	Name           string
	Teammember     string
	Project        *Project
	ElapsedSeconds int32
	IsRunning      bool
	ID             string
	Description    string
}

// type TimerResolver struct {
// 	R *Timer
// }

func (r TimerResolver) ID() *string {
	return &r.R.ID
}

func (r TimerResolver) Description() *string {
	return &r.R.Description
}

// func (r TimerResolver) Project() *ProjectResolver {
// 	return &ProjectResolver{r.R.Project}
// }

func (r TimerResolver) ElapsedSeconds() *int32 {
	return &r.R.ElapsedSeconds
}

func (r TimerResolver) IsRunning() *bool {
	return &r.R.IsRunning
}

func (r TimerResolver) TimerEnd() *string {
	return &r.R.TimerEnd
}

func (r TimerResolver) IsBilled() *bool {
	return &r.R.IsBilled
}

func (r TimerResolver) Name() *string {
	return &r.R.Name
}

func (r TimerResolver) Teammember() *string {
	return &r.R.Teammember
}

func (r TimerResolver) Tags() *string {
	return &r.R.Tags
}

func (r TimerResolver) TimerStart() *string {
	return &r.R.TimerStart
}

type TimerInput struct {
	Name           *string
	Description    *string
	Client         *string
	Project        *string
	Tags           *string
	ElapsedSeconds *int32
	IsBilled       *bool
	ID             *string
	TimerStart     *string
	TimerEnd       *string
	IsRunning      *bool
	Teammember     *string
}

type User struct {
	ID    string
	Name  string
	Email string
}

type UserResolver struct {
	R *User
}

func (r UserResolver) ID() *string {
	return &r.R.ID
}

func (r UserResolver) Name() *string {
	return &r.R.Name
}

func (r UserResolver) Email() *string {
	return &r.R.Email
}

type UserInput struct {
	Name     *string
	Email    *string
	Password *string
}

type CreateUserRequest struct {
	User UserInput
}

type StartTimerRequest struct {
	TimerId string
}

type StopTimerRequest struct {
	TimerId string
}

type UpdateTimerRequest struct {
	T TimerInput
}

type CreateProjectRequest struct {
	P ProjectInput
}

type UpdateProjectRequest struct {
	P ProjectInput
}

type UpdateClientRequest struct {
	C ClientInput
}

type DeleteClientRequest struct {
	ClientId string
}

type CreateTimerRequest struct {
	T CreateTimerInput
}

type DeleteTimerRequest struct {
	TimerId string
}

type DeleteProjectRequest struct {
	ProjectId string
}

type CreateClientRequest struct {
	C ClientInput
}

type GetTimerRequest struct {
	ID *string
}

type GetProjectRequest struct {
	ID *string
}

type AllTimerRequest struct {
	Filter AllTimerFilter
}

type GetClientRequest struct {
	ID *string
}

type HelloRequest struct {
	Name string
}

type UserRequest struct {
	ID *string
}

type GqlResolver interface {
	CreateClient(CreateClientRequest) *ClientResolver
	User(UserRequest) *UserResolver
	AllClients() *[]*ClientResolver
	AllTimer(AllTimerRequest) *[]*TimerResolver
	UpdateClient(UpdateClientRequest) *ClientResolver
	StopTimer(StopTimerRequest) *TimerResolver
	CreateProject(CreateProjectRequest) *ProjectResolver
	CreateTimer(CreateTimerRequest) *TimerResolver
	Hello(HelloRequest) string
	GetTimer(GetTimerRequest) *TimerResolver
	CreateUser(CreateUserRequest) *string
	DeleteClient(DeleteClientRequest) *ClientResolver
	DeleteProject(DeleteProjectRequest) *ProjectResolver
	GetClient(GetClientRequest) *ClientResolver
	AllUsers() *[]*UserResolver
	RunningTimer() *TimerResolver
	UpdateTimer(UpdateTimerRequest) *TimerResolver
	UpdateProject(UpdateProjectRequest) *ProjectResolver
	DeleteTimer(DeleteTimerRequest) *TimerResolver
	AllProjects() *[]*ProjectResolver
	GetProject(GetProjectRequest) *ProjectResolver
	StartTimer(StartTimerRequest) *TimerResolver
}

var Schema = `

schema {
  query: Query
  mutation: Mutation
}
type Query {
  hello(name: String!): String!
  User(id: String): User
  allUsers: [User]

  allTimer(filter: AllTimerFilter!): [Timer]
  runningTimer: Timer
  getTimer(id: String): Timer

  getProject(id: String): Project
  allProjects: [Project]

  getClient(id: String): Client
  allClients: [Client]
}

type Mutation {
  createUser(user: UserInput!): String
  createTimer(t: CreateTimerInput!): Timer
  startTimer(timerId: String!): Timer
  stopTimer(timerId: String!): Timer
  deleteTimer(timerId: String!): Timer
  updateTimer(t: TimerInput!): Timer

  createProject(p: ProjectInput!): Project
  updateProject(p: ProjectInput!): Project
  deleteProject(projectId: String!): Project

  createClient(c: ClientInput!): Client
  updateClient(c: ClientInput!): Client
  deleteClient(clientId: String!): Client
}

type User {
  id: String
  name: String
  email: String
}
input UserInput {
  name: String
  email: String
  password: String
}

input AllTimerFilter {
  dayrange: String
}

input CreateTimerInput {
  description: String
  project: String
}
input TimerInput {
  id: String
  name: String
  description: String
  teammember: String
  client: String
  project: String
  tags: String
  elapsedSeconds: Int
  timerStart: String
  timerEnd: String
  isRunning: Boolean
  isBilled: Boolean
}

type Timer {
  id: String
  name: String
  description: String
  teammember: String
  client: Client
  project: Project
  tags: String
  elapsedSeconds: Int
  timerStart: String
  timerEnd: String
  isRunning: Boolean
  isBilled: Boolean
}

input ProjectInput {
  id: String
  name: String
  description: String
  team: String
  client: String
  status: String
}
type Project {
  id: String
  name: String
  description: String
  team: String
  client: Client
  status: String
}

input ClientInput {
  id: String
  description: String
  name: String
  address: String
}

type Client {
  id: String
  description: String
  name: String
  address: String
}

`
