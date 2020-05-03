package types

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

type CreateTimerInput struct {
	Project     *string
	Description *string
}

type Project struct {
	Status      string
	ID          string
	Description string
	Name        string
	Team        string
	Client      string
}

type ProjectResolver struct {
	R *Project
}

func (r ProjectResolver) Client() *string {
	return &r.R.Client
}

func (r ProjectResolver) Status() *string {
	return &r.R.Status
}

func (r ProjectResolver) ID() *string {
	return &r.R.ID
}

func (r ProjectResolver) Description() *string {
	return &r.R.Description
}

func (r ProjectResolver) Name() *string {
	return &r.R.Name
}

func (r ProjectResolver) Team() *string {
	return &r.R.Team
}

type Timer struct {
	ID             string
	Description    string
	Client         *Client
	Project        *Project
	Tags           string
	TimerStart     string
	TimerEnd       string
	Teammember     string
	ElapsedSeconds int32
	IsRunning      bool
	IsBilled       bool
}

type TimerResolver struct {
	R *Timer
}

func (r TimerResolver) ElapsedSeconds() *int32 {
	return &r.R.ElapsedSeconds
}

func (r TimerResolver) IsRunning() *bool {
	return &r.R.IsRunning
}

func (r TimerResolver) IsBilled() *bool {
	return &r.R.IsBilled
}

func (r TimerResolver) Teammember() *string {
	return &r.R.Teammember
}

func (r TimerResolver) Description() *string {
	return &r.R.Description
}

func (r TimerResolver) Client() *ClientResolver {
	return &ClientResolver{r.R.Client}
}

func (r TimerResolver) Project() *ProjectResolver {
	return &ProjectResolver{r.R.Project}
}

func (r TimerResolver) Tags() *string {
	return &r.R.Tags
}

func (r TimerResolver) TimerStart() *string {
	return &r.R.TimerStart
}

func (r TimerResolver) TimerEnd() *string {
	return &r.R.TimerEnd
}

func (r TimerResolver) ID() *string {
	return &r.R.ID
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

type CreateTimerRequest struct {
	Timer CreateTimerInput
}

type StartTimerRequest struct {
	TimerId string
}

type StopTimerRequest struct {
	TimerId string
}

type DeleteTimerRequest struct {
	TimerId string
}

type HelloRequest struct {
	Name string
}

type UserRequest struct {
	ID *string
}

type AllTimerRequest struct {
	Filter AllTimerFilter
}

type TimerRequest struct {
	ID *string
}

type GqlResolver interface {
	DeleteTimer(DeleteTimerRequest) *TimerResolver
	CreateTimer(CreateTimerRequest) *TimerResolver
	StopTimer(StopTimerRequest) *TimerResolver
	AllTimer(AllTimerRequest) *[]*TimerResolver
	Timer(TimerRequest) *TimerResolver
	CreateUser(CreateUserRequest) *string
	StartTimer(StartTimerRequest) *TimerResolver
	RunningTimer() *TimerResolver
	Hello(HelloRequest) string
	User(UserRequest) *UserResolver
	AllUsers() *[]*UserResolver
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
  Timer(id: String): Timer
}

type Mutation {
  createUser(user: UserInput!): String
  createTimer(timer: CreateTimerInput!): Timer
  startTimer(timerId: String!): Timer
  stopTimer(timerId: String!): Timer
  deleteTimer(timerId: String!): Timer
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

type Timer {
  id: String
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

type Project {
  id: String
  description: String
  name: String
  team: String
  client: String
  status: String
}

type Client {
  id: String
  description: String
  name: String
  address: String
}

`
